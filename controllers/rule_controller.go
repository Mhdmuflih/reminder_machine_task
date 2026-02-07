package controllers

import (
	"net/http"
	"reminder/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ===============================================================================================
func CreateRule(c *gin.Context) {
	var input struct {
		Name          string `json:"name"`
		MinutesBefore int    `json:"minutes_before"`
	}

	// Parse JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	// Call service
	rule, err := services.CreateRuleService(input.Name, input.MinutesBefore)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Rule created successfully",
		"rule":    rule,
	})
}


// ===============================================================================================
// GetRules handles GET /rules
func GetRules(c *gin.Context) {
	rules, err := services.GetAllRulesService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rules": rules})
}


// ===============================================================================================
// UpdateRule handles PUT /rules/:id
func UpdateRule(c *gin.Context) {
	// Get ID from URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rule ID"})
		return
	}

	// Parse JSON body
	var input struct {
		Name          string `json:"name"`
		MinutesBefore int    `json:"minutes_before"`
		IsActive      *bool  `json:"is_active"` // optional
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service
	updatedRule, err := services.UpdateRuleService(uint(id), input.Name, input.MinutesBefore, input.IsActive)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rule": updatedRule})
}


// ===============================================================================================
// ActivateRule handles PATCH /rules/:id/active
func ActivateRule(c *gin.Context) {
	// Get ID from URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rule ID"})
		return
	}

	// Call service
	updatedRule, err := services.ActivateRuleService(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rule": updatedRule})
}


// ===============================================================================================
// DeactivateRule handles PATCH /rules/:id/deactivate
func DeactivateRule(c *gin.Context) {
	// Get ID from URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rule ID"})
		return
	}

	// Call service
	updatedRule, err := services.DeactivateRuleService(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rule": updatedRule})
}


// ===============================================================================================
// DeleteRule handles DELETE /rules/:id
func DeleteRule(c *gin.Context) {
	// Get ID from URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rule ID"})
		return
	}

	// Call service
	if err := services.DeleteRuleService(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Rule deleted successfully"})
}