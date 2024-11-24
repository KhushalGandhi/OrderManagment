package controllers

import (
	"OrderManagment/models"
	"OrderManagment/services"
	"github.com/gofiber/fiber/v2"
)

var productService *services.ProductService

func InitializeAdminController(service *services.ProductService) {
	productService = service
}

func AddProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := productService.AddProduct(&product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add product"})
	}
	return c.Status(fiber.StatusCreated).JSON(product)
}

func GetAllProducts(c *fiber.Ctx) error {
	products, err := productService.GetAllProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch products"})
	}
	return c.JSON(products)
}

func GetOrderStatistics(c *fiber.Ctx) error {
	startDate := c.Query("start_date", "")
	endDate := c.Query("end_date", "")

	stats, err := orderService.GetOrderStatistics(startDate, endDate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve statistics"})
	}
	return c.JSON(stats)
}

func GetFilteredOrders(c *fiber.Ctx) error {
	userID := c.Query("user_id", "")
	productID := c.Query("product_id", "")
	sortBy := c.Query("sort_by", "created_at")
	order := c.Query("order", "asc")

	orders, err := orderService.GetFilteredOrders(userID, productID, sortBy, order)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve orders"})
	}
	return c.JSON(orders)
}
