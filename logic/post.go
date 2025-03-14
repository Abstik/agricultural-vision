package logic

import (
	"go.uber.org/zap"

	"agricultural_vision/dao/mysql"
	"agricultural_vision/dao/redis"
	"agricultural_vision/models/entity"
	"agricultural_vision/models/request"
	"agricultural_vision/models/response"
)

// 创建帖子
func CreatePost(createPostRequest *request.CreatePostRequest, authorID int64) (err error) {
	post := &entity.Post{
		Content:     createPostRequest.Content,
		Image:       createPostRequest.Image,
		AuthorID:    authorID,
		CommunityID: createPostRequest.CommunityID,
	}

	//2.保存到数据库
	err = mysql.CreatePost(post)
	if err != nil {
		return
	}

	//3.保存到redis
	err = redis.CreatePost(post.ID, post.CommunityID)
	return
}

// 查询帖子列表，并按照指定方式排序
func GetPostList(p *request.PostListRequest) (postResponses []*response.PostResponse, err error) {
	//从redis中，根据指定的排序方式和查询数量，查询符合条件的id列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}

	//根据id列表去数据库查询帖子详细信息
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}

	//查询所有帖子的赞成票数——切片
	voteData, err := redis.GetPostVoteDataByIDs(ids)
	if err != nil {
		return
	}

	// 查询所有帖子的评论数——切片
	commentNum, err := redis.GetCommentNumByID(ids)

	//将帖子作者及分区信息查询出来填充到帖子中
	for idx, post := range posts {
		//查询作者简略信息
		userBriefInfo, err := mysql.GetUserBriefInfo(post.AuthorID)
		if err != nil {
			zap.L().Error("查询作者信息失败", zap.Error(err))
			continue
		}

		//查询社区详情
		community, err := mysql.GetCommunityById(post.CommunityID)
		if err != nil {
			zap.L().Error("查询社区详情失败", zap.Error(err))
			continue
		}

		//封装查询到的信息
		postResponse := &response.PostResponse{
			ID:           post.ID,
			Content:      post.Content,
			Image:        post.Image,
			Author:       *userBriefInfo,
			LikeCount:    voteData[idx],
			CommentCount: int64(commentNum[idx]),
			CreatedAt:    post.CreatedAt.Format("2006-01-02 15:04:05"),
			Community:    response.CommunityBriefResponse{CommunityID: community.ID, CommunityName: community.CommunityName},
		}

		postResponses = append(postResponses, postResponse)
	}
	return
}

// 查询该社区下的帖子列表，并按指定方式排序
func GetCommunityPostList(postListRequest *request.CommunityPostListRequest) (postResponses []*response.PostResponse, err error) {
	//从redis中，根据指定的排序方式和查询数量，查询符合条件的分页后的id列表
	ids, err := redis.GetCommunityPostIDsInOrder(postListRequest)
	if err != nil {
		return
	}

	//根据id列表去数据库查询帖子详细信息
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}

	//根据帖子id列表查询所有帖子的赞成票数
	voteData, err := redis.GetPostVoteDataByIDs(ids)
	if err != nil {
		return
	}

	// 查询所有帖子的评论数——切片
	commentNum, err := redis.GetCommentNumByID(ids)

	//将帖子作者及分区信息查询出来填充到帖子中
	for idx, post := range posts {
		//根据作者id查询作者信息
		userBriefInfo, err := mysql.GetUserBriefInfo(post.AuthorID)
		if err != nil {
			zap.L().Error("查询作者信息失败", zap.Error(err))
			continue
		}

		//根据社区id查询社区详情
		community, err := mysql.GetCommunityById(post.CommunityID)
		if err != nil {
			zap.L().Error("查询社区详情失败", zap.Error(err))
			continue
		}

		//封装查询到的信息
		postDetail := &response.PostResponse{
			ID:           post.ID,
			Content:      post.Content,
			Image:        post.Image,
			Author:       *userBriefInfo,
			LikeCount:    voteData[idx],
			CommentCount: int64(commentNum[idx]),
			CreatedAt:    post.CreatedAt.Format("2006-01-02 15:04:05"),
			Community:    response.CommunityBriefResponse{CommunityID: community.ID, CommunityName: community.CommunityName},
		}

		postResponses = append(postResponses, postDetail)
	}
	return
}
