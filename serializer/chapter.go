package serializer

import "singo/model"

// novel 用户序列化器
type Chapter struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	ContentId  uint64 `json:"content_id"`
	ChapterNum uint64 `json:"chapter_num"`
	NovelId    uint64 `json:"novel_id"`
	CreatedAt  int64  `json:"created_at"`
}

// 序列化视频
func BuildChapter(item model.Chapter) Chapter {
	// user, _ := model.GetUser(item.UserID)
	return Chapter{
		ID:         item.ID,
		Name:       item.Name,
		ContentId:  item.ContentId,
		ChapterNum: item.ChapterNum,
		NovelId:    item.NovelId,
		CreatedAt:  item.CreatedAt.Unix(),
	}
}

// 序列化视频列表
func BuildChapters(items []model.Chapter) (novels []Chapter) {
	for _, item := range items {
		novel := BuildChapter(item)
		novels = append(novels, novel)
	}
	return novels
}
