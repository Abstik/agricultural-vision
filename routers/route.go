package routers

import (
	"agricultural_vision/controller"
	"agricultural_vision/logger"
	"agricultural_vision/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //gin设置成发布模式
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//注册业务路由
	v1 := r.Group("/api/v1")
	//用户注册
	v1.POST("/signup", controller.SignUpHandler)
	//用户登录
	v1.POST("/login", controller.LoginHandler)

	//应用JWT认证中间件
	v1.Use(middlewares.JWTAuthMiddleware())

	{

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
