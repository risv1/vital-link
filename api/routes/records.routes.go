package routes

import (
	"vital-link/api/controllers"
	"vital-link/api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RecordRoutes(app *fiber.App) {
	adminGroup := app.Group("/api/records").Use(middlewares.CheckAdmin)
	authGroup := app.Group("/api/records").Use(middlewares.CheckAuthenticated)
	adminGroup.Get("/", controllers.GetRecords)
	adminGroup.Get("/:id", controllers.GetRecord)
	adminGroup.Get("/patient/:patientId", controllers.GetRecordsByPatient)
	adminGroup.Post("/", controllers.CreateRecord)
	adminGroup.Put("/:id", controllers.UpdateRecord)
	adminGroup.Delete("/:id", controllers.DeleteRecord)
	authGroup.Post("/docs", controllers.AddDocument)
	authGroup.Get("/user", controllers.GetUserRecords)
}