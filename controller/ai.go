package controller

import (
	"agricultural_vision/middleware"
	"agricultural_vision/response"
	"go.uber.org/zap"
	"net/http"

	"github.com/gin-gonic/gin"

	"agricultural_vision/logic"
	"agricultural_vision/models"
)

func AiHandler(c *gin.Context) {
	userID, err := middleware.GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("获取userID失败", zap.Error(err))
		response.ResponseError(c, http.StatusInternalServerError, response.CodeServerBusy)
		return
	}

	// 解析前端传来的请求体
	var aiRequest response.AiRequest
	if err := c.ShouldBindJSON(&aiRequest); err != nil {
		zap.L().Error("参数校验失败", zap.Error(err))
		response.ResponseError(c, http.StatusBadRequest, response.CodeInvalidParam)
	}

	// 调用逻辑层
	aiResponse, err := logic.AiTalk(&aiRequest, userID)
	if err != nil {
		zap.L().Error("AI对话失败", zap.Error(err))
		response.ResponseError(c, http.StatusInternalServerError, response.CodeServerBusy)
	} else {
		response.ResponseSuccess(c, aiResponse)
	}
}
