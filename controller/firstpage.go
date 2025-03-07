package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"agricultural_vision/dao/mysql"
	"agricultural_vision/models"
	"agricultural_vision/response"
)

func GetNewsHandler(c *gin.Context) {
	news, err := mysql.GetNews()
	if err != nil {
		zap.L().Error("获取新闻失败", zap.Error(err))
		response.ResponseError(c, http.StatusInternalServerError, models.CodeServerBusy)
		return
	}

	response.ResponseSuccess(c, news)
	return
}

func GetProverbHandler(c *gin.Context) {
	proverbs, err := mysql.GetProverb()
	if err != nil {
		zap.L().Error("获取谚语失败", zap.Error(err))
		response.ResponseError(c, http.StatusInternalServerError, models.CodeServerBusy)
		return
	}

	response.ResponseSuccess(c, proverbs)
	return
}

func GetCropHandler(c *gin.Context) {
	crops, err := mysql.GetCrop()
	if err != nil {
		zap.L().Error("获取农作物百科失败", zap.Error(err))
		response.ResponseError(c, http.StatusInternalServerError, models.CodeServerBusy)
		return
	}

	response.ResponseSuccess(c, crops)
	return
}

func GetVideoHandler(c *gin.Context) {
	videos, err := mysql.GetVideo()
	if err != nil {
		zap.L().Error("获取视频失败", zap.Error(err))
		response.ResponseError(c, http.StatusInternalServerError, models.CodeServerBusy)
		return
	}

	response.ResponseSuccess(c, videos)
	return
}
