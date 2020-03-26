package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"jcfw.com/legal-api/model"
)

type BaseService struct {
	DB *gorm.DB
	C  *gin.Context
}

//实例化一个新的service
func NewService(c *gin.Context) *BaseService {
	return &BaseService{
		DB: model.GetDB(c),
		C:  c,
	}
}
