package routes

import (
	"vital-link/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	app.Get("/api/login", controllers.Login)
	app.Get("/auth/google/callback", controllers.Callback)
	app.Post("/api/logout", controllers.Logout)
}