package services

import (
	"reminder/models"
	"reminder/repositories"
	"time"
)


// ===============================================================================================
// CreateLogService creates a new log entry
func CreateLogService(action string, user string) (models.AuditLog, error) {
	log := models.AuditLog{
		Action:    action,
		User:      user,
		CreatedAt: time.Now(),
	}
	return repositories.SaveLog(log)
}


// ===============================================================================================
// GetAllLogsService returns all logs
func GetAllLogsService() ([]models.AuditLog, error) {
	return repositories.GetAllLogs()
}