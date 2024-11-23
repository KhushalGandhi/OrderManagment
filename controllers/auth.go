package controllers

import "github.com/gofiber/fiber/v2"

func Signup(c *fiber.Ctx) error {
	// Signup logic
	return c.JSON(fiber.Map{"message": "User signed up"})
}

func Login(c *fiber.Ctx) error {
	// Login logic
	return c.JSON(fiber.Map{"message": "User logged in"})
}
