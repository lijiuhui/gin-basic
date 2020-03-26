package util

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	prefix            = "legal"
	RequestTraceIdKey = prefix + "trace-id"
)

//获取request trace id
func GetTraceId(c *gin.Context) string {
	return c.GetString(RequestTraceIdKey)
}

//设置 request trace id
func SetTraceId(c *gin.Context, v string) {
	c.Set(RequestTraceIdKey, v)
}

//获取一个uuid
func GenUUID() string {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
	}
	return u.String()
}
