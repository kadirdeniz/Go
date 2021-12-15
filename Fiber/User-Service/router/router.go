package router

import (
	"github.com/fiber/user/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupApi(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/register", controller.Register)
	api.Post("/login", controller.Login)
	api.Delete("/user/:userId", controller.Delete)
	api.Get("/user/:userId", controller.Get)
	api.Get("/userall", controller.GetAll)

}
