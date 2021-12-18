package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func IsTokenExists(c *fiber.Ctx) error {
	header := c.Get("Authorization")
	if header == "" {
		return c.JSON(fiber.Map{
			"status": false,
			"error":  "Token Not Provided",
		})
	}
	return c.Next()
}
