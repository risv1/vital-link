package routes

import (
	"vital-link/api/controllers"
	"vital-link/api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func DoctorRoutes(app *fiber.App) {
	adminGroup := app.Group("/api/doctors").Use(middlewares.CheckAdmin)
	app.Get("/", controllers.GetDoctors)
	app.Get("/:id", controllers.GetDoctor)
	adminGroup.Post("/", controllers.CreateDoctor)
	adminGroup.Put("/:id", controllers.UpdateDoctor)
	adminGroup.Delete("/:id", controllers.DeleteDoctor)
}