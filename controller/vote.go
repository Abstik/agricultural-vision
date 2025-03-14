package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"agricultural_vision/constants"
	"agricultural_vision/logic"
	"agricultural_vision/middleware"
	"agricultural_vision/models/request"
)

// 投票功能
func PostVoteController(c *gin.Context) {
	p := new(request.VotePostRequest)
	//参数校验
	err := c.ShouldBindJSON(p)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, http.StatusBadRequest, constants.CodeInvalidParam)
			return
		} else {
			errData := RemoveTopStruct(errs.Translate(trans)) //翻译错误
			ResponseError(c, http.StatusBadRequest, errData)
			return
		}
	}

	//业务逻辑
	userID, err := middleware.GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, constants.CodeInvalidParam)
		return
	}

	err = logic.VoteForPost(userID, p)
	if err != nil {
		zap.L().Error("投票失败", zap.Error(err))
		ResponseError(c, http.StatusBadRequest, constants.CodeInvalidParam)
		return
	}

	//返回响应
	ResponseSuccess(c, nil)
	return
}
