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
	api.Post("/product/search", controllers.SearchProducts)
	api.Post("/product/view", controllers.ViewProducts) // there will be some kind of pagination
	api.Post("/place", controllers.CreateOrder)         // placing orders
	api.Get("/view/orders", controllers.ViewUserOrders) // viewing orders kind of dashboard

	//	api.Get("/products", controllers.GetProducts)

	// Admin routes
	admin := api.Group("/admin", middleware.AuthMiddleware("admin"))
	admin.Post("/product", controllers.AddProduct)
	admin.Get("/orders", controllers.GetFilteredOrders)
	admin.Get("/stats", controllers.GetOrderStatistics)

	//	admin.Put("/inventory/:id", controllers.UpdateInventory)
}
