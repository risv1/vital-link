package controllers

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Login"})
}

func Logout (c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Logout"})
}