package response

import (
	"agricultural_vision/models/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

//封装响应信息

type ResponseData struct {
	Code response.ResCode `json:"code"`           // 编码，成功为1，失败为0
	Msg  interface{}      `json:"msg"`            // 响应码对应的响应信息
	Data interface{}      `json:"data,omitempty"` // 返回的数据
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: 1,
		Msg:  response.CodeSuccess,
		Data: data,
	})
}

func ResponseError(c *gin.Context, httpStatus int, msg interface{}) {
	c.JSON(httpStatus, &ResponseData{
		Code: 0,
		Msg:  msg,
		Data: nil,
	})
}
