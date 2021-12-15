package main

import (
	"github.com/fiber/gorm/datasource"
	"github.com/fiber/gorm/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	datasource.ConnectDB()
	router.SetupRoutes(app)
	app.Listen(":3000")
}
