package routes

import (
	"reminder/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRuleRoutes(router *gin.Engine) {

	rules := router.Group("/rules") 
	
	{
		rules.GET("/", controllers.GetRules);
		rules.POST("/", controllers.CreateRule);
		rules.PUT("/:id", controllers.UpdateRule);
		rules.PATCH("/:id/active", controllers.ActivateRule);
		rules.PATCH("/:id/deactivate", controllers.DeactivateRule);
		rules.DELETE("/:id", controllers.DeleteRule);
	}
}