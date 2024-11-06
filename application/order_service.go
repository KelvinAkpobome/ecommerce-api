package application

import (
	"ecommerce-api/domain"
	"ecommerce-api/ports"
	"ecommerce-api/utils"
)

type OrderService struct {
    OrderRepository ports.OrderRepository
}

func NewOrderService(repo ports.OrderRepository) *OrderService {
    return &OrderService{OrderRepository: repo}
}

func (s *OrderService) CreateOrder(order *domain.Order) error {
    return s.OrderRepository.CreateOrder(order)
}

func (s *OrderService) GetOrderByID(id uint) (*domain.Order, error) {
    return s.OrderRepository.GetOrderByID(id)
}

func (s *OrderService) GetOrdersByUserID(userID uint) ([]domain.Order, error) {
    return s.OrderRepository.GetOrdersByUserID(userID)
}

func (s *OrderService) UpdateOrderStatus(orderID uint, status utils.OrderStatus) error {
    return s.OrderRepository.UpdateOrderStatus(orderID, status)
}
