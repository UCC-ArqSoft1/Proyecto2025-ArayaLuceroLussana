//controladores HTTP que reciben las peticiones del cliente, llaman a los servicios correspondientes y devuelven respuestas. Es el punto de entrada del backend a cada funcionalidad.

package handlers

import (
	"alua/models"
	"alua/services"
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
	role := c.GetHeader("Role") //verifica el rol del usuario
	if role != "Admin" {
		c.JSON(http.StatusForbidden, gin.H{"message": "You do not have permission to perform this action"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64) //convierte el ID recibido como string a tipo uint
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID "})
		return
	}
	if err := services.DeleteActivity(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting the activity"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Activity deleted successfully"})
}
