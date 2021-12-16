package datasource

import (
	"fmt"

	"github.com/fiber/gorm/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db DB = DB{
	Name:     "PostgreSql",
	Host:     "localhost",
	Port:     5432,
	Db:       "golang",
	User:     "root",
	Password: "123456789",
}

var PostgreDB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(postgres.Open(db.DSNConverter()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("DB Connection Success")
	PostgreDB = db

	db.AutoMigrate(&model.Car{})
}
