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

			// 查询用户的帖子列表
			userGroup.GET("/posts", controller.GetUserPostListHandler)
			// 查询用户的点赞列表
			userGroup.GET("/likes", controller.GetUserLikeListHandler)
			// 查询用户的评论列表
			userGroup.GET("/comments", controller.GetUserCommentListHandler)
		}
	}

	// AI模块
	AIGroup := r.Group("/ai")
	{
		AIGroup.POST("/ai", middleware.JWTAuthMiddleware(), controller.AiHandler)
	}

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
		// 查询社区列表
		communityPost.GET("/community", controller.CommunityHandler)
		// 查询社区详情
		communityPost.GET("/community/:id", controller.CommunityDetailHandler)

		// 创建帖子
		communityPost.POST("/post", controller.CreatePostHandler)
		// 删除帖子
		communityPost.DELETE("/post/:id", controller.DeletePostHandler)
		// 查询帖子列表（指定排序方式，默认按时间倒序）
		communityPost.GET("/posts", controller.GetPostListHandler)
		// 查询帖子列表（指定社区）（指定排序方式，默认按时间倒序）
		communityPost.GET("/community/:id/posts", controller.GetCommunityPostListHandler)
		// 帖子投票
		communityPost.POST("/post/vote", controller.PostVoteController)
		// 发布评论
		communityPost.POST("/comment", controller.CreateCommentHandler)
		// 查询帖子评论
		communityPost.GET("/comment", controller.GetCommentListHandler)
		// 评论投票
		communityPost.POST("/comment/vote", controller.CommentVoteController)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
