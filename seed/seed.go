package seed

import (
	"fmt"
	"time"
	"reminder/config"
	"reminder/models"
)

// SeedTasks creates sample tasks for demo
func SeedTasks() {
	tasks := []models.Task{
		{Title: "Finish report", DueAt: time.Now().Add(30 * time.Minute), Status: "pending"},
		{Title: "Team meeting", DueAt: time.Now().Add(45 * time.Minute), Status: "pending"},
		{Title: "Call client", DueAt: time.Now().Add(1 * time.Hour), Status: "pending"},
		{Title: "Buy groceries", DueAt: time.Now().Add(2 * time.Hour), Status: "pending"},
		{Title: "Gym workout", DueAt: time.Now().Add(3 * time.Hour), Status: "pending"},
	}

	for _, task := range tasks {
		err := config.DB.Create(&task).Error
		if err != nil {
			fmt.Println("Failed to seed task:", task.Title, err)
		} else {
			fmt.Println("Seeded task:", task.Title)
		}
	}
}

// SeedReminderRules creates sample reminder rules
func SeedReminderRules() {
	rules := []models.ReminderRule{
		{Name: "30 min reminder", MinutesBefore: 30, IsActive: true},
		{Name: "15 min reminder", MinutesBefore: 15, IsActive: true},
	}

	for _, rule := range rules {
		err := config.DB.Create(&rule).Error
		if err != nil {
			fmt.Println("Failed to seed rule:", rule.Name, err)
		} else {
			fmt.Println("Seeded rule:", rule.Name)
		}
	}
}