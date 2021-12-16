package router

import (
	"github.com/fiber/gorm/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/car", controller.Create)
	api.Put("/car/:carId", controller.Update)
	api.Delete("/car/:carId", controller.Delete)
	api.Get("/car/:carId", controller.GetById)
	api.Get("/allcars", controller.GetAll)

}
