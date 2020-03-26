package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"jcfw.com/legal-api/service"
	"jcfw.com/legal-api/util"
)

type loginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

//登录
func Login(c *gin.Context) {
	var loginForm loginForm
	if err := c.BindJSON(&loginForm); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	//验证参数
	data, err := service.NewService(c).LoginByPhoneAndPassword(loginForm.Username, loginForm.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	//生成jwt token
	token, err := util.GenerateToken(data.UUID, data.Phone, data.Name, data.SaasID, data.CompanyID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data": map[string]interface{}{
			"token":  token,
			"expire": viper.GetInt("jwt.expire"),
		},
	})
}
