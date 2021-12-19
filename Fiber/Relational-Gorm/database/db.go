package database

import (
	"fmt"

	"github.com/fiber/advanced/gorm/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=root password=123456789 dbname=golang port=5432 sslmode=disable TimeZone=Asia/Istanbul",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	fmt.Println("DB Connection Succes")
	DB = db

	DB.AutoMigrate(&model.User{}, &model.Social{}, &model.Images{})
}
