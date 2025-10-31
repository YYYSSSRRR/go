package controller

import (
	"ginTest/internal/model"
	"ginTest/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	SendCode(c *gin.Context)
	Login(c *gin.Context)
}
type AuthControllerImpl struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		authService: authService,
	}
}

func (authController *AuthControllerImpl) SendCode(c *gin.Context) {
	//var email model.EmailAddress
	//err := c.BindJSON(&email)
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, model.Failure("输入无效，请重试"))
		return
	}
	code, err := authController.authService.SendCode(email, c)
	if err != nil {
		c.JSON(http.StatusForbidden, model.Failure(err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.Success(map[string]string{
		"code": code,
	}))
}

func (authController *AuthControllerImpl) Login(ctx *gin.Context) {
	//var loginParam model.LoginParam
	//err := ctx.BindJSON(&loginParam)
	email := ctx.Query("email")
	code := ctx.Query("code")
	if email == "" || code == "" {
		ctx.JSON(http.StatusBadRequest, model.Failure("输入无效，请重试"))
		return
	}
	user, token, err := authController.authService.Login(email, code, ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, model.Failure(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(
		map[string]interface{}{
			"userInfo": user,
			"token":    token,
		},
	))

}
