package repositories

import (
	"reminder/config"
	"reminder/models"
)

// SaveLog inserts a new log entry into the database
func SaveLog(log models.AuditLog) (models.AuditLog, error) {
	err := config.DB.Create(&log).Error
	return log, err
}

// GetAllLogs fetches all audit logs from the database
func GetAllLogs() ([]models.AuditLog, error) {
	var logs []models.AuditLog
	err := config.DB.Find(&logs).Error
	return logs, err
}