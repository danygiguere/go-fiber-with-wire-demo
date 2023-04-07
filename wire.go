//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"go-fiber-demo/controllers"
)

func InitializeBaseController() (controllers.BaseController, error) {
	wire.Build(controllers.BaseControllerSet)
	return controllers.BaseController{}, nil
}
