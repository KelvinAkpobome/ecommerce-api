package application

import (
    "ecommerce-api/domain"
    "ecommerce-api/ports"
)

type ProductService struct {
    ProductRepository ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) *ProductService {
    return &ProductService{ProductRepository: repo}
}

func (s *ProductService) CreateProduct(product *domain.Product) error {
    return s.ProductRepository.CreateProduct(product)
}

func (s *ProductService) GetProductByID(id uint) (*domain.Product, error) {
    return s.ProductRepository.GetProductByID(id)
}

func (s *ProductService) UpdateProduct(product *domain.Product) error {
    return s.ProductRepository.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
    return s.ProductRepository.DeleteProduct(id)
}

func (s *ProductService) GetAllProducts() ([]domain.Product, error) {
    return s.ProductRepository.GetAllProducts()
}

func (s *ProductService) GetProductsById(ids []uint) ([]domain.Product, error) {
    return s.ProductRepository.GetProductsById(ids)
}