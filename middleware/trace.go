package middleware

import (
	"jcfw.com/legal-api/util"

	"github.com/gin-gonic/gin"
)

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 优先从请求头中获取请求ID，如果没有则使用UUID
		traceID := c.GetHeader("X-Request-Id")
		if traceID == "" {
			traceID = util.GenUUID()
		}
		util.SetTraceId(c, traceID)
		c.Next()
	}
}
