package ports

import "ecommerce-api/domain"

type ProductRepository interface {
    CreateProduct(product *domain.Product) error
    GetProductByID(id uint) (*domain.Product, error)
    UpdateProduct(product *domain.Product) error
    DeleteProduct(id uint) error
    GetAllProducts() ([]domain.Product, error)
    GetProductsById(ids []uint)([]domain.Product, error)
}
