package cmd

import (
	"fiber-test/internal/delivery"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Execute() {
	app := fiber.New()

	delivery.RegisterRoutes(app)

	log.Fatal(app.Listen(":8080"))
}