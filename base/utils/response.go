package utils

import (
	"github.com/gofiber/fiber/v2"
)

// APIResponse defines the structure of a response following the JSend spec.
type APIResponse struct {
	Status  string      `json:"status"`            // "success", "fail", or "error"
	Data    interface{} `json:"data,omitempty"`    // payload for success/fail responses
	Message string      `json:"message,omitempty"` // error message for error responses
}

func SendSuccessResponse(context *fiber.Ctx, message string, data interface{}, code int) error {
	response := APIResponse{
		Data:    data,
		Message: message,
		Status:  "success",
	}
	return context.Status(code).JSON(response)
}

func SendErrorResponse(context *fiber.Ctx, message string, code int) error {
	response := APIResponse{
		Data:    nil,
		Status:  "error",
		Message: message,
	}
	return context.Status(code).JSON(response)
}
