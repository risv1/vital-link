package main

import (
	"vital-link/api/routes"
	"os"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	port := os.Getenv("PORT")
	
	if port == "" {
		port = "8000"
	}

	routes.RecordRoutes(app)
	routes.PatientRoutes(app)
	routes.DoctorRoutes(app)
	routes.AuthRoutes(app)
	routes.AppointmentRoutes(app)

	if err := app.Listen(":" + port); err != nil {
		panic(err)
	}
}