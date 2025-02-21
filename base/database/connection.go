package database

import (
	"context"
	"fmt"
	"time"

	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	_ "github.com/uptrace/bun/driver/sqliteshim"

	"github.com/Raman5837/task-management/base/configuration"
)

// Connect to database
func EstablishConnection() error {

	Logger := configuration.GetLogger()

	// SQLite Database Configuration
	config := SQLiteConfig{
		MaxOpenConnection: 25,
		MaxIdleConnection: 10,
		DBPath:            "storage.db",
		Timeout:           5 * time.Second,
	}
	connection, err := sql.Open("sqlite", config.DBPath)

	if err != nil {
		return fmt.Errorf("Failed to open SQLite DB: %w", err)
	}

	DBManager.SQLiteDB = bun.NewDB(connection, sqlitedialect.New())

	// Configure the connection pool.
	DBManager.SQLiteDB.SetConnMaxLifetime(config.Timeout)
	DBManager.SQLiteDB.SetMaxOpenConns(config.MaxOpenConnection)
	DBManager.SQLiteDB.SetMaxIdleConns(config.MaxIdleConnection)

	// Verify the connection with a ping.
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()

	if err := DBManager.SQLiteDB.PingContext(ctx); err != nil {
		return fmt.Errorf("Failed to ping SQLite DB: %w", err)
	}

	Logger.Info("SQLite connection established successfully")

	if err := Migrate(); err != nil {
		Logger.Fatal(err, "Something went wrong while migrating tables")
	}

	return nil

}
