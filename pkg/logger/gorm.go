package logger

import (
	"github.com/sirupsen/logrus"
)

// gorm日志记录
type GormLogger struct {
	TraceId string
}

func (g *GormLogger) Print(v ...interface{}) {
	logger := NewLogger()
	switch v[0] {
	case "sql":
		entry := logger.WithFields(
			logrus.Fields{
				"module":        "gorm",
				"type":          "sql-exec",
				"rows_returned": v[5],
				"src":           v[1],
				"values":        v[4],
				"sql":           v[3],
				"duration":      v[2],
				"trace_id":      g.TraceId,
			},
		)
		entry.Message = "sql-log"
		s, _ := entry.String()
		ToChan(SqlChannel, s)
	case "log":
		entry := logger.WithFields(
			logrus.Fields{
				"module":     "gorm",
				"type":       "sql-log",
				"src":        v[1],
				"error-info": v[2],
				"trace_id":   g.TraceId,
			},
		)
		entry.Message = "sql-log"
		s, _ := entry.String()
		ToChan(SqlChannel, s)
	case "info":
		entry := logger.WithFields(
			logrus.Fields{
				"module":   "gorm",
				"type":     "sql-info",
				"src":      v[1],
				"trace_id": g.TraceId,
			},
		)
		entry.Message = "sql-log"
		s, _ := entry.String()
		ToChan(SqlChannel, s)
	}
}
