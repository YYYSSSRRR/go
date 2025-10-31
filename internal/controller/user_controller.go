package controller

import (
	"ginTest/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetAllUsers(c *gin.Context)
}

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{userService: userService}
}

func (uc *UserControllerImpl) GetAllUsers(c *gin.Context) {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get users",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
