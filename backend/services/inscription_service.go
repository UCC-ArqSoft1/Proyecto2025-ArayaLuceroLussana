package services

import (
	"alua/config"
	"alua/models"
	"errors"
	"time"
)

// Create a new inscription for an activity
func CreateInscription(UserID uint, ActivityID uint) error {

	// Verify if the user is already registered for the activity
	var inscription models.Inscription
	err := config.DB.Where("UserID = ? AND ActivityID = ?", UserID, ActivityID).First(&inscription).Error
	if err == nil {
		return errors.New("user already registered for this activity")
	}

	// verify if the activity already exists and there's a spot available
	var activity models.Activity
	if err := config.DB.First(&activity, ActivityID).Error; err != nil {
		return errors.New("activity not found")
	}
	// Check if the activity is active
	if activity.State != "Activo" {
		return errors.New("inscription couldn't be done: activity is not active")
	}

	//Count the number of registered users for the activity
	var totalInscription int64
	config.DB.Model(&models.Inscription{}).Where("activity_id = ?", ActivityID).Count(&totalInscription)
	if totalInscription >= int64(activity.Cupo) {
		return errors.New("there are no spots available for this activity")
	}
	// Create a new inscription
	new := models.Inscription{
		UserID:     UserID,
		ActivityID: ActivityID,
		Date:       time.Now().Format("2006-01-02"),
		State:      "Confirmed", // Default state
	}

	return config.DB.Create(&new).Error
}

// Change the state of an inscription
func EditInscription(id uint, new models.Inscription, UserID uint) error {
	var inscription models.Inscription

	//Search for the inscription by id
	if err := config.DB.First(&inscription, id).Error; err != nil {
		return errors.New("inscription not found")
	}

	//Verify that the inscription belongs to the user
	if inscription.UserID != UserID {
		return errors.New("no permission to edit this inscription")
	}
	//Only change the state of the inscription
	inscription.State = new.State

	return config.DB.Save(&inscription).Error
}

// Delete an inscription
func DeleteInscription(id uint, UserID uint) error {
	var inscription models.Inscription

	//Search for the inscription by id
	if err := config.DB.First(&inscription, id).Error; err != nil {
		return errors.New("inscription not found")
	}

	//Verify that the inscription belongs to the user
	if inscription.UserID != UserID {
		return errors.New("no permission to delete this inscription")
	}

	return config.DB.Delete(&inscription).Error
}
