package handlers

import "github.com/gofiber/fiber/v2"

// Base API Handler
func BaseHandler(context *fiber.Ctx) error {
	errors := context.SendString("Task management service is ready to serve!")
	return errors
}
