package routes

import (
	"vital-link/api/controllers"
	"vital-link/api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func PatientRoutes(app *fiber.App) {
	adminGroup := app.Group("/api/appointments").Use(middlewares.CheckAdmin)
	authGroup := app.Group("/api/appointments").Use(middlewares.CheckAuthenticated)
	authGroup.Post("/api/patients", controllers.CreatePatient)
	adminGroup.Get("/api/patients", controllers.GetPatients)
	adminGroup.Get("/api/patients/:id", controllers.GetPatient)
	adminGroup.Put("/api/patients/:id", controllers.UpdatePatient)
	adminGroup.Delete("/api/patients/:id", controllers.DeletePatient)
}