package main

import (
	"OrderManagment/config"
	"OrderManagment/database"
	"OrderManagment/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	// Load configuration
	config.LoadEnv()

	// Initialize database
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Run migrations
	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize Fiber app
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
