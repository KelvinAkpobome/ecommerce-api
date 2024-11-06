package db

import (
	"ecommerce-api/domain"
	"ecommerce-api/ports"
	"ecommerce-api/utils"

	"gorm.io/gorm"
)

type OrderRepositoryDB struct {
    DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) ports.OrderRepository {
    return &OrderRepositoryDB{DB: db}
}

func (r *OrderRepositoryDB) CreateOrder(order *domain.Order) error {
    return r.DB.Create(order).Error
}

func (r *OrderRepositoryDB) GetOrderByID(id uint) (*domain.Order, error) {
    var order domain.Order
    if err := r.DB.First(&order, id).Error; err != nil {
        return nil, err
    }
    return &order, nil
}
func (r *OrderRepositoryDB) GetOrdersByUserID(userID uint) ([]domain.Order, error) {
    var orders []domain.Order
    if err := r.DB.Preload("Products").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
        return nil, err
    }
    return orders, nil
}


func (r *OrderRepositoryDB) UpdateOrderStatus(orderID uint, status utils.OrderStatus) error {
    return r.DB.Model(&domain.Order{}).Where("id = ?", orderID).Update("status", status).Error
}
