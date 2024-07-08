package routes

import (
	"vital-link/api/controllers"
	"vital-link/api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func DoctorRoutes(app *fiber.App) {
	adminGroup := app.Group("/api/appointments").Use(middlewares.CheckAdmin)
	app.Get("/api/doctors", controllers.GetDoctors)
	app.Get("/api/doctors/:id", controllers.GetDoctor)
	adminGroup.Post("/api/doctors", controllers.CreateDoctor)
	adminGroup.Put("/api/doctors/:id", controllers.UpdateDoctor)
	adminGroup.Delete("/api/doctors/:id", controllers.DeleteDoctor)
}