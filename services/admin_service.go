package services

import (
	"OrderManagment/models"
	"OrderManagment/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo}
}

func (s *ProductService) AddProduct(product *models.Product) error {
	return s.repo.CreateProduct(product)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetProducts()
}

func (s *ProductService) UpdateProduct(product *models.Product) error {
	return s.repo.UpdateProduct(product)
}

func (s *OrderService) GetOrderStatistics(startDate, endDate string) (map[string]interface{}, error) {
	return s.repo.GetOrderStatistics(startDate, endDate)
}
