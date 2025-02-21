package settings

import (
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	"github.com/Raman5837/task-management/base/configuration"
	"github.com/Raman5837/task-management/base/constants"
	"github.com/Raman5837/task-management/base/database"
	"github.com/Raman5837/task-management/base/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Add all the defined middleware (Ordering is important)
func addMiddleware(app *fiber.App) {

	// 1. Add a request id in all the requests
	app.Use(middleware.RequestId())

	// 2. Log all the requests
	app.Use(middleware.APILogger())

}

// Connect to DataBase
func connectDataBase() {

	Logger := configuration.GetLogger()

	// Initialize Database Connections
	if dbError := database.EstablishConnection(); dbError != nil {
		Logger.Fatal(dbError, "Error while connecting to database: ")
	}

}

// Initialize a new Fiber app (Setups all the Middleware, DB connections and external services)
func InitializeApp() *fiber.App {

	config := fiber.Config{
		ReadBufferSize: 8190,
		ServerHeader:   "task-management",
		AppName:        constants.GetEnv("APP_NAME"),
	}

	app := fiber.New(config)

	allowedMethods := strings.Join([]string{
		fiber.MethodGet,
		fiber.MethodPost,
		fiber.MethodHead,
		fiber.MethodPut,
		fiber.MethodPatch,
		fiber.MethodDelete,
		fiber.MethodOptions,
	}, ",")

	allowedHeaders := "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin"

	// Regular Expression To Match The Allowed Origins
	var domainRegex = regexp.MustCompile(`^https://task-management\.in/?$`)

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     allowedHeaders,
		AllowMethods:     allowedMethods,
		AllowOriginsFunc: func(origin string) bool {
			return domainRegex.MatchString(origin)
		},
	}))

	// Attach all the middleware
	addMiddleware(app)

	// Connect to database
	connectDataBase()

	return app

}

/*
Gracefully Shutdown The Application

The server will wait for all the active connections to process, and will not accept new connections.
*/
func GracefulShutdownHandler(app *fiber.App, shutdown chan os.Signal) {

	Logger := configuration.GetLogger()
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-shutdown
		Logger.Info("Shutting down gracefully...")

		// Close any other resources or perform cleanup before shutting down the app
		if exception := app.Shutdown(); exception != nil {
			Logger.Error(exception, "Error during shutdown")
		}

		// Close the shutdown channel to signal that the shutdown process is complete
		close(shutdown)
	}()

}

/*
Performs cleanup tasks before shutdown and after program exits.

It closes all the db connections and redis connections as of now.
*/
func InitiateCleanupProcess() {

	Logger := configuration.GetLogger()

	// Closing all DB connections
	sqlite := database.DBManager.SQLiteDB

	if sqlite != nil {

		closingError := sqlite.Close()

		if closingError != nil {
			Logger.Error(closingError, "Error while closing SQLite db connection")

		} else {
			Logger.Info("SQLite db connection closed")
		}

	}

}
