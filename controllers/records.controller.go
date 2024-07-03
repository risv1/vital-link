package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetRecord(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get Record"})
}

func GetRecords(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get Records"})
}

func CreateRecord(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Create Record"})
}

func UpdateRecord(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Update Record"})
}

func DeleteRecord(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Delete Record"})
}

func GetRecordsByPatient(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get Records by Patient"})
}