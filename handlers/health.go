package handlers

import "github.com/gofiber/fiber/v2"

// Health Check API Handler
func HealthCheck(context *fiber.Ctx) error {
	return context.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"data":    nil,
			"error":   nil,
			"message": "Task management service is up and running",
		},
	)
}
