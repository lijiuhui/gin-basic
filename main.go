package main

import (
	"fmt"

	"jcfw.com/legal-api/router"

	"jcfw.com/legal-api/bootstrap"

	"github.com/spf13/viper"
)

func main() {
	//初始化
	bootstrap.InitApp()
	//路由注册
	r := router.NewRouter()

	// 启动项目
	addr := viper.GetString("app.address")
	port := viper.GetString("app.port")
	if err := r.Run(addr + ":" + port); err != nil {
		panic(fmt.Sprintf("gin启动失败：%s", err))
	}
}
