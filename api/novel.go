package api

import (
	novel "singo/service/novel"

	"github.com/gin-gonic/gin"
)

// 视频列表
func ListNovel(c *gin.Context) {
	service := novel.ListNovelService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
