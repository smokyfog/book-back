package service

import (
	"singo/model"
	"singo/serializer"
)

// ListByNovelChapterService 列表服务
type ListByNovelChapterService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// List 列表
func (service *ListByNovelChapterService) List(NovelId uint64) serializer.Response {
	capters := []model.Chapter{}
	total := 0

	if service.Limit == 0 {
		service.Limit = 6
	}

	if err := model.DB.Model(model.Chapter{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}

	if err := model.DB.Where("novel_id = ?", NovelId).Limit(service.Limit).Offset(service.Start).Find(&capters).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildChapters(capters), uint(total))
}
