package services

import (
	"reminder/models"
	"reminder/repositories"
)

func CreateRuleService(name string, minutesBefore int) (models.ReminderRule, error) {
	rule := models.ReminderRule{
		Name:          name,
		MinutesBefore: minutesBefore,
		IsActive:      true,
	}
	return repositories.SaveRule(rule)
}


// GetAllRulesService calls repository to fetch all rules
func GetAllRulesService() ([]models.ReminderRule, error) {
	return repositories.GetAllRules()
}

// UpdateRuleService updates a rule
func UpdateRuleService(id uint, name string, minutesBefore int, isActive *bool) (models.ReminderRule, error) {
	updatedData := map[string]interface{}{
		"name":           name,
		"minutes_before": minutesBefore,
	}

	// Include isActive if provided
	if isActive != nil {
		updatedData["is_active"] = *isActive
	}

	return repositories.UpdateRule(id, updatedData)
}


// ActivateRuleService calls repository to activate rule
func ActivateRuleService(id uint) (models.ReminderRule, error) {
	return repositories.ActivateRule(id)
}

// DeactivateRuleService calls repository to deactivate rule
func DeactivateRuleService(id uint) (models.ReminderRule, error) {
	return repositories.DeactivateRule(id)
}


// DeleteRuleService calls repository to delete a rule
func DeleteRuleService(id uint) error {
	return repositories.DeleteRule(id)
}