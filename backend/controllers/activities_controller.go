// Pregunta: Como vincular con la DB??
package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

type Activity struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Day          string    `json:"day"`
	Cupo         int       `json:"cupo"`
	Schedule     time.Time `json:"schedule"`
	Category     string    `json:"category"`
	Instructor   string    `json:"instructor"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Review       string    `json:"review"`
	CreationDate time.Time `json:"creationDate" gorm:"autoCreateTime"`
}

func deleteActivity(ctx *gin.Context) {
	id := ctx.Param("id")

	result := db.Delete(&Activity{}, "id = ?", id)
	if result.RowsAffected == 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Activity not found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Activity deleted"})
}

func updateActivity(ctx *gin.Context) {
	id := ctx.Param("id")
	var modifyActivity Activity //
	var activity Activity
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
	var newActivity Activity

	if err := ctx.BindJSON(&newActivity); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	db.Create(&newActivity)
	ctx.IndentedJSON(http.StatusCreated, newActivity)
}

func getActivityById(ctx *gin.Context) {
	id := ctx.Param("id")

	var actividad Activity //model.activity
	if err := db.First(&actividad, id).Error; err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}
	ctx.JSON(http.StatusOK, actividad)
}

func showActivities(ctx *gin.Context) {
	var activities []Activity
	if err := db.Find(&activities).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, activities)
}
