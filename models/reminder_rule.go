package models

type ReminderRule struct {
	ID            uint `gorm:"primaryKey"`
	Name          string
	MinutesBefore int
	IsActive      bool
}