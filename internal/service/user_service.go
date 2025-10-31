package service

import (
	"ginTest/internal/model"
	"ginTest/internal/repository"
)

// UserService 接口
type UserService interface {
	GetAllUsers() ([]*model.User, error)
}
type UserServiceImpl struct {
	userRepo repository.UserRepo
}

// NewUserService 依赖注入接口实例
func NewUserService(userRepo repository.UserRepo) UserService {
	return &UserServiceImpl{userRepo: userRepo}
}

// GetAllUsers 实现接口的方法
func (us *UserServiceImpl) GetAllUsers() ([]*model.User, error) {
	return us.userRepo.GetAllUsers()
}
