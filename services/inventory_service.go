package services

import (
	"OrderManagment/models"
	"OrderManagment/repositories"
)

type InventoryService struct {
	repo *repositories.InventoryRepository
}

func NewInventoryService(repo *repositories.InventoryRepository) *InventoryService {
	return &InventoryService{repo}
}

func (s *InventoryService) UpdateInventory(inventory *models.Inventory) error {
	return s.repo.UpdateInventory(inventory)
}

func (s *InventoryService) GetInventoryByProductID(productID uint) (*models.Inventory, error) {
	return s.repo.GetInventoryByProductID(productID)
}