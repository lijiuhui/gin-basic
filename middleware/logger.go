package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"

	"jcfw.com/legal-api/pkg/logger"
	"jcfw.com/legal-api/util/response"

	"jcfw.com/legal-api/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

//日志中间件
func HttpLogger() gin.HandlerFunc {
	go logger.HandleAccessChannel()
	ruslog := logger.NewLogger()
	return func(c *gin.Context) {
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		// 开始时间
		startTime := util.GetCurrentMilliUnix()
		responseBody := bodyLogWriter.body.String()

		var responseCode int
		var responseMsg string
		var responseData interface{}

		if responseBody != "" {
			res := response.Response{}
			err := json.Unmarshal([]byte(responseBody), &res)
			if err == nil {
				responseCode = res.Code
				responseMsg = res.Message
				responseData = res.Data
			}
		}
		// 结束时间
		endTime := util.GetCurrentMilliUnix()
		// if c.Request.Method == "POST" {
		// 	_ = c.Request.ParseForm()
		// }
		str, _ := ruslog.WithFields(logrus.Fields{
			"trace_id":          util.GetTraceId(c),
			"request_time":      startTime,
			"request_method":    c.Request.RequestURI,
			"request_uri":       c.Request.RequestURI,
			"request_proto":     c.Request.Proto,
			"request_ua":        c.Request.UserAgent(),
			"request_referer":   c.Request.Referer(),
			"request_post_data": c.Request.PostForm.Encode(),
			"request_client_ip": c.ClientIP(),

			"response_time": endTime,
			"response_code": responseCode,
			"response_msg":  responseMsg,
			"response_data": responseData,

			"cost_time": fmt.Sprintf("%vms", endTime-startTime),
		}).String()
		logger.ToChan(logger.AccessChannel, str)
		// 处理请求
		c.Next()
	}
}
