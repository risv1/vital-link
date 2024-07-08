package routes

import (
	"vital-link/api/controllers"
	"vital-link/api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func PatientRoutes(app *fiber.App) {
	adminGroup := app.Group("/api/patients").Use(middlewares.CheckAdmin)
	authGroup := app.Group("/api/patients").Use(middlewares.CheckAuthenticated)
	authGroup.Post("/", controllers.CreatePatient)
	adminGroup.Get("/", controllers.GetPatients)
	adminGroup.Get("/:id", controllers.GetPatient)
	adminGroup.Put("/:id", controllers.UpdatePatient)
	adminGroup.Delete("/:id", controllers.DeletePatient)
}