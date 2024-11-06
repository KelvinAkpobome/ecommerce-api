package ports

import (
	"ecommerce-api/domain"
)

type UserRepository interface {
    CreateUser(user *domain.User) (*domain.UserResponse, error)
    GetUserByEmail(email string) (*domain.User, error)
    GetUserByID(id uint) (*domain.User, error)
}
