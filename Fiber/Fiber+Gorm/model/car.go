package model

import (
	"github.com/fiber/gorm/datasource"
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	Price  int
	Brand  string
	CModel string
	Year   int
	KM     int
	Color  string
}

func (c *Car) Create() *gorm.DB {
	return datasource.PostgreDB.Create(&c)
}

func (c *Car) Delete() *gorm.DB {
	return datasource.PostgreDB.Delete(&c)
}

func (c *Car) Update() *gorm.DB {
	return datasource.PostgreDB.Model(&c).Updates(&c)
}

func (c *Car) Get() *gorm.DB {
	return datasource.PostgreDB.First(&c)
}
