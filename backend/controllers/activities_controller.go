// Manejar las peticiones HTTP (GET,POST,etc)
// Son el primer punto de entrada del backend
// Funciones como GetUsers, CreateUser, y llama a los servicios correspondientes
// Encargan de la logica relacionada con las rutas HTTP y delegan el trabajo real a los servicios
// Pregunta: Como vincular con la DB??
package handlers

import (
	"net/http"

	"Proyecto2025-ArayaLuceroLussana/backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func deleteActivity(ctx *gin.Context) {
	id := ctx.Param("id")

	result := db.Delete(&models.Activity{}, "id = ?", id)
	if result.RowsAffected == 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Activity not found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Activity deleted"})
}

func updateActivity(ctx *gin.Context) {
	id := ctx.Param("id")
	var modifyActivity models.Activity //
	var activity models.Activity
	if err := ctx.BindJSON(&modifyActivity); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Error reading request body"})
		return
	}
	if err := db.First(&activity, id).Error; err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}
	activity.Day = modifyActivity.Day
	activity.Cupo = modifyActivity.Cupo
	activity.Schedule = modifyActivity.Schedule
	activity.Category = modifyActivity.Category
	activity.Instructor = modifyActivity.Instructor
	activity.Title = modifyActivity.Title
	activity.Description = modifyActivity.Description
	db.Save(&activity)
	ctx.IndentedJSON(http.StatusOK, activity)
}

func createActivity(ctx *gin.Context) {
	var newActivity models.Activity

	if err := ctx.BindJSON(&newActivity); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	db.Create(&newActivity)
	ctx.IndentedJSON(http.StatusCreated, newActivity)
}

func getActivityById(ctx *gin.Context) {
	id := ctx.Param("id")

	var actividad models.Activity //model.activity
	if err := db.First(&actividad, id).Error; err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}
	ctx.JSON(http.StatusOK, actividad)
}

func showActivities(ctx *gin.Context) {
	var activities []models.Activity
	if err := db.Find(&activities).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, activities)
}
