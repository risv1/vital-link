package routes

import (
	"vital-link/api/controllers"
	"vital-link/api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func AppointmentRoutes(app *fiber.App) {
	adminGroup := app.Group("/api/appointments").Use(middlewares.CheckAdmin)
	authGroup := app.Group("/api/appointments").Use(middlewares.CheckAuthenticated)
	authGroup.Post("/api/appointments", controllers.CreateAppointment)
	adminGroup.Get("/api/appointments", controllers.GetAppointments)
	adminGroup.Get("/api/appointments/:id", controllers.GetAppointment)
	adminGroup.Get("/api/appointments/doctor/:doctorId", controllers.GetAppointmentsByDoctor)
	adminGroup.Get("/api/appointments/patient/:patientId", controllers.GetAppointmentsByPatient)
	adminGroup.Put("/api/appointments/:id", controllers.UpdateAppointment)
	adminGroup.Delete("/api/appointments/:id", controllers.DeleteAppointment)
}