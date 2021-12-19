package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string
	Surname string
	Email   string
	Social  Social
	Images  []Images
}

type Social struct {
	UserID    uint
	Twitter   string
	Linkedin  string
	Instagram string
	Facebook  string
	Github    string
}

type Images struct {
	gorm.Model
	UserID uint
	Image  string
}
