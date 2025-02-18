package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"agricultural_vision/dao/mysql"
	"agricultural_vision/response"
)

func GetFirstPageHandler(c *gin.Context) {
	firstPage, err := mysql.GetFirstPage()
	if err != nil {
		zap.L().Error("获取首页失败", zap.Error(err))
		response.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.ResponseSuccess(c, firstPage)
	return
}
