package services

import (
	"Proyecto2025-ArayaLuceroLussana/backend/config"
	"Proyecto2025-ArayaLuceroLussana/backend/models"
	"errors"
	"time"
)

// Create a new inscription for an activity
func createInscription(UserID uint, ActivityID uint) error {

	// Verify if the user is already registered for the activity
	err := config.DB.Where("UserID = ? AND ActivityID = ?", UserID, ActivityID).First(&inscripcion).Error
	if err == nil {
		return errors.New("User already registered for this activity")
	}

	// verify if the activity already exists and there's a spot available
	var activity models.Actividad
	if err := config.DB.First(&activity, ActivityID).Error; err != nil {
		return errors.New("Activity not found")
	}

	if activity.State != "active" {
		return errors.New("Inscription couldn't be done: activity is not active")
	}

	//Count the number of registered users for the activity
	var totalInscription int64
	config.DB.Model(&models.Inscription{}).Where("ActivityID = ?", ActivityID).Count(&totalInscription)
	if totalInscription >= int64(activity.Cupo) {
		return errors.New("There are no spots available for this activity")
	}
	// Create a new inscription
	new := models.Inscription{
		UserID:     UserID,
		ActivityID: ActivityID,
		Date:       time.Now().Format("2006-01-02"),
		Estado:     "Confirmed", // Default state
	}

	return config.DB.Create(&new).Error
}

// Change the state of an inscription
func editInscriptionn(id uint, new models.Inscription, UserID uint) error {
	var inscription models.Inscription

	//Search for the inscription by id
	if err := config.DB.First(&inscription, id).Error; err != nil {
		return errors.New("Inscription not found")
	}

	//Verify that the inscription belongs to the user
	if inscription.UserID != UserID {
		return errors.New("No permission to edit this inscription")
	}
	//Only change the state of the inscription
	inscription.State = new.State

	return config.DB.Save(&inscription).Error
}

// Delete an inscription
func deleteInscripction(id uint, UserID uint) error {
	var inscription models.Inscription

	//Search for the inscription by id
	if err := config.DB.First(&inscription, id).Error; err != nil {
		return errors.New("Inscription not found")
	}

	//Verify that the inscription belongs to the user
	if inscription.UserID != UserID {
		return errors.New("No permission to delete this inscription")
	}

	return config.DB.Delete(&inscription).Error
}
