package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-demo/services"
)

type DemoController struct {
	baseController BaseController
	i18n           services.I18n
}

func NewDemoController(baseController BaseController) DemoController {
	controller := BaseController{i18n: baseController.i18n}
	return DemoController{i18n: controller.i18n}
}

func (controller *DemoController) Index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    "Hello World",
	})
}

func (controller *DemoController) I18nHelloWorld(ctx *fiber.Ctx) error {
	resp := controller.i18n.Localize(ctx, "helloWorld", nil)
	return ctx.SendString(resp)
}

func (controller *DemoController) I18nHelloName(ctx *fiber.Ctx) error {
	templateData := map[string]string{"Name": ctx.Params("name")}
	return ctx.SendString(controller.i18n.Localize(ctx, "helloName", templateData))
}
