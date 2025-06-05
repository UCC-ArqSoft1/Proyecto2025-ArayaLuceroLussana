// Contiene la logica del negocio
// logica de la app para funcionaar validaciones, reglas, calculos,etc
// Separa las reglas del negocio del acceso a datos y de la interfaz HTTP
package services

import (
	"Proyecto2025-ArayaLuceroLussana/backend/config"
	"Proyecto2025-ArayaLuceroLussana/backend/models"
	"errors"
)

// ValidateState checks if the provided state is valid.
var possibleStates = map[string]bool{
	"active":   true,
	"inactive": true,
	"finished": true,
}

// Get all the activities from the database
func showActivities() ([]models.Activity, error) {
	var activities []models.Activity
	result := config.DB.Find(&activities)
	return activities, result.Error
}

// Obtener una actividad por su ID
func getActivityById(id string) (*models.Activity, error) {
	var activity models.Activity //Devuelve un puntero a la estructura Actividad
	result := config.DB.First(&activity, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &activity, nil
}

// Add a new activity to the database
func addActivity(activity models.Activity) error {
	if !possibleStates[activity.State] {
		return errors.New("Invalid state ")
	}
	return config.DB.Create(&activity).Error
}

// Update an activitie preload and save the changes in the DB
func updateActivity(id string, updatedActividad models.Actividad) error {
	var activity models.Actividad
	result := config.DB.First(&activity, id)
	if result.Error != nil {
		return result.Error
	}
	//Update the fields of the activity with the new values
	activity.Name = updatedActividad.Name
	activity.Description = updatedActividad.Description
	activity.Day = updatedActividad.Day
	activity.Duration = updatedActividad.Duration
	activity.State = updatedActividad.State
	activity.Instructor = updatedActividad.Instructor
	activity.Category = updatedActividad.Category
	activity.Cupo = updatedActividad.Cupo
	if !possibleStates[updatedActividad.State] {
		return errors.New("Invalid State. Must be 'active', 'cancelled' o 'finished'")
	}

	return config.DB.Save(&activity).Error
}

// Delete an activity by id
func deleteActivity(id uint) error {
	result := config.DB.Delete(&models.Activity{}, id) //usa soft delete por defecto (gorm.model) el registro no se borra del todo sino se marca como eliminado
	return result.Error
}
