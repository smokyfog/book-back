package router

import (
	"singo/api"

	"github.com/gin-gonic/gin"
)

func InitAuthorRouter(Router *gin.RouterGroup) {
	AuthorRouter := Router.Group("author")
	{
		AuthorRouter.GET("list", api.ListAuthor)
	}
}
