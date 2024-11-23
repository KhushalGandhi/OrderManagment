package repositories

import (
	"OrderManagment/models"
	"gorm.io/gorm"
)

type InventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) *InventoryRepository {
	return &InventoryRepository{db}
}

func (repo *InventoryRepository) UpdateInventory(inventory *models.Inventory) error {
	return repo.db.Save(inventory).Error
}

func (repo *InventoryRepository) GetInventoryByProductID(productID uint) (*models.Inventory, error) {
	var inventory models.Inventory
	err := repo.db.Where("product_id = ?", productID).First(&inventory).Error
	return &inventory, err
}
