package delivery

import (
	"fiber-test/internal/middleware"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/public", publicHandler)
	app.Get("/protected", middleware.AuthMiddleware, protectedHandler)
	app.Get("/generate-token/:user_id", generateTokenHandler)
}

func publicHandler(c *fiber.Ctx) error {
	data := []fiber.Map{
		{"id": 1, "name": "Laptop ASUS ROG", "price": 25000000},
		{"id": 2, "name": "Keyboard Mechanical", "price": 1500000},
		{"id": 3, "name": "Mouse Gaming", "price": 500000},
	}

	return JSONResponse(c, fiber.StatusOK, "This is a public endpoint", data)
}

func protectedHandler(c *fiber.Ctx) error {
	data := []fiber.Map{
		{"id": 101, "name": "Laptop MSI", "price": 12000000},
		{"id": 102, "name": "Keyboard Mechanical", "price": 200000},
		{"id": 103, "name": "Mouse Gaming", "price": 120000},
	}

	return JSONResponse(c, fiber.StatusOK, "This is a protected endpoint", data)
}

func generateTokenHandler(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("user_id"))
	if err != nil {
		return JSONResponse(c, fiber.StatusBadRequest, "Invalid user ID", nil)
	}

	token, err := middleware.GenerateToken(userID)
	if err != nil {
		return JSONResponse(c, fiber.StatusInternalServerError, "Failed to generate token", nil)
	}

	return JSONResponse(c, fiber.StatusOK, "Token generated successfully", fiber.Map{"token": token})
}

func JSONResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  status < 400, 
		"message": message,
		"data":    data,
	})
}
