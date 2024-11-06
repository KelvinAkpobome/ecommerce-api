package domain

import (
	"ecommerce-api/utils"
	"gorm.io/gorm"
)

type Order struct {
    ID       uint       `gorm:"primaryKey"`
    UserID   uint       `gorm:"not null;index"`
    User     User       `gorm:"foreignKey:UserID"`
    Products []Product  `gorm:"many2many:order_products;"`
    Status   utils.OrderStatus   `gorm:"not null"`
    gorm.Model
}

type OrderProduct struct {
    OrderID   uint `gorm:"primaryKey"`
    ProductID uint `gorm:"primaryKey"`
}


func (order *Order) Validate() error {
	return utils.Validate.Struct(order)
}