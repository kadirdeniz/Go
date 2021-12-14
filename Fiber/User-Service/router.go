package main

import "github.com/gofiber/fiber/v2"

func SetupApi(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create")
}
