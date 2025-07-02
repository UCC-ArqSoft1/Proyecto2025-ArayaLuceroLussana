// Contiene la logica del negocio
// logica de la app para funcionaar validaciones, reglas, calculos,etc
// Separa las reglas del negocio del acceso a datos y de la interfaz HTTP
package services

import (
	"alua/config"
	"alua/models"
	"errors"
)

// Verifica que el estado sea valido
var possibleStates = map[string]bool{
	"Activo":     true,
	"Inactivo":   true,
	"Finalizado": true,
}

// Get las actividades de la base de datos
func ShowActivities() ([]models.Activity, error) {
	var activities []models.Activity
	result := config.DB.Find(&activities)
	return activities, result.Error
}

// Get actividad con el ID
func GetActivityByID(id string) (*models.Activity, error) {
	var activity models.Activity //Devuelve un puntero a la estructura Actividad
	result := config.DB.First(&activity, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &activity, nil
}

// Agrega una nueva actidad en la BD
func AddActivity(activity models.Activity) error {
	if !possibleStates[activity.State] {
		return errors.New("Invalid state ")
	}
	return config.DB.Create(&activity).Error
}

// Actualiza las actividades precargadas y guarda los cambios en la BD
func UpdateActivity(id string, updatedActivity models.Activity) error {
	var activity models.Activity
	result := config.DB.First(&activity, id)
	if result.Error != nil {
		return result.Error
	}

	// Actualiza los campos de las atividades con los nuevos valores
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

// Delete una actividad con el ID
func DeleteActivity(id uint) error {
	result := config.DB.Delete(&models.Activity{}, id) //usa soft delete por defecto (gorm.model) el registro no se borra del todo sino se marca como eliminado
	return result.Error
}
