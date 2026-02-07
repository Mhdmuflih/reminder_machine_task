package repositories

import (
	"errors"
	"reminder/config"
	"reminder/models"
)

// ===============================================================================================
// SaveRule inserts a new rule into the database
func SaveRule(rule models.ReminderRule) (models.ReminderRule, error) {
	err := config.DB.Create(&rule).Error
	return rule, err
}

// ===============================================================================================
// GetAllRules fetches all reminder rules from DB
func GetAllRules() ([]models.ReminderRule, error) {
	var rules []models.ReminderRule
	err := config.DB.Find(&rules).Error
	return rules, err
}

// ===============================================================================================
// Private helper: get rule by ID (reusable)
func getRuleByID(id uint) (models.ReminderRule, error) {
	var rule models.ReminderRule
	if err := config.DB.First(&rule, id).Error; err != nil {
		return rule, errors.New("ReminderRule not found")
	}
	return rule, nil
}

// ===============================================================================================
// UpdateRule updates a reminder rule by ID
func UpdateRule(id uint, updatedData map[string]interface{}) (models.ReminderRule, error) {
	rule, err := getRuleByID(id)
	if err != nil {
		return rule, err
	}

	if err := config.DB.Model(&rule).Updates(updatedData).Error; err != nil {
		return rule, err
	}

	return getRuleByID(id) // return updated rule
}

// ===============================================================================================
// ActivateRule sets is_active = true for the given rule ID
func ActivateRule(id uint) (models.ReminderRule, error) {
	rule, err := getRuleByID(id)
	if err != nil {
		return rule, err
	}

	if err := config.DB.Model(&rule).Update("is_active", true).Error; err != nil {
		return rule, err
	}

	return getRuleByID(id)
}

// ===============================================================================================
// DeactivateRule sets is_active = false for the given rule ID
func DeactivateRule(id uint) (models.ReminderRule, error) {
	rule, err := getRuleByID(id)
	if err != nil {
		return rule, err
	}

	if err := config.DB.Model(&rule).Update("is_active", false).Error; err != nil {
		return rule, err
	}

	return getRuleByID(id)
}

// ===============================================================================================
// DeleteRule deletes a reminder rule by ID
func DeleteRule(id uint) error {
	_, err := getRuleByID(id)
	if err != nil {
		return err
	}

	return config.DB.Delete(&models.ReminderRule{}, id).Error
}