//Estructuras que representan entidades del dominio
//Define structs que representan entidades como User, Activity, etc
//Solo contienen la estrutura d elos datos, no logica

package models

import (
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model                 // Incluye ID, CreatedAt, UpdatedAt, DeletedAt
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	Day          string        `json:"day"`
	Duration     int           `json:"duration"` // Duración en minutos
	Category     string        `json:"category"`
	State        string        `json:"state"`
	Instructor   string        `json:"instructor"`
	Cupo         int           `json:"cupo"`
	Inscriptions []Inscription `gorm:"foreignKey:ActivityID"`
}
