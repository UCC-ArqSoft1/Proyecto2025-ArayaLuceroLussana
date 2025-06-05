//controladores HTTP que reciben las peticiones del cliente, llaman a los servicios correspondientes y devuelven respuestas. Es el punto de entrada del backend a cada funcionalidad.

package handlers

import (
	"Proyecto2025-ArayaLuceroLussana/backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get all the activities
func showActivities(c *gin.Context) {
	activities, err := services.showActivities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting the activities"})
		return
	}
	c.JSON(http.StatusOK, activities)
}

// tener actividad especifica por ID
func getActivityByID(c *gin.Context) {
	id := c.Param("id")
	activity, err := services.getActivityById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Activity not found"})
		return
	}
	c.JSON(http.StatusOK, activity)
}

// Add a new activity (admin)
func addActivity(c *gin.Context) {
	var activity models.Actividad
	if err := c.ShouldBindJSON(&activity); err != nil { //usa sbj para tranformar el JSON en una estructura actividad
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	if err := services.addActivity(activity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating the activity"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Activity created successfully"})
}

// Update an activity preload (admin)
func updateActivity(c *gin.Context) {
	id := c.Param("id") //recibe el id por URL y los datos nuevos en el body de la petición (json)
	var data models.Activity
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}
	if err := services.updateActivity(id, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating the activity"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Activity updated successfully"})
}

// Delete an activity (admin)
func deleteActivity(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64) //convierte el ID recibido como string a tipo uint
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID "})
		return
	}
	if err := services.deleteActivity(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting the activity"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Activity deleted successfully"})
}
