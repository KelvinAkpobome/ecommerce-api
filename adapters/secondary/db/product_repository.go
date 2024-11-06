package db

import (
	"ecommerce-api/domain"
	"ecommerce-api/ports"
	"fmt"
	"gorm.io/gorm"
)

type ProductRepositoryDB struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ports.ProductRepository {
	return &ProductRepositoryDB{DB: db}
}

func (r *ProductRepositoryDB) CreateProduct(product *domain.Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductRepositoryDB) GetProductByID(id uint) (*domain.Product, error) {
	var product domain.Product
	if err := r.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepositoryDB) UpdateProduct(product *domain.Product) error {
	return r.DB.Omit("CreatedAt").Save(product).Error
}

func (r *ProductRepositoryDB) DeleteProduct(id uint) error {
	return r.DB.Delete(&domain.Product{}, id).Error
}

func (r *ProductRepositoryDB) GetAllProducts() ([]domain.Product, error) {
	var products []domain.Product
	if err := r.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepositoryDB) GetProductsById(ids []uint) ([]domain.Product, error) {
	var products []domain.Product

	if len(ids) == 0 {
		return nil, fmt.Errorf("no product IDs provided")
	}

	if err := r.DB.Where("id IN ?", ids).Find(&products).Error; err != nil {
		return nil, err
	}

	if len(products) != len(ids) {
		return nil, fmt.Errorf("one or more products not found")
	}

	return products, nil
}
