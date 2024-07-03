package controllers

import "github.com/gofiber/fiber/v2"

func CreateAppointment(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Create Appointment"})
}

func GetAppointment(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get Appointment"})
}

func GetAppointments(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get Appointments"})
}

func UpdateAppointment(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Update Appointment"})
}

func DeleteAppointment(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Delete Appointment"})
}

func GetAppointmentsByDoctor(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get Appointments by Doctor"})
}

func GetAppointmentsByPatient(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get Appointments by Patient"})
}