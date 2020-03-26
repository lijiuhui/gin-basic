package middleware

import (
	"net/http"
	"strings"

	"jcfw.com/legal-api/util"
	"jcfw.com/legal-api/util/e"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//JWT验证
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			code int
			data interface{}
		)
		Authorization := c.GetHeader("Authorization")
		token := strings.Split(Authorization, " ")

		if Authorization == "" {
			code = e.INVALID_PARAMS
		} else {
			cliams, err := util.ParseToken(token[1])
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
			//设置
			c.Set(util.UUIDKey, cliams.UUID)
			c.Set(util.SaasIDKey, cliams.SaasId)
			c.Set(util.CompanyIDKey, cliams.CompanyId)
			c.Set(util.PhoneKey, cliams.Phone)
			c.Set(util.NameKey, cliams.Name)
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
