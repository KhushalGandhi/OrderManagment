package controllers

import (
	"OrderManagment/models"
	"OrderManagment/services"
	"github.com/gofiber/fiber/v2"
)

var orderService *services.OrderService

func InitializeUserController(service *services.OrderService) {
	orderService = service
}

func SearchProducts(c *fiber.Ctx) error {
	// Query parameters for search
	query := c.Query("query", "")
	category := c.Query("category", "")
	limit := c.QueryInt("limit", 10)
	offset := c.QueryInt("offset", 0)

	products, err := productService.SearchProducts(query, category, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to search products"})
	}
	return c.JSON(products)
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := orderService.CreateOrder(&order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create order"})
	}
	return c.Status(fiber.StatusCreated).JSON(order)
}

func ViewUserOrders(c *fiber.Ctx) error {
	userID := c.Query("user_id", "")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User ID is required"})
	}

	orders, err := orderService.ViewUserOrders(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve orders"})
	}
	return c.JSON(orders)
}

func ViewProducts(c *fiber.Ctx) error {
	category := c.Query("category", "")
	limit := c.QueryInt("limit", 10)
	offset := c.QueryInt("offset", 0)

	products, err := productService.ViewProducts(category, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch products"})
	}
	return c.JSON(products)
}
