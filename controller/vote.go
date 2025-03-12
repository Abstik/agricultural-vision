package controller

import (
	"agricultural_vision/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"agricultural_vision/logic"
	"agricultural_vision/middleware"
	"agricultural_vision/models"
)

// 投票功能
func PostVoteController(c *gin.Context) {
	p := new(models.VoteData)
	//参数校验
	err := c.ShouldBindJSON(p)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.ResponseError(c, http.StatusBadRequest, models.CodeInvalidParam)
			return
		} else {
			errData := RemoveTopStruct(errs.Translate(trans)) //翻译错误
			response.ResponseError(c, http.StatusBadRequest, errData)
			return
		}
	}

	//业务逻辑
	userID, err := middleware.GetCurrentUserID(c)
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, models.CodeInvalidParam)
		return
	}
	err = logic.VoteForPost(userID, p)
	if err != nil {
		zap.L().Error("投票失败", zap.Error(err))
		response.ResponseError(c, http.StatusBadRequest, models.CodeInvalidParam)
		return
	}

	//返回响应
	response.ResponseSuccess(c, nil)
	return
}
