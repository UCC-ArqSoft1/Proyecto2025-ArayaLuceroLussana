package handlers

import (
	"alua/config"
	"alua/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all activities for a user
func GetActivitiesByUser(c *gin.Context) {
	role := c.GetHeader("Role")
	if role != "socio" {
		c.JSON(http.StatusForbidden, gin.H{"message": "You do not have permission to perform this action"})
		return
	}

	userID := c.Param("id")

	var activities []models.Activity

	err := config.DB.
		Joins("JOIN inscriptions ON inscriptions.activity_id = activities.id").
		Where("inscriptions.user_id = ?", userID).
		Find(&activities).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting activities for user", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, activities)
}
