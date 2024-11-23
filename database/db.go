package database

import (
	"fmt"
	"order-inventory-management/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.GetEnv("DB_HOST", "localhost"),
		config.GetEnv("DB_USER", "postgres"),
		config.GetEnv("DB_PASSWORD", "password"),
		config.GetEnv("DB_NAME", "inventory_db"),
		config.GetEnv("DB_PORT", "5432"),
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
