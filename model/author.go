package model

import (
	"github.com/e421083458/gorm"
)

type Author struct {
	gorm.Model
	Name     string
	Info     string `gorm:"size:2000"`
	Portrait string
}
