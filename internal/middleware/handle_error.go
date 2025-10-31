package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				Logger.Error("捕获到panic错误",
					zap.Any("error", err),
					zap.ByteString("stack", debug.Stack()),
					zap.String("path", c.Request.URL.Path),
					zap.String("method", c.Request.Method),
				)

				c.JSON(http.StatusInternalServerError, gin.H{
					"code": 500,
					"msg":  "服务器内部错误",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
