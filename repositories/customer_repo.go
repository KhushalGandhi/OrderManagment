package repositories

import (
	"OrderManagment/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (p *UserRepository) CreateUser(user *models.User) error {
	return repo.db.Create(user).Error
}

func (p *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := repo.db.Where("email = ?", email).First(&user).Error
	return &user, err
}
