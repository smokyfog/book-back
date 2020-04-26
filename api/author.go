package api

import (
	author "singo/service/author"

	"github.com/gin-gonic/gin"
)

// 视频列表
func ListAuthor(c *gin.Context) {
	service := author.ListAuthorService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
