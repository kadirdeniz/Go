package db

import (
	"fmt"
	config "gorm/environment"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(postgres.Open(config.ConnectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db

	fmt.Println("Database Connection Success")
}
