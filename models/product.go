package models

type Product struct {
	ID       uint    `gorm:"primaryKey"`
	Name     string  `gorm:"not null"`
	Price    float64 `gorm:"not null"`
	Stock    int     `gorm:"not null"`
	Category string
	Orders   []Order `gorm:"many2many:order_products;"`
}
