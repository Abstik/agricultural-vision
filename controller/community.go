package controller

import (
	response2 "agricultural_vision/models/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"agricultural_vision/logic"
	"agricultural_vision/models"
	"agricultural_vision/response"
)

// 社区模块

// 查询所有社区
func CommunityHandler(c *gin.Context) {
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("获取社区列表失败", zap.Error(err))
		response.ResponseError(c, http.StatusInternalServerError, response2.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, data)
}

// 查询社区详情
func CommunityDetailHandler(c *gin.Context) {
	//1.获取社区id
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	//如果获取请求参数失败
	if err != nil {
		zap.L().Error("获取社区详情的参数不正确", zap.Error(err))
		response.ResponseError(c, http.StatusBadRequest, response2.CodeInvalidParam)
		return
	}

	//查询到所有的社区，以列表形式返回
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("获取社区详情失败", zap.Error(err))
		response.ResponseError(c, http.StatusInternalServerError, response2.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, data)
}
