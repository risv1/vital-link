package routes

import (
	"vital-link/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func AppointmentRoutes(app *fiber.App) {
	app.Get("/api/appointments", controllers.GetAppointments)
	app.Get("/api/appointments/:id", controllers.GetAppointment)
	app.Get("/api/appointments/doctor/:doctorId", controllers.GetAppointmentsByDoctor)
	app.Get("/api/appointments/patient/:patientId", controllers.GetAppointmentsByPatient)
	app.Post("/api/appointments", controllers.CreateAppointment)
	app.Put("/api/appointments/:id", controllers.UpdateAppointment)
	app.Delete("/api/appointments/:id", controllers.DeleteAppointment)
}