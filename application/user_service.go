package application

import (
	"ecommerce-api/domain"
	"ecommerce-api/ports"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{UserRepository: repo}
}

func (s *UserService) RegisterUser(email, password string) (*domain.UserResponse, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := domain.User{Email: email, Password: string(hashedPassword)}
	createdUser, err := s.UserRepository.CreateUser(&user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (s *UserService) LoginUser(email, password string) (*domain.User, error) {
	user, err := s.UserRepository.GetUserByEmail(email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

func (s *UserService) GetUserByID(id uint) (*domain.User, error) {
	user, err := s.UserRepository.GetUserByID(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}
	return user, nil
}
