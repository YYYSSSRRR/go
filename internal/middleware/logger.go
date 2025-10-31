package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	Logger, _ = zap.NewDevelopment()
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		//执行请求
		c.Next()

		//请求执行的时间
		latency := time.Since(start)
		//响应的状态码
		status := c.Writer.Status()

		Logger.Info("请求日志",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", status),
			zap.Duration("latency", latency),
			zap.String("client_ip", c.ClientIP()),
		)

	}
}
