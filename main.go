package main

import (
	"fmt"
	"ginTest/config"
)

func main() {
	var err error
	//加载配置文件
	err = config.LoadConfig()
	if err != nil {
		fmt.Errorf("配置文件加载错误,%v", err)
	}
	app := InitializeApp()
	err = app.Run(fmt.Sprintf(":%d", config.Conf.App.Port))
	if err != nil {
		fmt.Errorf("服务器启动错误:%v", err)
		return
	}
}
