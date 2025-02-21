package main

import (
	"os"

	"github.com/Raman5837/task-management/base/configuration"
	"github.com/Raman5837/task-management/base/constants"
	"github.com/Raman5837/task-management/base/settings"
	"github.com/Raman5837/task-management/routes"
)

// EntryPoint of the app
func main() {

	// Returns a new Fiber app instance
	app := settings.InitializeApp()
	Logger := configuration.GetLogger()

	// Registering all the routes
	routes.RegisterAll(app)

	// Graceful Shutdown
	shutdown := make(chan os.Signal, 1)
	settings.GracefulShutdownHandler(app, shutdown)

	// This will get execute after the main function to cleanup the resource
	defer settings.InitiateCleanupProcess()

	serverPort := ":" + constants.GetEnv("SERVER_PORT")

	if serverError := app.Listen(serverPort); serverError != nil {
		Logger.Fatal(serverError, "Error starting task management service")
	}
}
