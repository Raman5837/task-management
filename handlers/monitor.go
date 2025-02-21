package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

// Handler To Monitor Service Resources
func ResourceMonitor() fiber.Handler {
	return monitor.New(monitor.Config{Title: "Task management service metrics page"})
}
