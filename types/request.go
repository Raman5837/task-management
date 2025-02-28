package types

import (
	"context"

	"github.com/Raman5837/task-management/base/constants"
)

type CreateTaskRequestEntity struct {
	Context     context.Context      `query:"-" json:"-"`
	Title       string               `query:"title" json:"title"`
	Status      constants.TaskStatus `query:"status" json:"status"`
	Description string               `query:"description" json:"description"`
}

type UpdateTaskRequestEntity struct {
	Context     context.Context      `query:"-" json:"-"`
	ID          int                  `query:"id" json:"id"`
	Title       string               `query:"title" json:"title"`
	Status      constants.TaskStatus `query:"status" json:"status"`
	Description string               `query:"description" json:"description"`
}

type GetTaskRequestEntity struct {
	Context context.Context `query:"-" json:"-"`
	ID      int             `query:"id" json:"id"`
}

type FilterTaskRequestEntity struct {
	Context context.Context      `query:"-" json:"-"`
	Limit   *int                 `query:"limit" json:"limit"`
	Offset  *int                 `query:"offset" json:"offset"`
	Status  constants.TaskStatus `query:"status" json:"status"`
}
