package router

import (
	"singo/api"

	"github.com/gin-gonic/gin"
)

func InitContentRouter(Router *gin.RouterGroup) {
	ContentRouter := Router.Group("content")
	{
		ContentRouter.GET(":contentId", api.ShowContent)
	}
}
