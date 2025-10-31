package routes

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewApp,
	NewUserRoute,
)
