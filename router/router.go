package router

import (
	"io"
	"os"

	"jcfw.com/legal-api/api/v1/account"
	"jcfw.com/legal-api/pkg/logger"

	"jcfw.com/legal-api/api"

	"jcfw.com/legal-api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//初始化路由
func NewRouter() *gin.Engine {
	go logger.HandleAppChannel()

	gin.SetMode(viper.GetString("app.mode"))

	gin.DefaultWriter = io.MultiWriter(&logger.AppLog{}, os.Stdout)
	r := gin.Default()
	//ping && health
	hp := r.Group("")
	{
		hp.GET("health", api.Health)
		hp.GET("ping", api.Health)
	}
	//v1版本接口,logger
	v1 := r.Group("api/v1", middleware.Trace(), middleware.HttpLogger(), middleware.Cors())
	v1.POST("public/login", account.Login)
	// auth := r.Group("api/v1", middleware.JWT())
	auth := v1.Group("")
	{
		acc := auth.Group("account")
		{
			acc.POST("addAccount", account.AddAcount)
			acc.GET("get")
		}
		system := auth.Group("system")
		{
			system.GET("get")
		}
	}
	return r
}
