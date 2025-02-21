package repository

import (
	"context"

	"github.com/Raman5837/task-management/base/configuration"
	"github.com/Raman5837/task-management/models"
	"github.com/uptrace/bun"
)

var Logger = configuration.GetLogger()

type TaskRepository struct {
	DB *bun.DB
}

func NewTaskRepository(DB *bun.DB) *TaskRepository {
	return &TaskRepository{DB: DB}
}

// Create new Task
func (repo *TaskRepository) CreateTask(context context.Context, instance *models.Task) (*models.Task, error) {

	_, err := repo.DB.NewInsert().Model(instance).Exec(context)
	return instance, err

}

// Returns Task basis ID
func (repo *TaskRepository) GetTask(context context.Context, taskID int) (*models.Task, error) {

	instance := &models.Task{}
	err := repo.DB.NewSelect().Model(instance).Where("id = ?", taskID).Scan(context)
	return instance, err

}

// Update Task basis ID
func (repo *TaskRepository) UpdateTask(context context.Context, task *models.Task) error {
	_, err := repo.DB.NewUpdate().Model(task).WherePK().Exec(context)
	return err
}

// Deletes Task basis ID
func (repo *TaskRepository) DeleteTask(context context.Context, taskID int) error {

	_, err := repo.DB.NewDelete().Model((*models.Task)(nil)).Where("id = ?", taskID).Exec(context)
	return err
}

// Filter Task
func (repo *TaskRepository) FilteredTask(context context.Context, status string, limit *int, offset *int) ([]models.Task, error) {

	var response []models.Task
	query := repo.DB.NewSelect().Model(&response)

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if limit != nil && *limit > 0 {
		query = query.Limit(*limit)
	}

	if offset != nil && *offset >= 0 {
		query = query.Offset(*offset)
	}

	err := query.Scan(context)
	return response, err
}

// Returns the total number of tasks (For pagination)
func (repo *TaskRepository) GetCount(context context.Context, status string) (int, error) {

	query := repo.DB.NewSelect().Model((*models.Task)(nil)).Where("status = ?", status)

	if status != "" {
		query = query.Where("status = ?", status)
	}

	return query.Count(context)
}
