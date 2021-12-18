package main

import (
	"github.com/fiber/jwt/database"
	"github.com/fiber/jwt/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	router.SetupApi(app)
	router.SetupAuth(app)

	app.Listen(":3000")
}
