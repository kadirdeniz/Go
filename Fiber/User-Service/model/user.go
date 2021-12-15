package model

import "fmt"

var Users []User

type User struct {
	Id       int
	Name     string
	Surname  string
	Email    string
	Password string
}

func (u *User) AddId() *User {
	u.Id = len(Users) + 1
	return u
}

func (u *User) Create() []User {
	u.AddId()
	Users = append(Users, *u)
	return Users
}

func RemoveIndex(s []User, index int) []User {
	return append(s[:index-1], s[index:]...)
}

func (u *User) Delete() (int, string) {
	userId := u.Id
	status, _ := u.GetById()
	if status == 0 {
		return 0, "Kullanıcı Bulunamadı"
	}
	Users = RemoveIndex(Users, userId)
	fmt.Println(Users)
	return 1, "Kullanıcı Silindi"
}

func (u *User) GetById() (int, User) {

	for _, user := range Users {
		if user.Id == u.Id {
			return 1, user
		}
	}

	return 0, User{}
}

func (u *User) GetByEmail() (int, User) {
	for _, user := range Users {
		if user.Email == u.Email {
			return 1, user
		}
	}

	return 0, User{}
}

func (u *User) CheckPassword(password string) bool {
	return u.Password == password
}
