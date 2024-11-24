package services

import (
	"OrderManagment/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func Signup(user *models.User) error {
	// Check if the user already exists
	existingUser, err := repostories.GetUserByEmail(user.Email)
	if err == nil && existingUser.ID != 0 {
		return errors.New("email already in use")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Save the user
	return s.repo.CreateUser(user)
}

func Login(email, password string) (string, error) {
	// Get the user by email
	user, err := s.repo.GetUserByEmail(email)
	if err != nil || user.ID == 0 {
		return "", errors.New("invalid email or password")
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := generateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
