package routes

import (
	"vital-link/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func DoctorRoutes(app *fiber.App) {
	app.Get("/api/doctors", controllers.GetDoctors)
	app.Get("/api/doctors/:id", controllers.GetDoctor)
	app.Post("/api/doctors", controllers.CreateDoctor)
	app.Put("/api/doctors/:id", controllers.UpdateDoctor)
	app.Delete("/api/doctors/:id", controllers.DeleteDoctor)
}