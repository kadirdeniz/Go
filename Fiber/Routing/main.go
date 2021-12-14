package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// Basic Routing Get for /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ana Sayfa")
	})

	// Parameters
	app.Get("/user/:name/:surname", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Hello %v %v", c.Params("name"), c.Params("surname")))
	})
	// Split With "."
	app.Get("/date/:day:.month:.year", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Date : %v / %v / %v ", c.Params("day"), c.Params("month"), c.Params("year")))
	})

	// Optional Parameters
	app.Get("/car/:model/:year?", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Model : %v , Year : %v", c.Params("model"), c.Params("year")))
	})

	// * Can Be Empy Optional Tag
	app.Get("/schools/*", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("*"))
	})

	// + Must Fullfilled
	app.Get("/school/name/+", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("School Name : %v", c.Params("+")))
	})

	// Using : At Url
	app.Get("/league::name", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("name"))
	})

	app.Listen(":3000")
}
