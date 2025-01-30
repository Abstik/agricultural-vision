package response

import (
	"agricultural_vision/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//封装响应信息

type ResponseData struct {
	Code models.ResCode `json:"code"`           // 编码，成功为1，失败为0
	Msg  interface{}    `json:"msg"`            // 响应码对应的响应信息
	Data interface{}    `json:"data,omitempty"` // 返回的数据
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: 1,
		Msg:  models.CodeSuccess,
		Data: data,
	})
}

func ResponseError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: 0,
		Msg:  msg,
		Data: nil,
	})
}
