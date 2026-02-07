package repositories

import (
	"errors"
	"reminder/config"
	"reminder/models"
)

// ===============================================================================================
// SaveTask inserts a new task into the database
func SaveTask(task models.Task) (models.Task, error) {
	err := config.DB.Create(&task).Error
	return task, err
}

// ===============================================================================================
// GetAllTasks fetches all tasks from the database
func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := config.DB.Find(&tasks).Error
	return tasks, err
}

// ===============================================================================================
// Private helper: get task by ID (reusable)
func getTaskByID(id uint) (models.Task, error) {
	var task models.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		return task, errors.New("Task not found")
	}
	return task, nil
}

// ===============================================================================================
// UpdateTask updates a task by ID
func UpdateTask(id uint, updatedData map[string]interface{}) (models.Task, error) {
	task, err := getTaskByID(id)
	if err != nil {
		return task, err
	}

	if err := config.DB.Model(&task).Updates(updatedData).Error; err != nil {
		return task, err
	}

	return getTaskByID(id) // return updated task
}

// ===============================================================================================
// DeleteTask deletes a task by ID
func DeleteTask(id uint) error {
	_, err := getTaskByID(id)
	if err != nil {
		return err
	}

	return config.DB.Delete(&models.Task{}, id).Error
}