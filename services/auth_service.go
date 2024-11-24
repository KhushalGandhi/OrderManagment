package services

import (
	"OrderManagment/models"
	"OrderManagment/repositories"
	"OrderManagment/utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func Signup(user *models.SignUpRequest) error {
	// Check if the user already exists
	existingUser, err := repositories.GetUserByEmail(user.Email)
	if err == nil && existingUser.ID != 0 {
		return errors.New("email already in use")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	//user.Password = string(hashedPassword)

	baseModel := models.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Time{},
	}

	// Save the user
	return repositories.CreateUser(&baseModel)
}

func Login(email, password string) (string, error) {
	// Get the user by email
	user, err := repositories.GetUserByEmail(email)
	if err != nil || user.ID == 0 {
		return "", errors.New("invalid email or password")
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
