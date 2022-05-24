package entity

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Storage int64   `json:"storage"`
}

func (Products) TableName() string {
	return "Products"
}

func NewProduct(name string, price float64, storage int64) *Products {
	return &Products{
		Name:    name,
		Price:   price,
		Storage: storage,
	}
}
