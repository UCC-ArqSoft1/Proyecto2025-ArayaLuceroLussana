package models

import (
	"gorm.io/gorm"
)

// Una relación entre un usuario y una actividad. Es decir, el registro que indica que un socio se inscribió a una actividad.
type Inscription struct {
	gorm.Model        // Incluye ID, CreatedAt, UpdatedAt, DeletedAt
	UserID     uint   `json:"user_ID"`                               //fk que apunta al usuario que se inscribe
	ActivityID uint   `json:"activity_ID" gorm:"column:activity_id"` //fk que apunta a la actividad seleccionada
	Date       string `json:"date"`                                  // Fecha especifica de la sesion
	State      string `json:"state"`
}
