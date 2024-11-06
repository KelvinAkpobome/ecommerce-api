package domain

import (
	"ecommerce-api/utils"
	"gorm.io/gorm"
)

type Product struct {
	ID    uint    `json:"id" gorm:"primaryKey"`          
	Name  string  `json:"name" gorm:"not null" validate:"required,min=3,max=100"` 
	Price float64 `json:"price" gorm:"not null" validate:"required,gt=0"`
	gorm.Model
}

func (product *Product) Validate() error {
	return utils.Validate.Struct(product)
}
