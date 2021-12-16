package controller

import (
	"github.com/fiber/gorm/datasource"
	"github.com/fiber/gorm/model"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	car := model.Car{}
	c.BodyParser(&car)
	response := datasource.PostgreDB.Create(&car)
	if response.Error != nil {
		return c.JSON(fiber.Map{
			"status": false,
			"error":  response.Error,
		})
	}
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Car Intance Created",
	})
}

func Update(c *fiber.Ctx) error {
	carId, _ := c.ParamsInt("carId")
	car := model.Car{}
	car.ID = uint(carId)

	carInstance := datasource.PostgreDB.First(&car)
	if carInstance.RowsAffected == 0 {
		return c.JSON(fiber.Map{
			"status": false,
			"error":  "Instance Couldn't Found",
		})
	}
	c.BodyParser(&car)
	response := datasource.PostgreDB.Save(&car)
	if response.Error != nil {
		return c.JSON(fiber.Map{
			"status": false,
			"error":  response.Error,
		})
	}
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Car Intance Updated",
	})
}

func Delete(c *fiber.Ctx) error {
	carId, _ := c.ParamsInt("carId")
	car := model.Car{}
	car.ID = uint(carId)

	response := datasource.PostgreDB.Delete(&car)
	if response.RowsAffected == 0 {
		return c.JSON(fiber.Map{
			"status": false,
			"error":  "Instance Couldn't Found",
		})
	}
	if response.Error != nil {
		return c.JSON(fiber.Map{
			"status": false,
			"error":  response.Error,
		})
	}
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Car Intance Deleted",
	})
}

func GetById(c *fiber.Ctx) error {
	carId, _ := c.ParamsInt("carId")
	car := model.Car{}
	car.ID = uint(carId)

	response := datasource.PostgreDB.First(&car)
	if response.Error != nil {
		return c.JSON(fiber.Map{
			"status": false,
			"error":  response.Error,
		})
	}
	return c.JSON(fiber.Map{
		"status": true,
		"data":   car,
	})
}

func GetAll(c *fiber.Ctx) error {
	car := []model.Car{}

	response := datasource.PostgreDB.Find(&car)
	if response.Error != nil {
		return c.JSON(fiber.Map{
			"status": false,
			"error":  response.Error,
		})
	}
	return c.JSON(fiber.Map{
		"status": true,
		"data":   &car,
	})
}
