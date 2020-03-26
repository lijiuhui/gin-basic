package logger

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type logChan chan string

var SqlChannel = make(logChan, 100)
var AppChannel = make(logChan, 100)
var AccessChannel = make(logChan, 100)

//接收sql channel 内容
func HandleSqlChannel() {
	logToFile(viper.GetString("log.sql_log"), SqlChannel)
}

//接收access channel内容
func HandleAccessChannel() {
	logToFile(viper.GetString("log.access_log"), AccessChannel)
}

//接收app log
func HandleAppChannel() {
	logToFile(viper.GetString("log.app_log"), AppChannel)
}

//记录日志到文件中
func logToFile(file string, c logChan) {
	if f, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err != nil {
		log.Println(err)
	} else {
		for log := range c {
			_, _ = f.WriteString(log + "\n")
		}
	}
	return
}

//logger instance
func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger

}

//将日志放入channel
func ToChan(c logChan, s string) {
	c <- s
}
