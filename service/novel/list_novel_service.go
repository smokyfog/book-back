package service

import (
	"singo/model"
	"singo/serializer"
)

// ListNovelService 视频列表服务
type ListNovelService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// List 视频列表
func (service *ListNovelService) List() serializer.Response {
	novels := []model.Novel{}
	total := 0

	if service.Limit == 0 {
		service.Limit = 6
	}

	if err := model.DB.Model(model.Novel{}).Count(&total).Error; err != nil {
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

	return serializer.BuildListResponse(serializer.BuildNovels(novels), uint(total))
}
