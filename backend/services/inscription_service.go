package services

import (
	"alua/config"
	"alua/models"
	"errors"
	"time"
)

// Crea inscripcion a una actividad
func CreateInscription(UserID uint, ActivityID uint) error {

	// Verify if the user is already registered for the activity
	var existing models.Inscription
	err := config.DB.Where("user_id = ? AND activity_id = ?", UserID, ActivityID).First(&existing).Error
	if err == nil {
		return errors.New("user already registered for this activity")
	}

	// Verifica que la actividad exista y haya lugar
	var activity models.Activity
	if err := config.DB.First(&activity, ActivityID).Error; err != nil {
		return errors.New("activity not found")
	}

	// Verifica si la actividad tiene estado activo
	if activity.State != "Activo" {
		return errors.New("inscription couldn't be done: activity is not active")
	}

	// Cuenta el numero de registrados para la actividad
	var totalInscription int64
	config.DB.Model(&models.Inscription{}).Where("activity_id = ?", ActivityID).Count(&totalInscription)
	if totalInscription >= int64(activity.Cupo) {
		return errors.New("there are no spots available for this activity")
	}

	// Crea una nueva inscripcion
	new := models.Inscription{
		UserID:     UserID,
		ActivityID: ActivityID,
		Date:       time.Now().Format("2006-01-02"),
		State:      "Confirmed", // Default state
	}

	return config.DB.Create(&new).Error
}

// Cambia el estado de la inscripcion
func EditInscription(id uint, new models.Inscription, UserID uint) error {
	var inscription models.Inscription

	// Busca una inscripcion por ID
	if err := config.DB.First(&inscription, id).Error; err != nil {
		return errors.New("inscription not found")
	}

	// Verifica que la inscripcion sea del usuario
	if inscription.UserID != UserID {
		return errors.New("no permission to edit this inscription")
	}

	// Cambia el estado de la inscripcion
	inscription.State = new.State

	return config.DB.Save(&inscription).Error
}

// Borra una inscripcion
func DeleteInscription(activityID uint, userID uint) error {
	var inscription models.Inscription

	// Buscar la inscripción que coincida con el userID y activityID
	if err := config.DB.Where("user_id = ? AND activity_id = ?", userID, activityID).First(&inscription).Error; err != nil {
		return errors.New("inscripción no encontrada")
	}

	//Cambia el estado de la inscripción a "cancelado"
	inscription.State = "Cancelled"
	if err := config.DB.Save(&inscription).Error; err != nil {
		return errors.New("error al cancelar la inscripción")
	}

	// Eliminar la inscripción encontrada
	if err := config.DB.Delete(&inscription).Error; err != nil {
		return errors.New("error al eliminar la inscripción")
	}

	return nil
}
