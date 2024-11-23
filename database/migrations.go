package database

import (
	"OrderManagment/models"
	"log"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	log.Println("Running database migrations...")

	// AutoMigrate models
	if err := db.AutoMigrate(&models.Customer{}, &models.Product{}, &models.Order{}, &models.Inventory{}); err != nil {
		return err
	}

	// Create trigger for dynamic pricing
	triggerQuery := `
	CREATE OR REPLACE FUNCTION update_product_price()
	RETURNS TRIGGER AS $$
	BEGIN
		IF NEW.stock < 10 THEN
			UPDATE products SET price = price * 1.10 WHERE id = NEW.product_id;
		ELSE
			UPDATE products SET price = price * 0.90 WHERE id = NEW.product_id;
		END IF;
		RETURN NEW;
	END;
	$$ LANGUAGE plpgsql;

	CREATE TRIGGER trigger_update_price
	AFTER UPDATE ON inventories
	FOR EACH ROW
	EXECUTE FUNCTION update_product_price();
	`
	if err := db.Exec(triggerQuery).Error; err != nil {
		return err
	}

	log.Println("Migrations completed successfully!")
	return nil
}
