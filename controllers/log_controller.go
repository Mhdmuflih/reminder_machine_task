package controllers

import (
	"net/http"
	"reminder/services"
	"github.com/gin-gonic/gin"
)

// GetLogs handles GET /logs
func GetLogs(c *gin.Context) {
	logs, err := services.GetAllLogsService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"logs": logs})
}