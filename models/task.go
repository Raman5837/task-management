package models

import "github.com/uptrace/bun"

// Represents a task entity
type Task struct {
	AbstractModel
	bun.BaseModel `bun:"table:tasks"`
	ID            int `bun:"id,pk,autoincrement"`

	Title       string `bun:"title,notnull"`
	Description string `bun:"description,nullzero"`
	Status      string `bun:"status,notnull" json:"status"`
}
