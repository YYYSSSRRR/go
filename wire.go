//go:build wireinject
// +build wireinject

package main

import (
	"ginTest/internal/controller"
	"ginTest/internal/repository"
	"ginTest/internal/routes"
	"ginTest/internal/service"

	"github.com/google/wire"
)

func InitializeApp() *routes.App {
	wire.Build(
		routes.ProviderSet,
		controller.ProviderSet,
		service.ProviderSet,
		repository.ProviderSet,
	)
	return nil
}
