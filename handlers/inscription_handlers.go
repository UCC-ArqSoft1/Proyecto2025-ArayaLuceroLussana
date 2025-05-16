package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Inscription struct {
	ID         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID     uint   `json:"user_ID"`
	ActivityID uint   `json:"activity_ID"`
	Date       string `json:"date"` // Fecha especifica de la sesion
	//IsActivo         bool      `json:"estado"`
	StartTime string `json:"start_time"` //Hora de inicio  Ej: "18:00"
	EndTime   string `json:"end_time"`   //Hora de finalizacion  Ej: "19:00"
}

func createInscription(ctx *gin.Context) {
	var newInscription Inscription

	if err := ctx.BindJSON(&newInscription); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	db.Create(&newInscription)
	ctx.IndentedJSON(http.StatusCreated, newInscription)
}
