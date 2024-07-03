package routes

import (
	"event-booking/controllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	app.Post("/api/login", controllers.Login)
	app.Post("/api/logout", controllers.Logout)
}