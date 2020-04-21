package model

//执行数据迁移

func migration() {
	// 自动迁移模式
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Video{})
	DB.AutoMigrate(&Novel{})
	DB.AutoMigrate(&Chapter{})
	DB.AutoMigrate(&Author{})
	DB.AutoMigrate(&Content{})
}
