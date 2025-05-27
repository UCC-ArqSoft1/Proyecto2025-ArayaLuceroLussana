package domain

// Representa el dominio del problema, con logica o interfaces de uso
import "time"

type Activity struct {
	ID           int       `json:"id"`
	Day          string    `json:"day"`
	Cupos        int       `json:"cups"`
	Schedule     string    `json:"schedule"`
	Category     string    `json:"category"`
	Instructor   string    `json:"instructor"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Review       string    `json:"review"`
	DateCreation time.Time `json:"date_creation" gorm:"autoCreateTime"`
}
