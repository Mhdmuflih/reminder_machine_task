package services

import (
	"reminder/models"
	"reminder/repositories"
	"time"
)


// ===============================================================================================
// CreateTaskService creates a new task
func CreateTaskService(title string, dueAt time.Time, status string) (models.Task, error) {
	task := models.Task{
		Title:  title,
		DueAt:  dueAt,
		Status: status,
	}
	return repositories.SaveTask(task)
}


// ===============================================================================================
// GetAllTasksService returns all tasks
func GetAllTasksService() ([]models.Task, error) {
	return repositories.GetAllTasks()
}


// ===============================================================================================
// UpdateTaskService updates a task
func UpdateTaskService(id uint, title string, dueAt time.Time, status string) (models.Task, error) {
	updatedData := map[string]interface{}{
		"title":  title,
		"due_at": dueAt,
		"status": status,
	}

	return repositories.UpdateTask(id, updatedData)
}


// ===============================================================================================
// DeleteTaskService calls repository to delete a task
func DeleteTaskService(id uint) error {
	return repositories.DeleteTask(id)
}