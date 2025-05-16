package models

import "time"

type TimeSlot struct {
	ID         int          `json:"id" gorm:"primaryKey;autoIncrement"`
	ActivityID int          `json:"activity_id"`
	Weekday    time.Weekday `json:"weekday"`
	StartTime  string       `json:"start_time" gorm:"not null"`
	EndTime    string       `json:"end_time" gorm:"not null"`
}
