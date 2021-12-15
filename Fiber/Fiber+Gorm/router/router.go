package router

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/car/:carId", func(c *fiber.Ctx) error {
		return c.SendString("kdei<")
	})
	api.Post("/car", func(c *fiber.Ctx) error {
		return c.SendString("kdei<")
	})
	api.Delete("/car/:carId", func(c *fiber.Ctx) error {
		return c.SendString("kdei<")
	})
	api.Put("/car/:carId", func(c *fiber.Ctx) error {
		return c.SendString("kdei<")
	})
	api.Get("/allcars", func(c *fiber.Ctx) error {
		return c.SendString("kdei<")
	})

}
