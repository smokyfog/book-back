package router

import (
	"singo/api"
	"singo/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func InitAuthRouter(Router *gin.RouterGroup) {
	AuthorRouter := Router.Group("Auth")
	AuthorRouter.Use(jwt.JWTAuth())
	{
		AuthorRouter.GET("test", api.GetDataByTime)
	}
}
