// controllers/task_controller.go
package controllers

import (
	"time"

	"github.com/vikash-parashar/task-remainder/config"
	"github.com/vikash-parashar/task-remainder/models"
	"gorm.io/gorm"
)

// GetAllTasksForUser retrieves all tasks for a specific user from the database
func GetAllTasksForUser(userID uint) ([]*models.Task, error) {
	var tasks []*models.Task
	err := config.DB.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

// GetTaskByIDForUser retrieves a task by ID for a specific user from the database
func GetTaskByIDForUser(taskID uint, userID uint) (*models.Task, error) {
	var task models.Task
	err := config.DB.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &models.Task{}, gorm.ErrRecordNotFound
		}
		return &models.Task{}, err
	}
	return &task, nil
}

// CreateTaskForUser inserts a new task for a specific user into the database
func CreateTaskForUser(task *models.Task) error {
	return config.DB.Create(task).Error
}

// UpdateTaskForUser updates an existing task by ID for a specific user in the database (partial update with PATCH)
func UpdateTaskForUser(taskID uint, userID uint, updatedTask *models.Task) error {
	// Check if the task exists for the user
	existingTask, err := GetTaskByIDForUser(taskID, userID)
	if err != nil {
		return err
	}

	// Prepare the update data
	updateData := make(map[string]interface{})

	if updatedTask.Title != "" {
		updateData["title"] = updatedTask.Title
	}
	if updatedTask.Description != "" {
		updateData["description"] = updatedTask.Description
	}
	if updatedTask.Priority != 0 {
		updateData["priority"] = updatedTask.Priority
	}
	if !updatedTask.DueDate.IsZero() {
		updateData["due_date"] = updatedTask.DueDate
	}
	// Omit fields with nil values
	if updatedTask.CreatedAt != (time.Time{}) {
		updateData["created_at"] = updatedTask.CreatedAt
	}

	// Update the task
	err = config.DB.Model(existingTask).Where("id = ? AND user_id = ?", taskID, userID).Updates(updateData).Error
	return err
}

// DeleteTaskForUser deletes a task by ID for a specific user from the database
func DeleteTaskForUser(taskID uint, userID uint) error {
	// Check if the task exists for the user
	_, err := GetTaskByIDForUser(taskID, userID)
	if err != nil {
		return err
	}

	return config.DB.Where("id = ? AND user_id = ?", taskID, userID).Delete(&models.Task{}).Error
}
