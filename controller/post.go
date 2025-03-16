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

// 发布帖子
func CreatePostHandler(c *gin.Context) {
	//1.获取参数及参数的校验
	p := new(request.CreatePostRequest)
	//将参数绑定到p中
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("请求参数错误", zap.Error(err))
		ResponseError(c, http.StatusBadRequest, constants.CodeInvalidParam)
		return
	}
	//在请求上下文中获取userID
	userID, err := middleware.GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, constants.CodeServerBusy)
		return
	}

	//2.创建帖子
	if err := logic.CreatePost(p, userID); err != nil {
		zap.L().Error("创建帖子失败", zap.Error(err))
		ResponseError(c, http.StatusInternalServerError, constants.CodeServerBusy)
		return
	}

	//3.返回响应
	ResponseSuccess(c, constants.CodeSuccess)
}

// 删除帖子
func DeletePostHandler(c *gin.Context) {
	postID := c.Param("post_id")

	postIDStr, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		zap.L().Error("请求参数错误", zap.Error(err))
		ResponseError(c, http.StatusBadRequest, constants.CodeInvalidParam)
		return
	}

	if err := logic.DeletePost(postIDStr); err != nil {
		zap.L().Error("删除帖子失败", zap.Error(err))
		ResponseError(c, http.StatusInternalServerError, constants.CodeServerBusy)
		return
	}
	ResponseSuccess(c, constants.CodeSuccess)
}

// 查询帖子列表
// 根据前端传来的参数,动态获取帖子列表（按照创建时间or分数排序）
// 1.获取参数
// 2.去redis查询id列表
// 3.根据id去数据库查询帖子详细信息
func GetPostListHandler(c *gin.Context) {
	//初始化结构体时指定初始默认参数
	p := &request.ListRequest{
		Page:  1,
		Size:  10,
		Order: constants.OrderTime,
	}
	err := c.ShouldBindQuery(p)
	if err != nil {
		zap.L().Error("请求参数错误", zap.Error(err))
		ResponseError(c, http.StatusBadRequest, constants.CodeInvalidParam)
		return
	}

	//获取数据
	data, err := logic.GetPostList(p)
	if err != nil {
		zap.L().Error("指定顺序查询帖子列表失败", zap.Error(err))
		ResponseError(c, http.StatusInternalServerError, constants.CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
	return
}

// 查询该社区分类下的帖子详情列表
func GetCommunityPostListHandler(c *gin.Context) {
	//初始化结构体时指定初始默认参数
	p := &request.ListRequest{
		Page:  1,
		Size:  10,
		Order: constants.OrderTime,
	}
	err1 := c.ShouldBindQuery(p)

	communityIDStr := c.Param("id")
	communityID, err2 := strconv.ParseInt(communityIDStr, 10, 64)

	if err1 != nil || err2 != nil {
		zap.L().Error("请求参数错误", zap.Error(err1), zap.Error(err2))
		ResponseError(c, http.StatusBadRequest, constants.CodeInvalidParam)
		return
	}

	//根据社区查询该社区分类下的帖子列表
	data, err := logic.GetCommunityPostList(p, communityID)
	if err != nil {
		zap.L().Error("根据社区查询帖子列表失败", zap.Error(err))
		ResponseError(c, http.StatusInternalServerError, constants.CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
	return
}
