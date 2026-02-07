package models

import "time"

type AuditLog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Action    string    `json:"action"`
	User      string    `json:"user"`
	CreatedAt time.Time `json:"created_at"`
}