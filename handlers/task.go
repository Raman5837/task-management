package handlers

import (
	"strconv"
	"sync"

	"github.com/Raman5837/task-management/base/configuration"
	"github.com/Raman5837/task-management/base/constants"
	"github.com/Raman5837/task-management/base/database"
	"github.com/Raman5837/task-management/base/middleware"
	"github.com/Raman5837/task-management/base/utils"
	"github.com/Raman5837/task-management/repository"
	"github.com/Raman5837/task-management/services"
	"github.com/Raman5837/task-management/types"
	"github.com/gofiber/fiber/v2"
)

var Logger = configuration.GetLogger()

// Handler For Create Task
func CreateTaskHandler(context *fiber.Ctx) error {

	payload := types.CreateTaskRequestEntity{}
	if err := context.BodyParser(&payload); err != nil {
		Logger.Warn("[CreateTaskHandler]: Error while parsing request payload %v", err)
		return utils.SendErrorResponse(context, "Invalid payload", fiber.StatusBadRequest)
	}

	if payload.Status != "" && !payload.Status.IsValid() {
		Logger.Warn("[CreateTaskHandler]: Invalid value for status")
		return utils.SendErrorResponse(context, "Invalid payload", fiber.StatusBadRequest)
	}

	payload.Context = context.Context()
	service := services.NewTaskService(repository.NewTaskRepository(database.DBManager.SQLiteDB))

	response, err := service.CreateTask(payload)

	if err != nil {
		Logger.Warn("[CreateTaskHandler]: Error creating new task %v", err)
		return utils.SendErrorResponse(context, "Something went wrong", fiber.StatusInternalServerError)
	}

	return utils.SendSuccessResponse(context, "Task created successfully", response, fiber.StatusCreated)

}

// Handler For Get Task
func GetTaskHandler(context *fiber.Ctx) error {

	taskID, err := strconv.Atoi(context.Params("id"))
	if err != nil {
		Logger.Warn("[GetTaskHandler]: Error while parsing request payload %v", err)
		return utils.SendErrorResponse(context, "Invalid task ID", fiber.StatusBadRequest)
	}

	service := services.NewTaskService(repository.NewTaskRepository(database.DBManager.SQLiteDB))
	response, err := service.GetTask(types.GetTaskRequestEntity{Context: context.Context(), ID: taskID})

	if err != nil {
		Logger.Warn("[GetTaskHandler]: Error while retrieving task %v", err)
		return utils.SendErrorResponse(context, "Task does not exists", fiber.StatusNotFound)
	}

	return utils.SendSuccessResponse(context, "Task retrieved successfully", response, fiber.StatusOK)

}

// Handler For Update Task
func UpdateTaskHandler(context *fiber.Ctx) error {

	taskID, err := strconv.Atoi(context.Params("id"))
	if err != nil {
		Logger.Warn("Error while parsing request payload %v", err)
		return utils.SendErrorResponse(context, "Invalid task ID", fiber.StatusBadRequest)
	}

	payload := types.UpdateTaskRequestEntity{}
	if err := context.BodyParser(&payload); err != nil {
		return utils.SendErrorResponse(context, "Invalid payload", fiber.StatusBadRequest)
	}

	if payload.Status != "" && !payload.Status.IsValid() {
		Logger.Warn("[UpdateTaskHandler]: Invalid value for status")
		return utils.SendErrorResponse(context, "Invalid payload", fiber.StatusBadRequest)
	}

	payload.ID = taskID
	payload.Context = context.Context()

	service := services.NewTaskService(repository.NewTaskRepository(database.DBManager.SQLiteDB))
	response, err := service.UpdateTask(payload)

	if err != nil {
		Logger.Warn("[UpdateTaskHandler]: Error while updating task %v", err)
		return utils.SendErrorResponse(context, "Failed to update task", fiber.StatusInternalServerError)
	}

	return utils.SendSuccessResponse(context, "Task updated successfully", response, fiber.StatusOK)

}

// Handler For Delete Task
func DeleteTaskHandler(context *fiber.Ctx) error {

	taskID, err := strconv.Atoi(context.Params("id"))
	if err != nil {
		Logger.Warn("[DeleteTaskHandler]: Error while parsing request payload %v", err)
		return utils.SendErrorResponse(context, "Invalid task ID", fiber.StatusBadRequest)
	}

	payload := types.GetTaskRequestEntity{Context: context.Context(), ID: taskID}
	service := services.NewTaskService(repository.NewTaskRepository(database.DBManager.SQLiteDB))
	err = service.DeleteTask(payload)

	if err != nil {
		Logger.Warn("[DeleteTaskHandler]: Error while deleting task %v", err)
		return utils.SendErrorResponse(context, "Failed to delete task", fiber.StatusInternalServerError)
	}

	return utils.SendSuccessResponse(context, "Task deleted successfully", nil, fiber.StatusOK)

}

// Handler To List Task
func ListTaskHandler(context *fiber.Ctx) error {

	status := constants.TaskStatus(context.Query("status", ""))
	if status != "" && !status.IsValid() {
		return utils.SendErrorResponse(context, "Invalid payload", fiber.StatusBadRequest)
	}

	pagination := context.Locals("paginationOptions").(middleware.OffsetPaginationRequestOptions)
	payload := types.FilterTaskRequestEntity{
		Status:  status,
		Context: context.Context(),
		Limit:   &pagination.Limit,
		Offset:  &pagination.Offset,
	}

	service := services.NewTaskService(repository.NewTaskRepository(database.DBManager.SQLiteDB))

	var count int
	var taskErr error
	var countErr error
	var waitGroup sync.WaitGroup
	var allTasks []types.TaskResponseEntity

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		allTasks, taskErr = service.ListTask(payload)
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		count, countErr = service.GetCountOfTask(payload)
	}()

	waitGroup.Wait()

	if taskErr != nil {
		Logger.Warn("[ListTaskHandler]: Error while retrieving all tasks %v", taskErr)
		return utils.SendErrorResponse(context, "Failed to fetch tasks", fiber.StatusInternalServerError)
	}
	if countErr != nil {
		Logger.Warn("[ListTaskHandler]: Error while calculating count %v", countErr)
		return utils.SendErrorResponse(context, "Failed to fetch task count", fiber.StatusInternalServerError)
	}

	return middleware.LimitOffsetBasedPaginatedResponse(context, count, allTasks)

}
