package model

import "time"

type User struct {
	ID        int    `gorm:"column:id;primaryKey"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
	NickName  string
	CreatedAt time.Time `gorm:"column:create_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
}

// 如果要指定表名的话
func (user *User) TableName() string {
	return "users"
}

//这是绑定到User结构体上的方法
