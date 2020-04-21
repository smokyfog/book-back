package model

import (
	"github.com/e421083458/gorm"
)

type Chapter struct {
	gorm.Model
	NovelId    uint64
	ChapterNum uint64
	ContentId  uint64
	Name       string
}
