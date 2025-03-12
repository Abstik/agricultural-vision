package routers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"agricultural_vision/controller"
	"agricultural_vision/logger"
	"agricultural_vision/middleware"
)

func SetupRouter(mode string) *gin.Engine {
	// 如果当前代码是运行模式，则将gin设置成发布模式
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
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

	// 用户模块
	userGroup := r.Group("/user")
	{
		// 用户注册
		userGroup.POST("/signup", controller.SignUpHandler)
		// 用户登录
		userGroup.POST("/login", controller.LoginHandler)
		// 发送邮箱验证码
		userGroup.POST("/email", controller.VerifyEmailHandler)
		// 修改密码
		userGroup.POST("/change-password", controller.ChangePasswordHandler)

		// jwt校验
		userGroup.Use(middleware.JWTAuthMiddleware())
		{
			// 查询个人信息
			userGroup.GET("/info", controller.GetUserInfoHandler)
			// 修改个人信息
			userGroup.PUT("/info", controller.UpdateUserInfoHandler)
			// 修改个人头像
			userGroup.POST("/avatar", controller.UpdateUserAvatarHandler)

		}
	}

	r.POST("/ai", middleware.JWTAuthMiddleware(), controller.AiHandler)

	// 首页模块
	firstPageGroup := r.Group("/firstpage")
	{
		firstPageGroup.GET("/news", controller.GetNewsHandler)
		firstPageGroup.GET("/proverb", controller.GetProverbHandler)
		firstPageGroup.GET("/crop", controller.GetCropHandler)
		firstPageGroup.GET("/video", controller.GetVideoHandler)
		firstPageGroup.GET("/poetry", controller.GetPoetryHandler)
	}

	// 社区帖子模块
	communityPost := r.Group("/community-post", middleware.JWTAuthMiddleware())
	{
		//查询所有社区
		communityPost.GET("/community", controller.CommunityHandler)
		//查询社区详情（根据id）
		communityPost.GET("/community/:id", controller.CommunityDetailHandler)

		//创建帖子
		communityPost.POST("/post", controller.CreatePostHandler)
		//查询帖子详情（根据id）
		communityPost.GET("/post/:id", controller.GetPostDetailHandler)
		//查询帖子详情列表(分页)
		communityPost.GET("/post1", controller.GetPostListHandler)
		//查询帖子详情列表（分页）（指定排序方式）
		communityPost.GET("post2", controller.GetPostListHandler2)
		//查询帖子详情列表（分页）（指定社区）
		communityPost.GET("post3", controller.GetCommunityPostListHandler)

		//投票
		communityPost.POST("/vote", controller.PostVoteController)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
