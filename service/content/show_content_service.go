package service

import (
	"singo/model"
	"singo/serializer"
)

// ShowNovelService 小说详情
type ShowContentService struct {
}

// 查看小说
func (service *ShowContentService) Show(id string) serializer.Response {
	var content model.Content
	err := model.DB.First(&content, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "章节不存在",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildContent(content),
	}
}
