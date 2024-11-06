package domain

import (
	"ecommerce-api/utils"

	"gorm.io/gorm"
)

// User represents the user model
type User struct {
    ID       uint      `gorm:"primaryKey"`
    Email    string    `gorm:"unique;not null"`
    Password string    `gorm:"not null"`
    IsAdmin  bool       `gorm:"not null"`
    gorm.Model
}

type UserResponse struct {
    ID    uint   `json:"id"`
    Email string `json:"email"`
}

func (user *User) Validate() error {
	return utils.Validate.Struct(user)
}