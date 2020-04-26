package service

import (
	"singo/model"
	"singo/serializer"
)

// ShowNovelService 小说详情
type ShowNovelService struct {
}

// 查看小说
func (service *ShowNovelService) Show(id string) serializer.Response {
	var novel model.Novel
	err := model.DB.First(&novel, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "小说不存在",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildNovel(novel),
	}
}
