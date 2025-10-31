package middleware

import (
	"ginTest/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TokenParser 设置为全局变量，便于在测试代码中可以改掉它
var TokenParser = utils.ParseToken

// VerifyToken 登录拦截器
func VerifyToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "400",
			"message": "没有携带token",
		})
		return
	}
	userInfo, err := TokenParser(token)
	if err != nil || userInfo == nil {
		c.JSON(http.StatusForbidden, map[string]interface{}{
			"code":    "403",
			"message": "身份验证失败",
		})
		return
	}
	c.Next()
}
