package serializer

import "singo/model"

// novel 用户序列化器
type Novel struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	AuthorId  uint   `json:"author_id"`
	Score     uint   `json:"score"`
	Cover     string `json:"cover"`
	WordCount uint   `json:"word_count"`
	Type      uint   `json:"type"`
	Hits      uint   `json:"hits"`
	CreatedAt int64  `json:"created_at"`
}

// 序列化视频
func BuildNovel(item model.Novel) Novel {
	// user, _ := model.GetUser(item.UserID)
	return Novel{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		AuthorId:  item.AuthorId,
		Score:     item.Score,
		Cover:     item.CoverURL(),
		WordCount: item.WordCount,
		Type:      item.Type,
		Hits:      item.Hits,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// 序列化视频列表
func BuildNovels(items []model.Novel) (novels []Novel) {
	for _, item := range items {
		novel := BuildNovel(item)
		novels = append(novels, novel)
	}
	return novels
}
