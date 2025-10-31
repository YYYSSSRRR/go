package routes

import (
	"ginTest/internal/controller"

	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	userController controller.UserController
	authController controller.AuthController
}

func NewUserRoute(userController controller.UserController, authController controller.AuthController) *UserRoute {
	return &UserRoute{
		userController: userController,
		authController: authController,
	}
}

func (u *UserRoute) NewUserGroup(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.GET("getAll", u.userController.GetAllUsers)
		user.POST("/code", u.authController.SendCode)
		user.POST("/login", u.authController.Login)
	}

}
