package models

import "time"

// AbstractModel provides common metadata columns for all models.
// Embed this struct into models to include these columns automatically.
type AbstractModel struct {
	IsDeleted  bool      `bun:"is_deleted,default:false"`
	CreatedAt  time.Time `bun:"created_at,default:current_timestamp"`
	ModifiedAt time.Time `bun:"modified_at,default:current_timestamp"`
}
