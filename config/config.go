package config

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
)

//检测配置文件
func CheckConfig() {
	dir, _ := os.Getwd()
	configPath := path.Join(dir, "/config")
	configName := "config"
	viper.SetConfigName(configName) // 指定配置文件的文件名称(不需要指定配置文件的扩展名)
	viper.AddConfigPath(configPath) // 设置配置文件的搜索目录
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("配置文件读取失败：%s", err))
	}
}
