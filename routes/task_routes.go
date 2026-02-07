package routes

import (
	"reminder/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(router *gin.Engine) {
	task := router.Group("/tasks")

	{
		task.GET("/", controllers.GetTasks);
		task.POST("/", controllers.CreateTask);
		task.PUT("/:id", controllers.UpdateTask);
		task.DELETE("/:id", controllers.DeleteTask);
	}
}