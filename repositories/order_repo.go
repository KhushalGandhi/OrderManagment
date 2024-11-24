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

func (repo *OrderRepository) ViewUserOrders(userID string) ([]models.Order, error) {
	var orders []models.Order
	err := repo.db.Preload("Products").Where("customer_id = ?", userID).Find(&orders).Error
	return orders, err
}

func (repo *OrderRepository) GetFilteredOrders(userID, productID, sortBy, order string) ([]models.Order, error) {
	var orders []models.Order
	tx := repo.db.Preload("Customer").Preload("Product")

	if userID != "" {
		tx = tx.Where("customer_id = ?", userID)
	}
	if productID != "" {
		tx = tx.Where("product_id = ?", productID)
	}
	err := tx.Order(sortBy + " " + order).Find(&orders).Error
	return orders, err
}

func (repo *OrderRepository) GetOrderStatistics(startDate, endDate string) (map[string]interface{}, error) {
	var totalOrders int64
	var totalRevenue float64

	tx := repo.db.Model(&models.Order)
	if startDate != "" && endDate != "" {
		tx = tx.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}

	err := tx.Count(&totalOrders).Select("SUM(total_amount) as total_revenue").Scan(&totalRevenue).Error
	if err != nil {
		return nil, err
	}

	stats := map[string]interface{}{
		"total_orders":  totalOrders,
		"total_revenue": totalRevenue,
	}
	return stats, nil
}
