package models

import "time"

type Activity struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Day          string    `json:"day"`
	Cupo         int       `json:"cupo"`
	Schedule     time.Time `json:"schedule"`
	Category     string    `json:"category"`
	Instructor   string    `json:"instructor"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Review       string    `json:"review"`
	CreationDate time.Time `json:"creationDate" gorm:"autoCreateTime"`
}
