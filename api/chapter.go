package api

import (
	capter "singo/service/capter"
	"strconv"

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

// 小说详情
func ListByNovelChapter(c *gin.Context) {
	service := capter.ListByNovelChapterService{}
	if err := c.ShouldBind(&service); err == nil {
		novelIdStr := c.Param("novelId")
		intNum, _ := strconv.Atoi(novelIdStr)
		novelId := uint64(intNum)
		res := service.List(novelId)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
