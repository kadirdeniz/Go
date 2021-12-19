package main

import (
	"github.com/fiber/advanced/gorm/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router.SetupApi(app)

	app.Listen(":3000")
}
