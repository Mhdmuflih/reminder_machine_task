package repositories

import (
	"reminder/config"
	"reminder/models"
)

func SaveRule(rule models.ReminderRule) (models.ReminderRule, error) {
	err := config.DB.Create(&rule).Error
	return rule, err
}


// GetAllRules fetches all reminder rules from DB
func GetAllRules() ([]models.ReminderRule, error) {
	var rules []models.ReminderRule
	err := config.DB.Find(&rules).Error
	return rules, err
}


// UpdateRule updates a reminder rule by ID
func UpdateRule(id uint, updatedData map[string]interface{}) (models.ReminderRule, error) {
	var rule models.ReminderRule

	// Find the rule first
	if err := config.DB.First(&rule, id).Error; err != nil {
		return rule, err
	}

	// Update fields
	if err := config.DB.Model(&rule).Updates(updatedData).Error; err != nil {
		return rule, err
	}

	// Fetch the updated rule
	if err := config.DB.First(&rule, id).Error; err != nil {
		return rule, err
	}

	return rule, nil
}



// ActivateRule sets is_active = true for the given rule ID
func ActivateRule(id uint) (models.ReminderRule, error) {
	var rule models.ReminderRule

	// Find the rule
	if err := config.DB.First(&rule, id).Error; err != nil {
		return rule, err
	}

	// Update is_active to true
	if err := config.DB.Model(&rule).Update("is_active", true).Error; err != nil {
		return rule, err
	}

	// Fetch updated rule
	if err := config.DB.First(&rule, id).Error; err != nil {
		return rule, err
	}

	return rule, nil
}



// DeactivateRule sets is_active = false for the given rule ID
func DeactivateRule(id uint) (models.ReminderRule, error) {
	var rule models.ReminderRule

	// Find the rule
	if err := config.DB.First(&rule, id).Error; err != nil {
		return rule, err
	}

	// Update is_active to false
	if err := config.DB.Model(&rule).Update("is_active", false).Error; err != nil {
		return rule, err
	}

	// Fetch updated rule
	if err := config.DB.First(&rule, id).Error; err != nil {
		return rule, err
	}

	return rule, nil
}


// DeleteRule deletes a reminder rule by ID
func DeleteRule(id uint) error {
	var rule models.ReminderRule

	// Find the rule
	if err := config.DB.First(&rule, id).Error; err != nil {
		return err
	}

	// Delete the rule
	if err := config.DB.Delete(&rule).Error; err != nil {
		return err
	}

	return nil
}