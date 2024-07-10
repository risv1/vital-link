package middlewares

import "github.com/gofiber/fiber/v2"

func CheckAuthenticated(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Next()
}