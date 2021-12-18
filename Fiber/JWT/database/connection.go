package database

import (
	"fmt"

	"github.com/fiber/jwt/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=root password=123456789 dbname=golang port=5432 sslmode=disable TimeZone=Asia/Istanbul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Db Connection Success")
	DB = db
	db.AutoMigrate(&model.User{})
}
