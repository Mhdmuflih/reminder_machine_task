package repositories

import (
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
// UpdateTask updates a task by ID
func UpdateTask(id uint, updatedData map[string]interface{}) (models.Task, error) {
	var task models.Task

	// Find the task
	if err := config.DB.First(&task, id).Error; err != nil {
		return task, err
	}

	// Update fields
	if err := config.DB.Model(&task).Updates(updatedData).Error; err != nil {
		return task, err
	}

	// Fetch the updated task
	if err := config.DB.First(&task, id).Error; err != nil {
		return task, err
	}

	return task, nil
}


// ===============================================================================================
// DeleteTask deletes a task by ID
func DeleteTask(id uint) error {
	var task models.Task

	// Find the task first
	if err := config.DB.First(&task, id).Error; err != nil {
		return err
	}

	// Delete the task
	if err := config.DB.Delete(&task).Error; err != nil {
		return err
	}

	return nil
}