package ports

import (
	"ecommerce-api/domain"
	"ecommerce-api/utils"
)

type OrderRepository interface {
    CreateOrder(order *domain.Order) error
    GetOrderByID(id uint) (*domain.Order, error)
    GetOrdersByUserID(userID uint) ([]domain.Order, error)
    UpdateOrderStatus(orderID uint, status utils.OrderStatus) error
}
