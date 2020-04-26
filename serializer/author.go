package serializer

import "singo/model"

// novel 用户序列化器
type Author struct {
	ID        uint   `json:"id"`
	Info      string `json:"info"`
	Portrait  string `json:"portrait"`
	CreatedAt int64  `json:"created_at"`
}

// 序列化
func BuildAuthor(item model.Author) Author {
	// user, _ := model.GetUser(item.UserID)
	return Author{
		ID:        item.ID,
		Info:      item.Info,
		Portrait:  item.PortraitURL(),
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// 序列化列表
func BuildAuthors(items []model.Author) (capters []Author) {
	for _, item := range items {
		novel := BuildAuthor(item)
		capters = append(capters, novel)
	}
	return capters
}
