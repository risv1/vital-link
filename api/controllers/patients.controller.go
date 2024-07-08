package controllers

import (
	"vital-link/api/database"
	"vital-link/api/models"
	"github.com/google/uuid"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetPatient(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.GetDatabase()

	collection := db.Collection("patients")
	patient := models.Patient{}

	err := collection.FindOne(c.Context(), bson.M{"_id": id}).Decode(&patient)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Patient retrieval failed", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Patient retrieved successfully", "data": patient})
}

func GetPatients(c *fiber.Ctx) error {
	db := database.GetDatabase()

	collection := db.Collection("patients")
	patients := []models.Patient{}

	cursor, err := collection.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Patients retrieval failed", "error": err})
	}

	if err := cursor.All(c.Context(), &patients); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Patients retrieval failed", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Patients retrieved successfully", "data": patients})
}

func CreatePatient(c *fiber.Ctx) error {
	var body map[string]string

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Patient creation failed", "error": err})
	}

	patient := models.Patient{
		ID:                uuid.New().String(),
		FirstName:         body["firstName"],
		LastName:          body["lastName"],
		DateOfBirth: 	 body["dateOfBirth"],
		Gender: 		 body["gender"],
		ContactInformation: models.ContactInformation{
			Email: body["email"],
			Phone: body["phone"],
			Address: models.Address{
				Street: body["street"],
				City:   body["city"],
				State: body["state"],
				Pincode: body["pincode"],
				Country: body["country"],
			},
		},
		MedicalHistory: []models.MedicalHistory{},
		Appointments: []string{},
		Image: body["image"],
	}

	db := database.GetDatabase()

	collection := db.Collection("patients")
	_, err := collection.InsertOne(c.Context(), patient)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Patient creation failed", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Patient created successfully", "data": patient})
}

func UpdatePatient(c *fiber.Ctx) error {
	id := c.Params("id")
	var body map[string]string

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Patient update failed", "error": err})
	}

	db := database.GetDatabase()

	updatedPatient := models.Patient{
		FirstName:         body["firstName"],
		LastName:          body["lastName"],
		DateOfBirth: 	 body["dateOfBirth"],
		Gender: body["gender"],
		ContactInformation: models.ContactInformation{
			Email: body["email"],
			Phone: body["phone"],
			Address: models.Address{
				Street: body["street"],
				City:   body["city"],
				State: body["state"],
				Pincode: body["pincode"],
				Country: body["country"],
			},
		},
		MedicalHistory: []models.MedicalHistory{},
		Appointments: []string{},
		Image: body["image"],
	}

	collection := db.Collection("patients")
	_, err := collection.UpdateOne(c.Context(), bson.M{"_id": id}, bson.M{"$set": updatedPatient})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Patient update failed", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Patient updated successfully", "data": updatedPatient})
}

func DeletePatient(c *fiber.Ctx) error {
	id := c.Params("id")
	
	db := database.GetDatabase()

	collection := db.Collection("patients")
	_, err := collection.DeleteOne(c.Context(), bson.M{"_id": id})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Patient deletion failed", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Patient deleted successfully"})
}