package routes

import (
	"vital-link/api/controllers"
	"vital-link/api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func AppointmentRoutes(app *fiber.App) {
	adminGroup := app.Group("/api/appointments").Use(middlewares.CheckAdmin)
	authGroup := app.Group("/api/appointments").Use(middlewares.CheckAuthenticated)
	authGroup.Post("/", controllers.CreateAppointment)
	adminGroup.Get("/", controllers.GetAppointments)
	adminGroup.Get("/:id", controllers.GetAppointment)
	adminGroup.Get("/doctor/:doctorId", controllers.GetAppointmentsByDoctor)
	adminGroup.Get("/patient/:patientId", controllers.GetAppointmentsByPatient)
	adminGroup.Put("/:id", controllers.UpdateAppointment)
	adminGroup.Delete("/:id", controllers.DeleteAppointment)
}