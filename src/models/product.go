package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Gender  string
	Style   string
	Color   string
	Pattern string
	Figure  string
	Size    string
	Price   float64
}
