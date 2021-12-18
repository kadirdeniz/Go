package controller

import (
	"github.com/fiber/jwt/database"
	"github.com/fiber/jwt/model"
	"github.com/fiber/jwt/token"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	user := model.User{}
	c.BodyParser(&user)
	user.Hash()
	response := database.DB.Create(&user)
	if response.Error != nil {
		return c.JSON(fiber.Map{
			"status": false,
			"error":  response.Error,
		})
	}

	token, err := token.Create(user.ID)
	if err != nil {
		return c.JSON(fiber.Map{
			"status": false,
			"error":  err,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Registration Success",
		"token":   token,
	})
}

func Login(c *fiber.Ctx) error {
	user := model.User{}
	c.BodyParser(&user)
	password := user.Password
	if user.Email == "" || user.Password == "" {
		return c.JSON(fiber.Map{
			"status": false,
			"error":  "Email and Password Fields Must Be Filled",
		})
	}

	isEmailExists := database.DB.Where("email", user.Email).First(&user)
	if isEmailExists.RowsAffected == 0 {
		return c.JSON(fiber.Map{
			"status": false,
			"error":  "Provide a Valid Email",
		})
	}
	isMatch := user.Compare(password)
	if !isMatch {
		return c.JSON(fiber.Map{
			"status": false,
			"error":  "Wrong Password!",
		})
	}

	token, err := token.Create(user.ID)
	if err != nil {
		return c.JSON(fiber.Map{
			"status": false,
			"error":  err,
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"error":  "Logged In",
		"token":  token,
	})
}
