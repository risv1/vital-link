package controllers

import (
	"vital-link/api/database"
	"vital-link/api/models"
	"github.com/google/uuid"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetDoctor(c *fiber.Ctx) error {
	id := c.Params("id")
	
	db := database.GetDatabase()

	collection := db.Collection("doctors")
	doctor := models.Doctor{}

	err := collection.FindOne(c.Context(), bson.M{"_id": id}).Decode(&doctor)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Doctor retrieval failed", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Doctor retrieved successfully", "data": doctor})
}

func GetDoctors(c *fiber.Ctx) error {
	db := database.GetDatabase()

	collection := db.Collection("doctors")
	doctors := []models.Doctor{}

	cursor, err := collection.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Doctors retrieval failed", "error": err})
	}

	if err := cursor.All(c.Context(), &doctors); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Doctors retrieval failed", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Doctors retrieved successfully", "data": doctors})
}

func CreateDoctor(c *fiber.Ctx) error {
	var body map[string]string

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Doctor creation failed", "error": err})
	}

	doctor := models.Doctor{
		ID:                uuid.New().String(),
		FirstName:         body["firstName"],
		LastName:          body["lastName"],
		Specialty:         body["specialty"],
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
		Availability: []models.Availability{},
		Image:            body["image"],
	}

	db := database.GetDatabase()

	collection := db.Collection("doctors")
	_, err := collection.InsertOne(c.Context(), doctor)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Doctor creation failed", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Doctor created successfully", "data": doctor})
}

func UpdateDoctor(c *fiber.Ctx) error {
	id := c.Params("id")
	var body map[string]string

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Doctor update failed", "error": err})
	}

	db := database.GetDatabase()

	updatedDoctor := models.Doctor{
		FirstName:         body["firstName"],
		LastName:          body["lastName"],
		Specialty:         body["specialty"],
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
		Image:            body["image"],
	}

	collection := db.Collection("doctors")
	doctor := models.Doctor{}

	err := collection.FindOneAndUpdate(c.Context(), bson.M{"_id": id}, bson.M{"$set": updatedDoctor}).Decode(&doctor)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Doctor update failed", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Doctor updated successfully", "data": doctor})
}

func DeleteDoctor(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.GetDatabase()

	collection := db.Collection("doctors")
	_, err := collection.DeleteOne(c.Context(), bson.M{"_id": id})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Doctor deletion failed", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Doctor deleted successfully"})
}
