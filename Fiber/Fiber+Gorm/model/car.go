package model

import (
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
