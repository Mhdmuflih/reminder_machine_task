package models

import "time"

type Task struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	DueAt     time.Time `json:"due_at"`
	Status    string    `json:"status"` // e.g., "pending", "done"
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}