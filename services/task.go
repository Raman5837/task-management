package services

import (
	"github.com/Raman5837/task-management/models"
	"github.com/Raman5837/task-management/repository"
	"github.com/Raman5837/task-management/types"
)

type TaskService struct {
	repo *repository.TaskRepository
}

// Initializes new TaskService
func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

// Create new Task
func (service *TaskService) CreateTask(payload types.CreateTaskRequestEntity) (*types.TaskResponseEntity, error) {

	request := &models.Task{Title: payload.Title, Status: string(payload.Status), Description: payload.Description}
	instance, err := service.repo.CreateTask(payload.Context, request)

	if err != nil {
		return nil, err
	}

	response := &types.TaskResponseEntity{
		ID:          instance.ID,
		Title:       instance.Title,
		Status:      instance.Status,
		Description: instance.Description,
	}

	return response, nil
}

// Retrieve a Task using ID
func (service *TaskService) GetTask(payload types.GetTaskRequestEntity) (*types.TaskResponseEntity, error) {

	instance, err := service.repo.GetTask(payload.Context, payload.ID)

	if err != nil {
		return nil, err
	}

	response := &types.TaskResponseEntity{
		ID:          instance.ID,
		Title:       instance.Title,
		Status:      instance.Status,
		Description: instance.Description,
	}

	return response, nil
}

// Update an existing task
func (service *TaskService) UpdateTask(payload types.UpdateTaskRequestEntity) (*types.TaskResponseEntity, error) {

	instance := &models.Task{
		ID:          payload.ID,
		Title:       payload.Title,
		Description: payload.Description,
		Status:      string(payload.Status),
	}

	err := service.repo.UpdateTask(payload.Context, instance)
	if err != nil {
		return nil, err
	}

	return service.GetTask(types.GetTaskRequestEntity{ID: instance.ID, Context: payload.Context})
}

// Delete a Task using ID
func (service *TaskService) DeleteTask(payload types.GetTaskRequestEntity) error {

	return service.repo.DeleteTask(payload.Context, payload.ID)

}

// Returns list of tasks
func (service *TaskService) ListTask(payload types.FilterTaskRequestEntity) ([]types.TaskResponseEntity, error) {

	var response []types.TaskResponseEntity
	results, err := service.repo.FilteredTask(payload.Context, payload.Status, payload.Limit, payload.Offset)

	if err != nil {
		return nil, err
	}

	for _, task := range results {
		response = append(response, types.TaskResponseEntity{
			ID:          task.ID,
			Title:       task.Title,
			Status:      task.Status,
			Description: task.Description,
		})
	}

	return response, nil

}

// Returns the count of tasks
func (service *TaskService) GetCountOfTask(payload types.FilterTaskRequestEntity) (int, error) {

	return service.repo.GetCount(payload.Context, payload.Status)

}
