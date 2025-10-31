package repository

import (
	"fmt"
	"ginTest/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化mysql连接配置
func NewDB() *gorm.DB {
	var err error

	dsn := config.Conf.Database.Source

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	fmt.Println("数据库连接成功")
	return db
}
