package repositories

import "OrderManagment/models"

func CreateUser(user *models.User) error {
	return repo.db.Create(user).Error
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := repo.db.Where("email = ?", email).First(&user).Error
	return &user, err
}
