package model

import (
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/e421083458/gorm"
)

type Novel struct {
	gorm.Model
	Title     string
	Info      string `gorm:"size:1000"`
	AuthorId  uint
	Score     uint
	Cover     string `gorm:"size:1000"`
	WordCount uint
	Type      uint
	genre     uint
	Hits      uint
}

// AvatarURL 封面地址
func (novel *Novel) CoverURL() string {
	if novel.Cover == "" {
		return ""
	}
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(novel.Cover, oss.HTTPGet, 600)
	return signedGetURL
}
