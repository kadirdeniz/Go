package model

import (
	"fmt"
	"gorm/db"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string
	Surname string
	Email   string
}

var database *gorm.DB = db.DB

func init() {
	database.AutoMigrate(&User{})
}

func CreateUser(u *User) *User {

	user := database.Create(u)

	if user.Error != nil {
		panic("User Couldnt Created")
	}

	return user
}

func DeleteUser(userId int) string {
	user := database.Delete(&User{}, userId)
	if user.Error != nil {
		panic("User Couldnt Deleted")
	}

	return "{status:true,message:'User Deleted'}"
}

func FindUserById(userId int) string {
	user := database.First(&User{}, userId)
	if user.Error != nil {
		fmt.Println(user.Error)
	}
	return "{status:true,data:" + user + "}"
}
