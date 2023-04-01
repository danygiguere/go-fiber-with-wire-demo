package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func main() {
	_, err := InitializeApiRoutes()
	//routes.load()
	if err != nil {
		fmt.Printf("failed to create apiRoutes: %s\n", err)
		os.Exit(2)
	}
	app := fiber.New()
	err = app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatal("Failed to init app. \n", err)
	}
}
