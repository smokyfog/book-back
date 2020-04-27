package serializer

import "singo/model"

// novel 用户序列化器
type Content struct {
	ID        uint   `json:"id"`
	ChapterId uint64 `json:"chapter_id"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
}

// 序列化小说
func BuildContent(item model.Content) Content {
	return Content{
		ID:        item.ID,
		ChapterId: item.ChapterId,
		Content:   item.Content,
		CreatedAt: item.CreatedAt.Unix(),
	}
}
