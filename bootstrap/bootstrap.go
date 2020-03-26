package bootstrap

import (
	"jcfw.com/legal-api/config"
	"jcfw.com/legal-api/model"
	"jcfw.com/legal-api/pkg/logger"
)

//统一进行一些必要的初始化操作
func InitApp() {
	config.CheckConfig()

	//开始记录app 日志
	go logger.HandleAppChannel()

	model.InitModel()
}
