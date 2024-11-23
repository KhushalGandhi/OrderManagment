package models

import "time"

type Order struct {
	ID         uint    `gorm:"primaryKey"`
	CustomerID uint    `gorm:"not null"`
	TotalPrice float64 `gorm:"not null"`
	CreatedAt  time.Time
	Products   []Product `gorm:"many2many:order_products;"`
}
