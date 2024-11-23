package repositories

import (
	"OrderManagment/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (repo *OrderRepository) CreateOrder(order *models.Order) error {
	return repo.db.Create(order).Error
}

func (repo *OrderRepository) GetOrdersByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := repo.db.Where("customer_id = ?", userID).Find(&orders).Error
	return orders, err
}
