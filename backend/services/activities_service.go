// Contiene la logica del negocio
// logica de la app para funcionaar validaciones, reglas, calculos,etc
// Separa las reglas del negocio del acceso a datos y de la interfaz HTTP
package services

import (
	"alua/config"
	"alua/models"
	"errors"
)

// ValidateState checks if the provided state is valid.
var possibleStates = map[string]bool{
	"Activo":     true,
	"Inactivo":   true,
	"Finalizado": true,
}

// Get all the activities from the database
func ShowActivities() ([]models.Activity, error) {
	var activities []models.Activity
	result := config.DB.Find(&activities)
	return activities, result.Error
}

// Get an activity by ID
func GetActivityByID(id string) (*models.Activity, error) {
	var activity models.Activity //Devuelve un puntero a la estructura Actividad
	result := config.DB.First(&activity, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &activity, nil
}

// Add a new activity to the database
func AddActivity(activity models.Activity) error {
	if !possibleStates[activity.State] {
		return errors.New("Invalid state ")
	}
	return config.DB.Create(&activity).Error
}

// Update an activitie preload and save the changes in the DB
func UpdateActivity(id string, updatedActivity models.Activity) error {
	var activity models.Activity
	result := config.DB.First(&activity, id)
	if result.Error != nil {
		return result.Error
	}
	//Update the fields of the activity with the new values
	activity.Title = updatedActivity.Title
	activity.Description = updatedActivity.Description
	activity.Day = updatedActivity.Day
	activity.Duration = updatedActivity.Duration
	activity.State = updatedActivity.State
	activity.Instructor = updatedActivity.Instructor
	activity.Category = updatedActivity.Category
	activity.Cupo = updatedActivity.Cupo
	if !possibleStates[updatedActivity.State] {
		return errors.New("Invalid State. Must be 'active', 'cancelled' o 'finished'")
	}

	return config.DB.Save(&activity).Error
}

// Delete an activity by id
func DeleteActivity(id uint) error {
	result := config.DB.Delete(&models.Activity{}, id) //usa soft delete por defecto (gorm.model) el registro no se borra del todo sino se marca como eliminado
	return result.Error
}
