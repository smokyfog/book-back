package api

import (
	capter "singo/service/capter"

	"github.com/gin-gonic/gin"
)

// 视频列表
func ListChapter(c *gin.Context) {
	service := capter.ListChapterService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
