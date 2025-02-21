package database

import (
	"context"
	"fmt"

	"github.com/Raman5837/task-management/models"
)

// Migrate all the tables
func Migrate() error {
	context := context.Background()

	// List all models that need to be migrated.
	modelsToMigrate := []interface{}{
		(*models.Task)(nil),
	}

	for _, model := range modelsToMigrate {

		_, err := DBManager.SQLiteDB.NewCreateTable().Model(model).IfNotExists().Exec(context)

		if err != nil {
			return fmt.Errorf("Failed to migrate model %T: %w", model, err)
		}
	}

	return nil
}
