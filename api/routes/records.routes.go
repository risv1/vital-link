package routes

import (
	"vital-link/api/controllers"
	"vital-link/api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RecordRoutes(app *fiber.App) {
	adminGroup := app.Group("/api/records").Use(middlewares.CheckAdmin)
	adminGroup.Get("/api/records", controllers.GetRecords)
	adminGroup.Get("/api/records/:id", controllers.GetRecord)
	adminGroup.Get("/api/records/patient/:patientId", controllers.GetRecordsByPatient)
	adminGroup.Post("/api/records", controllers.CreateRecord)
	adminGroup.Put("/api/records/:id", controllers.UpdateRecord)
	adminGroup.Delete("/api/records/:id", controllers.DeleteRecord)
}