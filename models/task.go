package models

import (
	"time"

	"gorm.io/gorm"
)

// Task represents the task model
type Task struct {
	ID          int            `json:"id"`
	UserID      int            `json:"user_id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Priority    int            `json:"priority"`
	DueDate     time.Time      `json:"due_date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   gorm.DeletedAt `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
