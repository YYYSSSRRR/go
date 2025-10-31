package repository

import (
	"fmt"
	"ginTest/internal/model"

	"gorm.io/gorm"
)

type UserRepo interface {
	// GetAllUsers 查询所有用户
	GetAllUsers() ([]*model.User, error)

	// GetUserByEmail 根据email查询用户
	GetUserByEmail(email string) (*model.User, error)

	CreateNewUser(user *model.User) error
}

// UserRepository 需要注入db操作数据库
// 当作UserMapper类
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository 创建Repository实例
func NewUserRepository(db *gorm.DB) UserRepo {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAllUsers() ([]*model.User, error) {
	var results []*model.User
	err := r.db.Find(&results).Error
	if err != nil {
		return nil, fmt.Errorf("查询用户错误:%v", err)
	}
	return results, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	result := &model.User{}

	dbResult := r.db.Where("email=?", email).First(result)
	if dbResult.Error != nil {
		return nil, fmt.Errorf("查询用户错误:%v", dbResult.Error)
	}
	return result, nil
}

func (r *UserRepository) CreateNewUser(user *model.User) error {
	err := r.db.Create(user)
	if err != nil {
		return fmt.Errorf("注册用户失败：%v", err)
	}
	return nil
}
