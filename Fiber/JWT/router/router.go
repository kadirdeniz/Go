package router

import (
	"github.com/fiber/jwt/controller"
	"github.com/fiber/jwt/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupApi(app *fiber.App) {
	api := app.Group("/api", middleware.IsTokenExists)
	api.Get("/user/:userId", controller.Get)
	api.Get("/allusers", controller.GetAll)
	api.Put("/user/:userId", controller.Update)
	api.Delete("/user/:userId", controller.Delete)
}

func SetupAuth(app *fiber.App) {
	api := app.Group("/auth")
	api.Post("/login", controller.Login)
	api.Post("/register", controller.Register)
}
