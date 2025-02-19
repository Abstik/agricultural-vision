package routers

import (
	"agricultural_vision/controller"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"agricultural_vision/logger"
	"agricultural_vision/middleware"
)

func SetupRouter(mode string) *gin.Engine {
	// 如果当前代码是运行模式，则将gin设置成发布模式
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //gin设置成发布模式
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.Use(cors.New(cors.Config{
		// 允许的域名（前端地址）
		AllowOrigins: []string{"*"}, // 允许所有源
		// 允许的请求方法
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// 允许的请求头
		AllowHeaders: []string{"Content-Type", "Authorization"},
		// 允许携带认证信息
		AllowCredentials: true,
	}))

	// 用户注册
	r.POST("/signup", controller.SignUpHandler)
	// 用户登录
	r.POST("/login", controller.LoginHandler)
	// 发送邮箱验证码
	r.POST("/email", controller.VerifyEmailHandler)
	// 修改密码
	r.POST("/changePassword", controller.ChangePasswordHandler)

	// 应用JWT认证中间件
	r.Use(middleware.JWTAuthMiddleware())
	{
		// 查询个人信息
		r.GET("/user", controller.GetUserInfoHandler)
		// 修改个人信息
		r.PUT("/user", controller.UpdateUserInfoHandler)
		// ai对话
		r.POST("/ai", controller.AiHandler)

		// 查询首页信息
		firstPageGroup := r.Group("/firstPage")
		{
			firstPageGroup.GET("/news", controller.GetNewsHandler)
			firstPageGroup.GET("/proverb", controller.GetProverbHandler)
			firstPageGroup.GET("/crop", controller.GetCropHandler)
			firstPageGroup.GET("/video", controller.GetVideoHandler)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
