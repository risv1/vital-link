package controllers 

import "github.com/gofiber/fiber/v2"

func GetDoctor(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get Doctor"})
}

func GetDoctors(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get Doctors"})
}

func CreateDoctor(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Create Doctor"})
}

func UpdateDoctor(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Update Doctor"})
}

func DeleteDoctor(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Delete Doctor"})
}
