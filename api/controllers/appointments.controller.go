package controllers

import (
	"vital-link/api/database"
	"vital-link/api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateAppointment(c *fiber.Ctx) error {
	var body map[string]string
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Appointment creation failed", "error": err})
	}

	appointment := models.Appointments{
		ID: uuid.New().String(),
		DoctorId: body["doctorId"],
		PatientId: body["patientId"],
		Reason: body["reason"],
		Date: body["date"],
		Time: body["time"],
		Status: "pending",
		Notes: "",
	}

	db := database.GetDatabase()

	collection := db.Collection("appointments")
	_, err := collection.InsertOne(c.Context(), appointment)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Appointment creation failed", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Appointment created successfully", "data": appointment})
}

func GetAppointment(c *fiber.Ctx) error {
	
	id := c.Params("id")
	
	db := database.GetDatabase()

	collection := db.Collection("appointments")
	appointment := models.Appointments{}

	err := collection.FindOne(c.Context(), bson.M{"_id": id}).Decode(&appointment)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Appointment retrieval failed", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Appointment retrieved successfully", "data": appointment})
}

func GetAppointments(c *fiber.Ctx) error {
	db := database.GetDatabase()

	collection := db.Collection("appointments")
	appointments := []models.Appointments{}

	cursor, err := collection.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Appointments retrieval failed", "error": err})
	}

	if err = cursor.All(c.Context(), &appointments); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Appointments retrieval failed", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Appointments retrieved successfully", "data": appointments})
}

func UpdateAppointment(c *fiber.Ctx) error {
	id := c.Params("id")
	var body map[string]string

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Appointment update failed", "error": err})
	}

	updatedAppointment := models.Appointments{
		DoctorId: body["doctorId"],
		PatientId: body["patientId"],
		Reason: body["reason"],
		Date: body["date"],
		Time: body["time"],
		Status: body["status"],
		Notes: body["notes"],
	}

	db := database.GetDatabase()

	collection := db.Collection("appointments")
	appointment := models.Appointments{}

	err := collection.FindOneAndUpdate(c.Context(), bson.M{"_id": id}, bson.M{"$set": updatedAppointment}).Decode(&appointment)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Appointment update failed", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Appointment updated successfully", "data": appointment})
}

func DeleteAppointment(c *fiber.Ctx) error {
	id := c.Params("id")
	
	db := database.GetDatabase()

	collection := db.Collection("appointments")
	appointment := models.Appointments{}

	err := collection.FindOneAndDelete(c.Context(), bson.M{"_id": id}).Decode(&appointment)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Appointment deletion failed", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Appointment deleted successfully", "data": appointment})
}

func GetAppointmentsByDoctor(c *fiber.Ctx) error {
	doctorId := c.Params("doctorId")

	db := database.GetDatabase()

	collection := db.Collection("appointments")
	appointments := []models.Appointments{}

	cursor, err := collection.Find(c.Context(), bson.M{"doctorId": doctorId})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Appointments retrieval failed", "error": err})
	}

	if err = cursor.All(c.Context(), &appointments); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Appointments retrieval failed", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Appointments retrieved successfully", "data": appointments})
}

func GetAppointmentsByPatient(c *fiber.Ctx) error {
	patientId := c.Params("patientId")

	db := database.GetDatabase()

	collection := db.Collection("appointments")
	appointments := []models.Appointments{}

	cursor, err := collection.Find(c.Context(), bson.M{"patientId": patientId})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Appointments retrieval failed", "error": err})
	}

	if err = cursor.All(c.Context(), &appointments); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Appointments retrieval failed", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Appointments retrieved successfully", "data": appointments})
}