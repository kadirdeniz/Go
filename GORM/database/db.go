package database

import (
	"fmt"

	config "github.com/gorm/package"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		fmt.Println("Db Connection Error :", err)
	}
	fmt.Println("Db Connection Success")
	DB = db
}
