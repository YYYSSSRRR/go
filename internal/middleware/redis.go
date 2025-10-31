package middleware

import (
	"ginTest/config"

	"github.com/redis/go-redis/v9"
)

func NewRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.Addr,
		Password: config.Conf.Redis.Password,
		DB:       config.Conf.Redis.DB,
	})
}
