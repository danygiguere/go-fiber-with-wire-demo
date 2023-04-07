package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-demo/controllers"
)

type ApiRoutes struct {
	route          fiber.Router
	baseController controllers.BaseController
}

func NewApiRoutes(route fiber.Router, baseController controllers.BaseController) ApiRoutes {
	return ApiRoutes{route: route, baseController: baseController}
}

func (apiRoutes ApiRoutes) Load() {

	demoController := controllers.NewDemoController(apiRoutes.baseController)
	api := apiRoutes.route
	api.Get("/demo", demoController.Index)
	api.Get("/demo/i18n", demoController.I18nHelloWorld)
	api.Get("/demo/i18n/:name", demoController.I18nHelloName)
}
