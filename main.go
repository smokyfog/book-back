package main

import (
	"singo/conf"
	"singo/init/initRouter"
	"singo/init/log"
)

func main() {
	// 从配置文件读取配置
	conf.Init()
	log.Init()
	// 装载路由
	r := initRouter.NewRouter()
	r.Run(":3000")
}
