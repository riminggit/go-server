package logMiddleware

import (
	"github.com/gin-gonic/gin"
	"go-server-template/pkg/log"
)

func LogerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logging.WriteMiddlewareLogs(c)
	}
}

// 日志记录到 MongoDB
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
