package middleware

import (
	"errors"
	"ginTest/internal/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// 测试鉴权的中间件
func TestAuthToken(t *testing.T) {
	TokenParser = func(token string) (*utils.UserInfo, error) {
		if token == "valid_token" {
			return &utils.UserInfo{Nickname: "mockUser"}, nil
		} else {
			return nil, errors.New("invalid token")
		}

	}
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name  string
		token string

		expectedStatus int
	}{
		{
			name:  "valid token",
			token: "valid_token",

			expectedStatus: http.StatusOK,
		},
		{
			name:  "invalid token",
			token: "bad_token",

			expectedStatus: http.StatusForbidden,
		},
		{
			name:  "empty token",
			token: "",

			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()
			r.Use(VerifyToken)
			r.GET("/test", func(c *gin.Context) {
				c.JSON(200, gin.H{"msg": "ok"})
			})

			req, _ := http.NewRequest("GET", "/test", nil)
			req.Header.Set("Authorization", tt.token)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatus, w.Code)

		})
	}
}
