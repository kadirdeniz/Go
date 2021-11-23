package user

import (
	"fmt"

	db "github.com/gorm/database"
	_"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Name    string
	Surname string
	Email   string
}

func init() {
	db.DB.AutoMigrate(&User{})
}

func (u *User) Create() string {
	db.DB.Create(&u)

	return "{status:true,message:'User Created'}"
}

func (u *User) Delete() string {
	db.DB.Delete(&User{}, u.ID)

	return "{status:true,message:'User Deleted'}"
}

func (u *User) Get() string {
	result := db.DB.First(&u)
	fmt.Println(result)
	return "User Found"
}
