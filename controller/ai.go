package controller

import (
	"go.uber.org/zap"
	"net/http"

	"github.com/gin-gonic/gin"

	"agricultural_vision/logic"
	"agricultural_vision/models"
	"agricultural_vision/response"
)

func AiHandler(c *gin.Context) {
	// 解析前端传来的请求体
	var aiRequest models.AiRequest
	if err := c.ShouldBindJSON(&aiRequest); err != nil {
		zap.L().Error("参数校验失败", zap.Error(err))
		response.ResponseError(c, http.StatusBadRequest, models.CodeInvalidParam)
	}

	// 调用逻辑层
	aiResponse, err := logic.AiTalk(&aiRequest)
	if err != nil {
		zap.L().Error("AI对话失败", zap.Error(err))
		response.ResponseError(c, http.StatusInternalServerError, models.CodeServerBusy)
	} else {
		response.ResponseSuccess(c, aiResponse)
	}
}
