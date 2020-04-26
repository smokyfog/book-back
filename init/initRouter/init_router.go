package initRouter

import (
	"os"
	"singo/api"
	"singo/middleware"
	"singo/router"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		v1.GET("video/:id", api.ShowVideo)
		v1.GET("videos", api.ListVideo)
		v1.PUT("video/:id", api.UpdateVideo)
		v1.DELETE("video/:id", api.DeleteVideo)

		// 排行榜
		v1.GET("rank/daily", api.DailyRank)
		// 上传oss获取
		v1.POST("upload/token", api.UploadToken)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
			// 提交视频
			auth.POST("videos", api.CreateVideo)
		}
	}
	router.InitNovelRouter(v1)
	router.InitChapterRouter(v1)
	router.InitAuthorRouter(v1)
	return r
}
