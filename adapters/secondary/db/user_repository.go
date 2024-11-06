package db

import (
    "ecommerce-api/domain"
    "ecommerce-api/ports"
    "gorm.io/gorm"
    "fmt"
    "strings"
)

type UserRepositoryDB struct {
    DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) ports.UserRepository {
    return &UserRepositoryDB{DB: db}
}

func (r *UserRepositoryDB) CreateUser(user *domain.User) (*domain.UserResponse, error) {
    if err := r.DB.Create(user).Error; err != nil {
        if strings.Contains(err.Error(), "unique constraint") {
            return nil,fmt.Errorf("User with this email already exists")
        }
        return nil, err
    }

    return &domain.UserResponse{
        ID:    user.ID,
        Email: user.Email,
    }, nil
}

func (r *UserRepositoryDB) GetUserByEmail(email string) (*domain.User, error) {
    var user domain.User
    if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepositoryDB) GetUserByID(id uint) (*domain.User, error) {
    var user domain.User
    if err := r.DB.First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
