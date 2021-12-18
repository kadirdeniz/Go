package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Surname  string
	Email    string
	Password string
}

var Cost int = 14

func (u *User) Hash() {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), Cost)
	if err != nil {
		panic(err)
	}
	u.Password = string(bytes)
}

func (u *User) Compare(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
