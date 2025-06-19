//controladores HTTP que reciben las peticiones del cliente, llaman a los servicios correspondientes y devuelven respuestas. Es el punto de entrada del backend a cada funcionalidad.

package handlers

import (
	"alua/config"
	"alua/models"
	"alua/services"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateInscription(c *gin.Context) {
	// Limpiar los parámetros
	activityParam := strings.TrimSpace(c.Param("ActivityID"))
	userParam := strings.TrimSpace(c.Param("UserID"))

	// Convertir a int
	activityID, err := strconv.Atoi(activityParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":  "Invalid Activity ID",
			"received": activityParam,
		})
		return
	}

	userID, err := strconv.Atoi(userParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":  "Invalid User ID",
			"received": userParam,
		})
		return
	}

	// Verificar que la actividad existe
	var activity models.Activity
	if err := config.DB.First(&activity, activityID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Activity not found",
			"id":      activityID,
		})
		return
	}

	// Verificar que el usuario existe
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
			"id":      userID,
		})
		return
	}

	// Crear inscripción
	inscription := models.Inscription{
		UserID:     uint(userID),
		ActivityID: uint(activityID),
	}

	if err := config.DB.Create(&inscription).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create inscription",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "Inscription created successfully",
		"inscription": inscription,
	})
}

func EditInscription(c *gin.Context) { //permite cambiar el estado sin eliminar la inscripcion
	role := c.GetHeader("Role") //verifica el rol del usuario
	if role != "socio" {
		c.JSON(http.StatusForbidden, gin.H{"message": "You do not have permission to perform this action"})
		return
	}

	// Obtener el ID de la inscripción desde la URL
	idStr := c.Param("id")
	idParsed, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Inscription ID invalid"})
		return
	}
	inscripcionID := uint(idParsed)

	// Obtener los nuevos datos del cuerpo de la solicitud
	var nueva models.Inscription
	if err := c.ShouldBindJSON(&nueva); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	// Validar que el estado sea uno permitido
	estadoValido := map[string]bool{
		"Esperando":  true,
		"Confirmado": true,
		"Cancelado":  true,
	}
	if !estadoValido[nueva.State] {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid state. Must be 'Waiting', 'Confirmed' o 'Cancelled'",
		})
		return
	}

	// Llamar al servicio para editar la inscripción
	if err := services.EditInscription(inscripcionID, nueva, 0); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inscripción actualizada correctamente"})
}

// maneja la eliminación de una inscripción
func DeleteInscription(c *gin.Context) {
	role := c.GetHeader("Role") //verifica el rol del usuario
	if role != "socio" {
		c.JSON(http.StatusForbidden, gin.H{"message": "You do not have permission to perform this action"})
		return
	}

	// Obtener ID de la inscripción
	idStr := c.Param("id")
	idParsed, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensaje": "ID inválido"})
		return
	}
	inscripcionID := uint(idParsed)

	if err := services.DeleteInscription(inscripcionID, 0); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"mensaje": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Inscripción eliminada exitosamente"})
}
