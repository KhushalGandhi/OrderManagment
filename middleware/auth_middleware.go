package middleware

import (
	"strings"
)

func AuthMiddleware(role string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Get the token from the Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization header required"})
		}

		// Extract token (Bearer token)
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token required"})
		}

		// For simplicity, assuming a basic check here
		// You should verify the token and role using a JWT library
		if role == "admin" && token != "admin_token" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You don't have the necessary permissions"})
		}

		// Proceed to the next handler
		return c.Next()
	}
}
