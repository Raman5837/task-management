package routes

import (
	"github.com/Raman5837/task-management/handlers"
	"github.com/gofiber/fiber/v2"
)

// Register all routes from all apps
func RegisterAll(app *fiber.App) {

	base := app.Group("/")

	base.Get("/", handlers.BaseHandler)
	base.Get("/api/check", handlers.HealthCheck)
	base.Get("/api/monitor", handlers.ResourceMonitor())

}
