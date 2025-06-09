package handlers

import (
	"alua/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all activities for a user
func GetActivitiesByUser(c *gin.Context) {
	id := c.Param("id")
	activity, err := services.GetActivityByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting the activities for the user"})
		return
	}
	c.JSON(http.StatusOK, activity)
}
