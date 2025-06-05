package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model                 // Incluye ID, CreatedAt, UpdatedAt, DeletedAt
	Name         string        `json:"name"`
	LastName     string        `json:"lastName"`
	DNI          uint          `json:"DNI"`
	Email        string        `json:"email"`
	Password     string        `json:"password"`          // Password hasheada, not shown in frontend
	Rol          string        `json:"rol"`               //Admin o Socio
	Inscriptions []Inscription `gorm:"foreignKey:UserID"` // Relación con Inscription
}
