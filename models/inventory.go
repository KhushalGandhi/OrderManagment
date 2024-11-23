package models

type Inventory struct {
	ID        uint    `gorm:"primaryKey"`
	ProductID uint    `gorm:"not null;unique"`
	Stock     int     `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID"`
}
