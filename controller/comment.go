package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"agricultural_vision/constants"
	"agricultural_vision/logic"
	"agricultural_vision/middleware"
	"agricultural_vision/models/request"
)

// 创建评论
func CreateCommentHandler(c *gin.Context) {
	// 获取参数
	createCommentRequest := &request.CreateCommentRequest{}
	if err := c.ShouldBindJSON(createCommentRequest); err != nil {
		zap.L().Error("参数不正确", zap.Error(err))
		ResponseError(c, http.StatusBadRequest, constants.CodeInvalidParam)
		return
	}

	// 获取userID
	id, err := middleware.GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("获取userID失败", zap.Error(err))
		ResponseError(c, http.StatusInternalServerError, constants.CodeServerBusy)
		return
	}

	// 创建评论
	err = logic.CreateComment(createCommentRequest, id)
	if err != nil {
		zap.L().Error("创建评论失败", zap.Error(err))
		ResponseError(c, http.StatusInternalServerError, constants.CodeServerBusy)
		return
	}

	ResponseSuccess(c, constants.CodeSuccess)
}

// 删除评论
func DeleteCommentHandler(c *gin.Context) {
	// 获取参数
	commentIDStr := c.Param("comment_id")
	commentID, err1 := strconv.ParseInt(commentIDStr, 10, 64)
	if err1 != nil {
		zap.L().Error("参数不正确", zap.Error(err1))
		ResponseError(c, http.StatusBadRequest, constants.CodeInvalidParam)
		return
	}

	// 获取用户id
	userID, err := middleware.GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("获取userID失败", zap.Error(err))
		ResponseError(c, http.StatusInternalServerError, constants.CodeServerBusy)
		return
	}

	// 删除帖子
	err = logic.DeleteComment(commentID, userID)
	if err != nil {
		zap.L().Error("删除评论失败", zap.Error(err))
		ResponseError(c, http.StatusInternalServerError, constants.CodeServerBusy)
		return
	}

	ResponseSuccess(c, constants.CodeSuccess)
}

// 查询一级评论
func GetFirstLevelCommentListHandler(c *gin.Context) {
	// 获取参数
	postIDStr := c.Param("post_id")
	postID, err1 := strconv.ParseInt(postIDStr, 10, 64)

	listRequest := &request.ListRequest{
		Page:  1,
		Size:  10,
		Order: constants.OrderTime,
	}
	err2 := c.ShouldBindQuery(listRequest)

	if err1 != nil || err2 != nil {
		zap.L().Error("参数不正确", zap.Error(err1), zap.Error(err2))
		ResponseError(c, http.StatusBadRequest, constants.CodeInvalidParam)
		return
	}

	// 查询一级评论
	commentListResponse, err := logic.GetFirstLevelCommentList(postID, listRequest)
	if err != nil {
		zap.L().Error("查询一级评论失败", zap.Error(err))
		ResponseError(c, http.StatusInternalServerError, constants.CodeServerBusy)
		return
	}

	ResponseSuccess(c, commentListResponse)
}

// 查询二级评论
func GetSecondLevelCommentListHandler(c *gin.Context) {
	commentIDStr := c.Param("comment_id")
	commentID, err1 := strconv.ParseInt(commentIDStr, 10, 64)

	listRequest := &request.ListRequest{
		Page: 1,
		Size: 10,
	}
	err2 := c.ShouldBindQuery(listRequest)

	if err1 != nil || err2 != nil {
		zap.L().Error("参数不正确", zap.Error(err1), zap.Error(err2))
		ResponseError(c, http.StatusBadRequest, constants.CodeInvalidParam)
		return
	}

	commentListResponse, err := logic.GetSecondLevelCommentList(commentID, listRequest)
	if err != nil {
		zap.L().Error("查询二级评论失败", zap.Error(err))
		ResponseError(c, http.StatusInternalServerError, constants.CodeServerBusy)
		return
	}
	ResponseSuccess(c, commentListResponse)
}

// 评论投票
func CommentVoteController(c *gin.Context) {
	votePostRequest := new(request.VoteRequest)
	err := c.ShouldBindJSON(votePostRequest)
	if err != nil {
		zap.L().Error("参数不正确", zap.Error(err))
		ResponseError(c, http.StatusBadRequest, constants.CodeInvalidParam)
	}

	userID, err := middleware.GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, constants.CodeServerBusy)
		return
	}

	err = logic.CommentVote(userID, votePostRequest)
	if err != nil {
		zap.L().Error("评论投票失败", zap.Error(err))
		ResponseError(c, http.StatusBadRequest, constants.CodeInvalidParam)
		return
	}

	ResponseSuccess(c, constants.CodeSuccess)
}
