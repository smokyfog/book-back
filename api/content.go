package api

import (
	content "singo/service/content"

	"github.com/gin-gonic/gin"
)

// 小说详情
func ShowContent(c *gin.Context) {
	service := content.ShowContentService{}
	if err := c.ShouldBind(&service); err == nil {
		contentIdStr := c.Param("contentId")
		res := service.Show(contentIdStr)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
