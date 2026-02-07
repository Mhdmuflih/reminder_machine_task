package scheduler

import (
    "fmt"
    "time"
    "reminder/services"
)

func StartReminderScheduler() {
    ticker := time.NewTicker(10 * time.Second) // 10s for demo, 1 * time.Minute for prod
    go func() {
        for range ticker.C {
            fmt.Println("Scheduler running at:", time.Now())

            tasks, _ := services.GetAllTasksService()
            rules, _ := services.GetAllRulesService()

            for _, task := range tasks {
                for _, rule := range rules {
                    if rule.IsActive {
                        reminderTime := task.DueAt.Add(-time.Minute * time.Duration(rule.MinutesBefore))

                        // Trigger only if reminder time <= now
                        if time.Now().After(reminderTime) {
                            fmt.Println("Reminder triggered for task:", task.Title, "at", reminderTime)
                            services.CreateLogService(
                                fmt.Sprintf("Reminder triggered: Task=%s, Rule=%s", task.Title, rule.Name),
                                "system",
                            )
                        }
                    }
                }
            }
        }
    }()
}