package database

import (
	"time"

	"github.com/uptrace/bun"
)

// DatabaseManager to handle multiple database connections
type DatabaseManager struct {
	SQLiteDB *bun.DB
}

var DBManager = &DatabaseManager{}

type SQLiteConfig struct {
	MaxOpenConnection int
	MaxIdleConnection int
	DBPath            string
	Timeout           time.Duration
}
