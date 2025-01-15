package controller

import (
	"agricultural_vision/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

//封装响应信息

type ResponseData struct {
	Code response.ResCode `json:"responsecode"`
	Msg  interface{}      `json:"msg"`
	Data interface{}      `json:"data,omitempty"`
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: response.CodeSuccess,
		Msg:  response.CodeSuccess.GetMsg(),
		Data: data,
	})
}

// 返回现成的错误状态码及信息
func ResponseError(c *gin.Context, code response.ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.GetMsg(),
		Data: nil,
	})
}

// 返回自定义的错误状态码就信息
func ResponseErrorWithMsg(c *gin.Context, code response.ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
