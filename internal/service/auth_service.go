package service

import (
	"fmt"
	"ginTest/internal/model"
	"ginTest/internal/repository"
	"ginTest/internal/utils"
	"strconv"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type AuthService interface {
	SendCode(email string, ctx *gin.Context) (string, error)
	Login(email string, code string, ctx *gin.Context) (*model.User, string, error)
}

type AuthServiceImpl struct {
	repo  repository.UserRepo
	redis *redis.Client
}

func NewAuthService(repo repository.UserRepo, redis *redis.Client) AuthService {
	return &AuthServiceImpl{
		repo:  repo,
		redis: redis,
	}
}

func (authService *AuthServiceImpl) SendCode(email string, ctx *gin.Context) (string, error) {
	//生成验证码
	code := utils.GenerateCode()

	//将验证码存到redis中，设置过期时间5分钟
	authService.redis.Set(ctx, utils.LOGIN_CODE+email, code, time.Minute*5)

	return code, nil
}

func (authService *AuthServiceImpl) Login(email string, code string, ctx *gin.Context) (*model.User, string, error) {
	var err error
	//从redis中获取验证码，验证
	codeFromRedis, err := authService.redis.Get(ctx, utils.LOGIN_CODE+email).Result()
	if err != nil || codeFromRedis != code {
		return nil, "", fmt.Errorf("验证码错误或过期")
	}

	//验证成功，判断是否存在该用户
	user, err := authService.repo.GetUserByEmail(email)
	if err != nil {
		return nil, "", err
	}

	//如果存在就直接返回登录的token
	if user != nil {
		token, err := utils.GenerateToken(strconv.Itoa(user.ID), user.Email, user.NickName)
		if err != nil {
			return nil, "", err
		}
		return user, token, nil
	}

	//如果没有就先创建用户（自动注册）
	nickname := uuid.New().String()
	newUser := &model.User{
		Email:    email,
		NickName: nickname,
	}

	err = authService.repo.CreateNewUser(newUser)
	if err != nil {
		return nil, "", err
	}
	token, err := utils.GenerateToken(strconv.Itoa(newUser.ID), newUser.Email, newUser.NickName)
	if err != nil {
		return nil, "", err
	}
	return newUser, token, nil
}
