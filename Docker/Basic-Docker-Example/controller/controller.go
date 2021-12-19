package controller

import "github.com/gofiber/fiber/v2"

func Anasayfa(c *fiber.Ctx) error {
	return c.SendString("Anasayfa")
}
