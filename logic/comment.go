package logic

import (
	"errors"
	"strconv"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"agricultural_vision/constants"
	"agricultural_vision/dao/mysql"
	"agricultural_vision/dao/redis"
	"agricultural_vision/models/entity"
	"agricultural_vision/models/request"
	"agricultural_vision/models/response"
)

// 创建评论
func CreateComment(createCommentRequest *request.CreateCommentRequest, userID int64) (*response.CommentResponse, error) {
	// 在mysql中查询postID是否存在
	_, err := mysql.GetPostById(createCommentRequest.PostID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 如果查询不到帖子
			return nil, constants.ErrorNoPost
		}
		return nil, err
	}

	// 在mysql中创建评论
	comment := &entity.Comment{
		Content:  createCommentRequest.Content,
		ParentID: createCommentRequest.ParentID,
		RootID:   createCommentRequest.RootID,
		AuthorID: userID,
		PostID:   createCommentRequest.PostID,
	}

	err = mysql.CreateComment(comment)
	if err != nil {
		return nil, err
	}

	// 查询作者信息
	author, err := mysql.GetUserBriefInfo(userID)
	if err != nil {
		return nil, err
	}

	// 构造返回值
	commentResponse := &response.CommentResponse{
		ID:        comment.ID,
		Content:   comment.Content,
		Author:    *author,
		CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	// 在redis中创建评论
	if createCommentRequest.ParentID == nil {
		// 在redis中创建一级评论
		err = redis.CreateFirstLevelComment(comment.ID, comment.PostID)
	} else {
		// 在redis中创建二级评论
		err = redis.CreateSecondLevelComment(*createCommentRequest.ParentID)
	}
	if err != nil {
		return nil, err
	}

	return commentResponse, nil
}

// 删除评论
func DeleteComment(commentID int64, userID int64) error {
	// 先从mysql中查找评论
	ids := []string{strconv.Itoa(int(commentID))}
	comment, err := mysql.GetCommentListByIDs(ids)
	if err != nil {
		return err
	}
	if len(comment) == 0 { // 如果未找到此评论
		return constants.ErrorNoComment
	}

	// 校验userID
	if comment[0].AuthorID != userID {
		return constants.ErrorNoPermission
	}

	// 在mysql中删除评论
	if err := mysql.DeleteComment(commentID); err != nil {
		return err
	}

	// 在redis中删除评论
	if err := redis.DeleteComment(commentID, comment[0].PostID, comment[0].ParentID); err != nil {
		return err
	}
	return nil
}

// 查询单个帖子的一级评论
func GetFirstLevelCommentList(postID int64, listRequest *request.ListRequest) (commentListResponse *response.CommentListResponse, err error) {
	commentListResponse = &response.CommentListResponse{
		Data: []*response.CommentResponse{},
	}

	//从redis中，根据指定的排序方式和查询数量，查询符合条件的id列表
	ids, total, err := redis.GetCommentIDsInOrder(listRequest, postID)
	if err != nil {
		return
	}
	commentListResponse.Total = total
	if len(ids) == 0 {
		return
	}

	//根据id列表去数据库查询评论详细信息
	comments, err := mysql.GetCommentListByIDs(ids)
	if err != nil {
		return
	}

	// 查询所有一级评论的赞成票数——切片
	voteData, err := redis.GetCommentVoteDataByIDs(ids)
	if err != nil {
		return
	}

	// 查询所有一级评论的二级评论数——切片
	commentNum, err := redis.GetSecondLevelCommentNumByIDs(ids)

	//将帖子作者及分区信息查询出来填充到帖子中
	for idx, comment := range comments {
		//查询作者简略信息
		userBriefInfo, err := mysql.GetUserBriefInfo(comment.AuthorID)
		if err != nil {
			zap.L().Error("查询作者信息失败", zap.Error(err))
			continue
		}

		//封装查询到的信息
		commentResponse := &response.CommentResponse{
			ID:           comment.ID,
			Content:      comment.Content,
			LikeCount:    voteData[idx],
			RepliesCount: int64(commentNum[idx]),
			Author:       *userBriefInfo,
			CreatedAt:    comment.CreatedAt.Format("2006-01-02 15:04:05"),
		}

		commentListResponse.Data = append(commentListResponse.Data, commentResponse)
	}
	return
}

// 查询单个一级评论的二级评论
func GetSecondLevelCommentList(commentID int64, listRequest *request.ListRequest) (commentListResponse *response.CommentListResponse, err error) {
	commentListResponse = &response.CommentListResponse{
		Data: []*response.CommentResponse{},
	}

	// 从mysql中查询二级评论
	comments, total, err := mysql.GetSecondLevelCommentList(commentID, listRequest.Page, listRequest.Size)
	if err != nil {
		return
	}
	commentListResponse.Total = int64(len(comments))
	if commentListResponse.Total == 0 {
		return
	}

	commentIDs := make([]string, commentListResponse.Total)
	for _, comment := range comments {
		commentIDs = append(commentIDs, strconv.FormatInt(comment.ID, 10)) // 提取每个二级评论的 ID
	}

	// 查询所有二级评论的赞成票数——切片
	voteData, err := redis.GetCommentVoteDataByIDs(commentIDs)
	if err != nil {
		return
	}

	//将帖子作者及分区信息查询出来填充到帖子中
	for idx, comment := range comments {
		//查询作者简略信息
		userBriefInfo, err := mysql.GetUserBriefInfo(comment.AuthorID)
		if err != nil {
			zap.L().Error("查询作者信息失败", zap.Error(err))
			continue
		}

		//封装查询到的信息
		commentResponse := &response.CommentResponse{
			ID:        comment.ID,
			Content:   comment.Content,
			LikeCount: voteData[idx],
			Author:    *userBriefInfo,
			CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
		}

		commentListResponse.Data = append(commentListResponse.Data, commentResponse)
	}
	commentListResponse.Total = total
	return
}
