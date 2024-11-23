package controllers

import (
	"OrderManagment/models"
	"OrderManagment/services"
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
