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
	ID        uint      `json:"id" gorm:"autoIncrement;primaryKey"`
	Name      string    `gorm:"name" json:"name"`
	Email     string    `json:"email" gorm:"email;uniqueIndex"`
	Password  string    `gorm:"password" json:"password"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
}

type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email" `
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
