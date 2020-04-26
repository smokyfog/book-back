package model

import (
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/e421083458/gorm"
)

type Author struct {
	gorm.Model
	Name     string
	Info     string `gorm:"size:2000"`
	Portrait string
}

// AvatarURL 封面地址
func (author *Author) PortraitURL() string {
	if author.Portrait == "" {
		return ""
	}
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(author.Portrait, oss.HTTPGet, 600)
	return signedGetURL
}
