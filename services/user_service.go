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

func (s *OrderService) GetOrdersByUserID(userID uint) ([]models.Order, error) {
	return s.repo.GetOrdersByUserID(userID)
}

func (s *OrderService) GetFilteredOrders(userID, productID, sortBy, order string) ([]models.Order, error) {
	return s.repo.GetFilteredOrders(userID, productID, sortBy, order)
}
