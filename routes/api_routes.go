package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

type ApiRoutes struct {
	route fiber.Router
}

func ProvideRoutes(route fiber.Router) ApiRoutes {
	return ApiRoutes{route: route}
}

var ApiRoutesSet = wire.NewSet(ProvideRoutes)

func (apiRoutes ApiRoutes) Load() {
	apiRoutes.route.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, 2023 World!")
	})
}
