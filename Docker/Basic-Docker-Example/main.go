package main

import (
	"github.com/docker/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router.Setup(app)
	app.Listen(":3000")
}
