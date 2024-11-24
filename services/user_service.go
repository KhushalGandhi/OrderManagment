package services

import (
	"OrderManagment/models"
	"OrderManagment/repositories"
)

type OrderService struct {
	repo *repositories.OrderRepository
}

func NewOrderService(repo *repositories.OrderRepository) *OrderService {
	return &OrderService{repo}
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	return s.repo.CreateOrder(order)
}

func (s *OrderService) ViewUserOrders(userID string) ([]models.Order, error) {
	return s.repo.ViewUserOrders(userID)
}

func (s *OrderService) GetFilteredOrders(userID, productID, sortBy, order string) ([]models.Order, error) {
	return s.repo.GetFilteredOrders(userID, productID, sortBy, order)
}

func (s *ProductService) ViewProducts(category string, limit, offset int) ([]models.Product, error) {
	return s.repo.ViewProducts(category, limit, offset)
}
