package router

import (
	"singo/api"

	"github.com/gin-gonic/gin"
)

func InitNovelRouter(Router *gin.RouterGroup) {
	NovelRouter := Router.Group("novel")
	{
		NovelRouter.GET("novels", api.ListNovel)
		NovelRouter.GET("detail/:id", api.ShowNovel)
	}
}
