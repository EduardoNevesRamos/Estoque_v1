package models

import (
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Name        string  `gorm:"column:name" json:"name"`
	Price       float64 `gorm:"column:price" json:"price"`
	Description string  `gorm:"column:description" json:"description"`
	// ImageURL    string  `gorm:"column:imageurl" json:"img_url"`
}

type Client struct {
	gorm.Model
	Name   string  `gorm:"column:name" json:"name"`
	Age    uint    `gorm:"column:age" json:"age"`
	Salary float64 `gorm:"column:salary" json:"salary"`
}
