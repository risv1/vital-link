package controllers

import "github.com/gofiber/fiber/v2"

func GetPatient(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get Patient"})
}

func GetPatients(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get Patients"})
}

func CreatePatient(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Create Patient"})
}

func UpdatePatient(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Update Patient"})
}

func DeletePatient(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Delete Patient"})
}