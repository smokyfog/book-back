package service

import (
	"singo/model"
	"singo/serializer"
)

// ListChapterService 视频列表服务
type ListChapterService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// List 视频列表
func (service *ListChapterService) List() serializer.Response {
	novels := []model.Chapter{}
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

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&novels).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildChapters(novels), uint(total))
}
