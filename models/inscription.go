package models

type Inscription struct {
	ID         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID     uint   `json:"user_ID"`
	ActivityID uint   `json:"activity_ID"`
	Date       string `json:"date"` // Fecha especifica de la sesion
	//IsActivo         bool      `json:"estado"`
	StartTime string `json:"start_time"` //Hora de inicio  Ej: "18:00"
	EndTime   string `json:"end_time"`   //Hora de finalizacion  Ej: "19:00"
}
