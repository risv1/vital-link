package routes

import (
	"vital-link/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func RecordRoutes(app *fiber.App) {
	app.Get("/api/records", controllers.GetRecords)
	app.Get("/api/records/:id", controllers.GetRecord)
	app.Get("/api/records/patient/:patientId", controllers.GetRecordsByPatient)
	app.Post("/api/records", controllers.CreateRecord)
	app.Put("/api/records/:id", controllers.UpdateRecord)
	app.Delete("/api/records/:id", controllers.DeleteRecord)
}