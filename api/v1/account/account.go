package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"jcfw.com/legal-api/model"
	"jcfw.com/legal-api/service"
)

//添加一个账户
func AddAcount(c *gin.Context) {
	var addJson model.LegalCompanyAccount
	if err := c.ShouldBindJSON(&addJson); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
	}
	res := service.NewService(c).AddAcount(&addJson, "123123")
	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}
