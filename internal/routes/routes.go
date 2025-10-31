package routes

import (
	"fmt"
	"ginTest/internal/middleware"

	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
	usv    *UserRoute
}

func NewApp(usv *UserRoute) *App {
	r := gin.New()
	r.Use(gin.Recovery())

	//注册日志中间件
	middleware.InitLogger()
	r.Use(middleware.LoggerMiddleware())

	//错误捕获中间件
	r.Use(middleware.RecoveryMiddleware())

	usv.NewUserGroup(r)
	return &App{
		router: r,
		usv:    usv,
	}
}

func (r *App) Run(port string) error {
	return r.router.Run(fmt.Sprintf("%s", port))
}
