package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"vital-link/api/database"
	"vital-link/api/models"
)

func GetRecord(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.GetDatabase()

	collection := db.Collection("records")
	record := models.Records{}

	err := collection.FindOne(c.Context(), bson.M{"_id": id}).Decode(&record)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Record retrieval failed", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Record retrieved successfully", "data": record})
}

func GetRecords(c *fiber.Ctx) error {
	db := database.GetDatabase()

	collection := db.Collection("records")
	records := []models.Records{}

	cursor, err := collection.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Records retrieval failed", "error": err})
	}

	if err := cursor.All(c.Context(), &records); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Records retrieval failed", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Records retrieved successfully", "data": records})
}

func CreateRecord(c *fiber.Ctx) error {
	var body map[string]string

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Record creation failed", "error": err})
	}

	record := models.Records{
		ID:          uuid.New().String(),
		PatientId:   body["patientId"],
		DoctorId:    body["doctorId"],
		Date:        body["date"],
		Time:        body["time"],
		Type:        body["type"],
		Description: body["description"],
		Documents:   []models.Documents{},
	}

	db := database.GetDatabase()
	collection := db.Collection("records")

	_, err := collection.InsertOne(c.Context(), record)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Record creation failed", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Record created successfully", "data": record})
}

func UpdateRecord(c *fiber.Ctx) error {
	var body map[string]string

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Record update failed", "error": err})
	}

	id := c.Params("id")

	db := database.GetDatabase()
	collection := db.Collection("records")
	record := models.Records{}

	err := collection.FindOne(c.Context(), bson.M{"_id": id}).Decode(&record)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Record update failed", "error": err})
	}

	updatedRecord := models.Records{
		PatientId:   body["patientId"],
		DoctorId:    body["doctorId"],
		Date:        body["date"],
		Time:        body["time"],
		Type:        body["type"],
		Description: body["description"],
	}

	_, err = collection.UpdateOne(c.Context(), bson.M{"_id": id}, bson.M{"$set": updatedRecord})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Record update failed", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Record updated successfully", "data": updatedRecord})
}

func DeleteRecord(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.GetDatabase()
	collection := db.Collection("records")

	_, err := collection.DeleteOne(c.Context(), bson.M{"_id": id})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Record deletion failed", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Record deleted successfully"})
}

func GetRecordsByPatient(c *fiber.Ctx) error {
	patientId := c.Params("patientId")
	db := database.GetDatabase()

	collection := db.Collection("records")
	records := []models.Records{}

	cursor, err := collection.Find(c.Context(), bson.M{"patientId": patientId})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Records retrieval failed", "error": err})
	}

	if err := cursor.All(c.Context(), &records); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Records retrieval failed", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Records retrieved successfully", "data": records})
}

func AddDocument(c *fiber.Ctx) error {
	recordId := c.Params("recordId")

	documentType := c.FormValue("documentType")

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Document creation failed", "error": err})
	}

	db := database.GetDatabase()
	collection := db.Collection("records")
	record := models.Records{}

	err = collection.FindOne(c.Context(), bson.M{"_id": recordId}).Decode(&record)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Document creation failed", "error": err})
	}

	// yet to implement a object storage such as s3 or cloudflare-r2
	filePath := "/" + file.Filename 

	addDoc := models.Records{
		Documents: []models.Documents{
			{
				DocumentId:   uuid.New().String(),
				DocumentType: documentType,
				Url:     filePath,
			},
		},
	}

	_, err = collection.UpdateOne(c.Context(), bson.M{"_id": recordId}, bson.M{"$push": addDoc})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Document creation failed", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Document added successfully"})
}
