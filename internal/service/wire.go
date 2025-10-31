package service

import (
	"ginTest/internal/middleware"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewUserService,
	NewAuthService,
	middleware.NewRedis,
)
