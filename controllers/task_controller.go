package controllers

import (
	"net/http"
	"reminder/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// ===============================================================================================
// CreateTask handles POST /tasks
func CreateTask(c *gin.Context) {
	var input struct {
		Title  string `json:"title"`
		DueAt  string `json:"due_at"`  // Accept string in ISO format
		Status string `json:"status"`  // e.g., "pending", "done"
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse DueAt string to time.Time
	dueAt, err := time.Parse(time.RFC3339, input.DueAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid due_at format, use ISO8601"})
		return
	}

	task, err := services.CreateTaskService(input.Title, dueAt, input.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	    // ✅ Create log automatically
    services.CreateLogService("Created Task: "+task.Title, "Muflih")

	c.JSON(http.StatusOK, gin.H{"task": task})
}


// ===============================================================================================
// GetTasks handles GET /tasks
func GetTasks(c *gin.Context) {
	tasks, err := services.GetAllTasksService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}


// ===============================================================================================
// UpdateTask handles PUT /tasks/:id
func UpdateTask(c *gin.Context) {
	// Get ID from URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	// Parse JSON body
	var input struct {
		Title  string `json:"title"`
		DueAt  string `json:"due_at"` // string in ISO8601
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert DueAt string to time.Time
	dueAt, err := time.Parse(time.RFC3339, input.DueAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid due_at format, use ISO8601"})
		return
	}

	// Call service
	updatedTask, err := services.UpdateTaskService(uint(id), input.Title, dueAt, input.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": updatedTask})
}



// ===============================================================================================
// DeleteTask handles DELETE /tasks/:id
func DeleteTask(c *gin.Context) {
	// Get ID from URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	// Call service
	if err := services.DeleteTaskService(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
