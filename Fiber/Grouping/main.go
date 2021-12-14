package main

import (
	"github.com/fiber/grouping/router"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	router.ApiRouter(app)
	app.Listen(":3000")
}
