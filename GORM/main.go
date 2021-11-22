package main

import (
	"fmt"
	_ "gorm/db"
	"gorm/db/model"
	config "gorm/environment"
)

func main() {
	fmt.Println("Server Started at " + config.PORT)

	user := model.CreateUser(&model.User{
		Name:    "Kadir",
		Surname: "Deniz",
		Email:   "kadirdenz@hotmail.com",
	})
	fmt.Println(user)

	find := model.FindUserById(int(user.ID))
	fmt.Println(find)

	deleted := model.DeleteUser(int(user.ID))
	fmt.Println(deleted)

}
