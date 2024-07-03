package main

import (
	"event-booking/database"
	"event-booking/routes"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	port := os.Getenv("PORT")
	
	if port == "" {
		port = "8000"
	}

	if err := app.Listen(":" + port); err != nil {
		panic(err)
	}
}