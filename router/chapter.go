package router

import (
	"singo/api"

	"github.com/gin-gonic/gin"
)

func InitChapterRouter(Router *gin.RouterGroup) {
	ChapterRouter := Router.Group("chapter")
	{
		ChapterRouter.GET("list", api.ListChapter)
	}
}
