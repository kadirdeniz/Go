package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Static("/image", "./public")

	app.Listen(":3000")
}
