package router

import (
	"github.com/docker/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/anasayfa", controller.Anasayfa)
}
