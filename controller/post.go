package controller

import (
	"agricultural_vision/models/entity"
	response2 "agricultural_vision/models/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"agricultural_vision/logic"
	"agricultural_vision/middleware"
	"agricultural_vision/models"
	"agricultural_vision/response"
)

// 创建帖子
func CreatePostHandler(c *gin.Context) {
	//1.获取参数及参数的校验
	p := new(entity.Post)
	//将参数绑定到p中
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("请求参数错误", zap.Error(err))
		response.ResponseError(c, http.StatusBadRequest, response2.CodeInvalidParam)
		return
	}
	//在请求上下文中获取userID
	userID, err := middleware.GetCurrentUserID(c)
	if err != nil {
		response.ResponseError(c, http.StatusInternalServerError, response2.CodeServerBusy)
		return
	}
	p.AuthorID = userID

	//2.创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("创建帖子失败", zap.Error(err))
		response.ResponseError(c, http.StatusInternalServerError, response2.CodeServerBusy)
		return
	}

	//3.返回响应
	response.ResponseSuccess(c, response2.CodeSuccess)
}

// 查询帖子详情
func GetPostDetailHandler(c *gin.Context) {
	//1.获取参数（从URL中获取帖子的id）
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("获取帖子详情的参数不正确", zap.Error(err))
		response.ResponseError(c, http.StatusBadRequest, response2.CodeInvalidParam)
		return
	}

	//2.根据id取出帖子数据（查数据库）
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("获取帖子详情失败", zap.Error(err))
		response.ResponseError(c, http.StatusInternalServerError, response2.CodeServerBusy)
		return
	}

	//3.返回响应
	response.ResponseSuccess(c, data)
}

// 查询帖子列表
func GetPostListHandler(c *gin.Context) {
	//获取分页参数
	//page表示第几页，size表示每页几条数据
	pageNumStr := c.Query("page")
	pageSizeStr := c.Query("size")

	var (
		pageNum  int64
		pageSize int64
		err      error
	)

	pageNum, err = strconv.ParseInt(pageNumStr, 10, 64)
	if err != nil {
		pageNum = 0
	}
	pageSize, err = strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		pageSize = 0
	}

	//获取数据
	data, err := logic.GetPostList(pageNum, pageSize)
	if err != nil {
		zap.L().Error("获取帖子列表失败", zap.Error(err))
		response.ResponseError(c, http.StatusBadRequest, response2.CodeInvalidParam)
		return
	}
	response.ResponseSuccess(c, data)
}

// 升级版查询帖子列表接口
// 根据前端传来的参数,动态获取帖子列表（按照创建时间or分数排序）
// 1.获取参数
// 2.去redis查询id列表
// 3.根据id去数据库查询帖子详细信息
func GetPostListHandler2(c *gin.Context) {
	//初始化结构体时指定初始默认参数
	p := &response.PostListParam{
		Page:  1,
		Size:  10,
		Order: response2.OrderTime,
	}
	err := c.ShouldBindQuery(p)
	if err != nil {
		zap.L().Error("请求参数错误", zap.Error(err))
		response.ResponseError(c, http.StatusBadRequest, response2.CodeInvalidParam)
		return
	}

	//获取数据
	data, err := logic.GetPostList2(p)
	if err != nil {
		zap.L().Error("指定顺序查询帖子列表失败", zap.Error(err))
		response.ResponseError(c, http.StatusInternalServerError, response2.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, data)
	return
}

// 根据社区查询该社区分类下的帖子详情列表
func GetCommunityPostListHandler(c *gin.Context) {
	//初始化结构体时指定初始默认参数
	p := &response.CommunityPostListParam{
		PostListParam: response.PostListParam{
			Page:  1,
			Size:  10,
			Order: response2.OrderTime,
		},
	}
	err := c.ShouldBindQuery(p)
	if err != nil {
		zap.L().Error("请求参数错误", zap.Error(err))
		response.ResponseError(c, http.StatusBadRequest, response2.CodeInvalidParam)
		return
	}

	//根据社区查询该社区分类下的帖子列表
	data, err := logic.GetCommunityPostList(p)
	if err != nil {
		zap.L().Error("根据社区查询帖子列表失败", zap.Error(err))
		response.ResponseError(c, http.StatusInternalServerError, response2.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, data)
	return
}
