//controladores HTTP que reciben las peticiones del cliente, llaman a los servicios correspondientes y devuelven respuestas. Es el punto de entrada del backend a cada funcionalidad.

package handlers

import (
	"alua/config"
	"alua/models"
	"alua/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get all the activities (public)
func ShowActivities(c *gin.Context) {
	activities, err := services.ShowActivities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting the activities"})
		return
	}
	c.JSON(http.StatusOK, activities)
}

// Get a specific activity by ID (public)
func GetActivityByID(c *gin.Context) {
	id := c.Param("id")
	activity, err := services.GetActivityByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Activity not found"})
		return
	}
	c.JSON(http.StatusOK, activity)
}

// Add a new activity (admin)
func AddActivity(c *gin.Context) {
	role := c.GetHeader("Role") //verifica el rol del usuario
	if role != "Admin" {
		c.JSON(http.StatusForbidden, gin.H{"message": "You do not have permission to perform this action"})
		return
	}

	var activity models.Activity
	if err := c.ShouldBindJSON(&activity); err != nil { //usa sbj para tranformar el JSON en una estructura actividad
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	if err := services.AddActivity(activity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating the activity"})
		return
	}

	c.JSON(http.StatusCreated, activity)
}

// Update an activity preload (admin)
func UpdateActivity(c *gin.Context) {
	role := c.GetHeader("Role") //verifica el rol del usuario
	if role != "Admin" {
		c.JSON(http.StatusForbidden, gin.H{"message": "You do not have permission to perform this action"})
		return
	}

	id := c.Param("id") //recibe el id por URL y los datos nuevos en el body de la petición (json)
	var data models.Activity
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}
	if err := services.UpdateActivity(id, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating the activity"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Activity updated successfully"})
}

// Delete an activity (admin)
func DeleteActivity(c *gin.Context) {
	role := c.GetHeader("Role") // Verifica el rol del usuario
	if role != "Admin" {
		c.JSON(http.StatusForbidden, gin.H{"message": "You do not have permission to perform this action"})
		return
	}

	// Obtener el ID de la actividad
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	activityID := uint(id)

	// Buscar inscripciones vinculadas a esta actividad
	var inscripciones []models.Inscription
	if err := config.DB.Where("activity_id = ?", activityID).Find(&inscripciones).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving inscriptions"})
		return
	}

	// Guardar los IDs de los usuarios inscriptos
	var userIDs []uint
	for _, ins := range inscripciones {
		userIDs = append(userIDs, ins.UserID)
	}

	// Eliminar inscripciones asociadas a esta actividad
	if err := config.DB.Where("activity_id = ?", activityID).Delete(&models.Inscription{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting inscriptions"})
		return
	}

	// Eliminar la actividad
	if err := services.DeleteActivity(activityID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting the activity"})
		return
	}

	// Notificar a los usuarios (simulado con logs por ahora)
	for _, userID := range userIDs {
		log.Printf("Actividad ID %d eliminada: notificar al usuario ID %d", activityID, userID)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activity deleted successfully and users notified"})
}
