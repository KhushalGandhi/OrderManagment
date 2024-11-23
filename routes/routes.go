package routes

import (
	"OrderManagment/controllers"
	"OrderManagment/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// User routes
	api.Post("/signup", controllers.Signup)
	api.Post("/login", controllers.Login)
	api.Get("/products", controllers.GetProducts)

	// Admin routes
	admin := api.Group("/admin", middleware.AuthMiddleware("admin"))
	admin.Post("/product", controllers.AddProduct)
	admin.Put("/inventory/:id", controllers.UpdateInventory)
}
