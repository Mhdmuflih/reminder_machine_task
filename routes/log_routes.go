package routes

import (
	"reminder/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterLogRoutes(router *gin.Engine) {
	router.GET("/logs", controllers.GetLogs)
}