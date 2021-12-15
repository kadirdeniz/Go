package controller

import (
	"github.com/fiber/user/model"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	user := model.User{}
	c.BodyParser(&user)
	Users := user.Create()
	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "Kullanıcı Oluşturuldu",
		"User":    Users,
	})
}

func Get(c *fiber.Ctx) error {
	userId, _ := c.ParamsInt("userId")

	user := model.User{
		Id: userId,
	}

	status, user := user.GetById()
	if status == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status": false,
			"error":  "Kullanıcı Bulunamadı",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": true,
		"User":   user,
	})
}

func GetAll(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"users": model.Users,
	})
}

func Delete(c *fiber.Ctx) error {
	userId, _ := c.ParamsInt("userId")
	user := model.User{
		Id: userId,
	}
	status, response := user.Delete()
	if status == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status": false,
			"error":  response,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "Kullanıcı Silindi",
	})
}

func Login(c *fiber.Ctx) error {
	user := model.User{}
	c.BodyParser(&user)
	isEmailExists, dbUser := user.GetByEmail()

	if isEmailExists == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status": false,
			"error":  "Email Adresi Bulunamadı",
		})
	}

	isMatch := user.CheckPassword(dbUser.Password)
	if !isMatch {
		return c.Status(400).JSON(fiber.Map{
			"status": false,
			"error":  "Şifre Hatalı",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "Giriş Başarılı",
	})
}
