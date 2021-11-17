package userInt

import (
	"interface/Interface/carInt"
	"interface/Struct/car"
	"interface/Struct/user"
)

type Emp int

type UserInt interface {
	getUserById(arr []user.User, id int) user.User
	buyCar(id int) string
}

func (e Emp) GetUserById(arr []user.User) user.User {
	id := e
	for index, value := range arr {
		if index == id {
			return value
		}
	}
	panic("User Couldnt Found")
}

func GetUserNameById(arr []user.User, id int) string {
	for index, value := range arr {
		if index == id {
			return value.Name
		}
	}
	panic("User Couldnt Found")
}

func (e Emp) BuyCar(arr []user.User, carId int) string {
	id := e
	carInt.GetCarById(car.Cars, carId)
	return "{status:true,message:'" + GetUserNameById(user.User, id) + "'}"
}
