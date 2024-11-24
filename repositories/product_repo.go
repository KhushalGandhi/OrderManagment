package repositories

import (
	"OrderManagment/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (repo *ProductRepository) CreateProduct(product *models.Product) error {
	return repo.db.Create(product).Error
}

func (repo *ProductRepository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	err := repo.db.Find(&products).Error
	return products, err
}

func (repo *ProductRepository) ViewProducts(category string, limit, offset int) ([]models.Product, error) {
	var products []models.Product
	tx := repo.db.Model(&models.Product)

	if category != "" {
		tx = tx.Where("category = ?", category)
	}
	err := tx.Limit(limit).Offset(offset).Find(&products).Error
	return products, err
}

func (repo *ProductRepository) UpdateProduct(product *models.Product) error {
	return repo.db.Save(product).Error
}

func (repo *ProductRepository) SearchProducts(query, category string, limit, offset int) ([]models.Product, error) {
	var products []models.Product
	tx := repo.db.Model(models.Product)

	if query != "" {
		tx = tx.Where("name ILIKE ?", "%"+query+"%")
	}
	if category != "" {
		tx = tx.Where("category = ?", category)
	}
	err := tx.Limit(limit).Offset(offset).Find(&products).Error
	return products, err
}
