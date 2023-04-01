//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"go-fiber-demo/routes"
)

func InitializeApiRoutes() (routes.ApiRoutes, error) {
	wire.Build(routes.ApiRoutesSet)
	return routes.ApiRoutes{}, nil
}
