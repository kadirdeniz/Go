package router

import "github.com/gofiber/fiber/v2"

func SetupApi(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/user/:userId", handler)
	// api.Get("/allusers")
	// api.Put("/user/:userId")
	// api.Delete("/user/:userId")
}

func handler(c *fiber.Ctx) error {
	return c.SendString("kadir")
}
