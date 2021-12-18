package controller

import (
	"github.com/fiber/jwt/database"
	"github.com/fiber/jwt/model"
	"github.com/gofiber/fiber/v2"
)

func Get(c *fiber.Ctx) error {
	userId, _ := c.ParamsInt("userId")
	user := model.User{}
	user.ID = uint(userId)
	response := database.DB.First(&user)

	if response.RowsAffected == 0 {
		return c.JSON(fiber.Map{
			"status": false,
			"data":   user,
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"data":   user,
	})
}

func Delete(c *fiber.Ctx) error {
	userId, _ := c.ParamsInt("userId")
	user := model.User{}
	user.ID = uint(userId)
	response := database.DB.Delete(&user)
	if response.Error != nil {
		c.JSON(fiber.Map{
			"status": false,
			"error":  response.Error,
		})
	}
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "User Deleted",
	})
}

func Update(c *fiber.Ctx) error {
	user := model.User{}
	userId, _ := c.ParamsInt("userId")
	user.ID = uint(userId)
	get := database.DB.First(&user)
	if get.RowsAffected == 0 {
		c.JSON(fiber.Map{
			"status": true,
			"error":  "User Couldn't Found",
		})
	}
	c.BodyParser(&user)
	response := database.DB.Save(&user)
	if response.Error != nil {
		return c.JSON(fiber.Map{
			"status": false,
			"error":  response.Error,
		})
	}
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "User Updated",
	})
}

func GetAll(c *fiber.Ctx) error {
	users := []model.User{}
	response := database.DB.Find(&users)
	if response.Error != nil {
		c.JSON(fiber.Map{
			"status": false,
			"error":  response.Error,
		})
	}
	return c.JSON(fiber.Map{
		"status": true,
		"data":   users,
	})
}
