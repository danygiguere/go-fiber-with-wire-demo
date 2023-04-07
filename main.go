package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-fiber-demo/configs"
	"go-fiber-demo/routes"
	"log"
	"os"
)

func StartApp() *fiber.App {
	app := fiber.New()
	baseController, err := InitializeBaseController()
	if err != nil {
		fmt.Printf("failed to create apiRoutes: %s\n", err)
		os.Exit(2)
	}

	routes.NewApiRoutes(app, baseController).Load()
	return app
}

func main() {
	configs.LoadEnvVariables()
	app := StartApp()
	err := app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatal("Failed to init app. \n", err)
	}
}
