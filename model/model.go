package model

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"jcfw.com/legal-api/pkg/logger"
	"jcfw.com/legal-api/util"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

//初始化db
func InitModel() {
	addr := viper.GetString("mysql.address")
	port := viper.GetString("mysql.port")
	user := viper.GetString("mysql.user")
	pwd := viper.GetString("mysql.password")
	database := viper.GetString("mysql.database")
	//"${user}:${pwd}@tcp(${addr}:${port})/${db}?charset=utf8&parseTime=True&loc=Local"
	connectStr := user + ":" + pwd + "@tcp(" + addr + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", connectStr)
	if err != nil {
		// log.Fatalln("数据库连接失败：%s", err)
	}
	var mode bool
	if viper.GetString("app.mode") == "debug" {
		mode = true
	}
	db.LogMode(mode)
	//禁用表名复数
	db.SingularTable(true)

	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)
	//开始监听数据
	DB = db

	go logger.HandleSqlChannel()
}

//获取一个db instance
func GetDB(c *gin.Context) *gorm.DB {
	db := DB.New()
	traceId := util.GetTraceId(c)
	saasID, _ := c.Get(util.SaasIDKey)
	uuid := util.GenUUID()
	companyId, _ := c.Get(util.CompanyIDKey)
	db.SetLogger(&logger.GormLogger{TraceId: traceId})
	db.Callback().Create().Replace("gorm:before_create", func(scope *gorm.Scope) {
		scope.SetColumn("uuid", uuid)
		scope.SetColumn("saas_id", saasID)
		scope.SetColumn("company_id", companyId)
	})
	return db
}
