package model

import (
	"github.com/e421083458/gorm"
)

type Content struct {
	gorm.Model
	ChapterId uint64
	Content   string `gorm:"type:text"`
}
