package models

import "time"

type Usuario struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	LastName     string    `json:"lastName"`
	DNI          uint      `json:"DNI"`
	Email        string    `json:"email"`
	Password     string    `json:"**********"`
	Rol          string    `json:"rol"`
	DateCreation time.Time `json:"fechaCreacion" gorm:"autoCreateTime"`
}
