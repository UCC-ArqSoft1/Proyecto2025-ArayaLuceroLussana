package models

import (
	"gorm.io/gorm"
)

// Una relación entre un usuario y una actividad. Es decir, el registro que indica que un socio se inscribió a una actividad.
type Inscription struct {
	gorm.Model
	UserID     uint   `json:"user_id" gorm:"column:user_id"`
	ActivityID uint   `json:"activity_id" gorm:"column:activity_id"`
	Date       string `json:"date"`
	State      string `json:"state"`
}
