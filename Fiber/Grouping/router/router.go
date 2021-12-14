package router

import "github.com/gofiber/fiber"

func ApiRouter(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/user", func(c *fiber.Ctx) {
		c.SendString("Grouping")
	})

	api.Get("/user/name", func(c *fiber.Ctx) {
		c.SendString("Grouping")
	})

	api.Get("/user/passwor", func(c *fiber.Ctx) {
		c.SendString("Grouping")
	})

}
