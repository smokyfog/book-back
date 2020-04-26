package api

import (
	novel "singo/service/novel"

	"github.com/gin-gonic/gin"
)

// 小说列表
func ListNovel(c *gin.Context) {
	service := novel.ListNovelService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 小说详情
func ShowNovel(c *gin.Context) {
	service := novel.ShowNovelService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
