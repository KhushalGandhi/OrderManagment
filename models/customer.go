package models

import "time"

type Customer struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Orders   []Order
}

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"uniqueIndex"`
	Password  string `json:"-"`
	CreatedAt time.Time
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
