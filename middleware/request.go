package middleware

import (
	"github.com/gin-gonic/gin"
)

// 从请求上下文中获取当前用户ID
func GetCurrentUserID(c *gin.Context) (userID int64, err error) {
	// 如果用户已登录则可以在请求上下文中获取到userID, 如果获取不到则用户未登录
	uid, ok := c.Get("userID")
	// 如果获取不到则返回错误
	if !ok {
		return
	}

	// 类型断言失败则直接返回错误
	userID, ok = uid.(int64)
	if !ok {
		// 类型断言失败
		return
	}

	// 没有错误，直接返回
	return
}
