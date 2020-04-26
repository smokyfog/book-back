package router

import (
	"singo/api"

	"github.com/gin-gonic/gin"
)

func InitNovelRouter(Router *gin.RouterGroup) {
	NovelRouter := Router.Group("novel")
	{
		NovelRouter.GET("list", api.ListNovel)
	}
}
