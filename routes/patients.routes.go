package routes 

import (
	"event-booking/controllers"
	"github.com/gofiber/fiber/v2"
)

func PatientRoutes(app *fiber.App) {
	app.Get("/api/patients", controllers.GetPatients)
	app.Get("/api/patients/:id", controllers.GetPatient)
	app.Post("/api/patients", controllers.CreatePatient)
	app.Put("/api/patients/:id", controllers.UpdatePatient)
	app.Delete("/api/patients/:id", controllers.DeletePatient)
}