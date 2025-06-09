//controladores HTTP que reciben las peticiones del cliente, llaman a los servicios correspondientes y devuelven respuestas. Es el punto de entrada del backend a cada funcionalidad.

package handlers

import (
	"alua/models"
	"alua/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateInscription(c *gin.Context) {
	// Obtener ID del usuario desde el token, verifica que el usuarioid de la URL coincida con el id del token
	tokenUserIDRaw, exists := c.Get("UserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}
	tokenUserID := uint(tokenUserIDRaw.(float64)) // JWT devuelve float64

	//Get userID and activityID from URL parameters
	userIDStr := c.Param("UserID")
	activityIDStr := c.Param("ActivityID")

	userIDParsed, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	activityIDParsed, err := strconv.ParseUint(activityIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	userID := uint(userIDParsed)
	activityID := uint(activityIDParsed)

	// Verify that the userID from the token matches the userID from the URL
	if tokenUserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"message": "You do not have permission to perform this action"})
		return
	}

	// Call the service to create the inscription
	err = services.CreateInscription(userID, activityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inscription created successfully"})
}

func EditInscription(c *gin.Context) { //permite cambiar el estado sin eliminar la inscripcion
	// Obtener el ID del usuario desde el token
	tokenUserIDRaw, exists := c.Get("UserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}
	tokenUserID := uint(tokenUserIDRaw.(float64))

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
		"Waiting":   true,
		"Confirmed": true,
		"Cancelled": true,
	}
	if !estadoValido[nueva.State] {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid state. Must be 'Waiting', 'Confirmed' o 'Cancelled'",
		})
		return
	}

	// Llamar al servicio para editar la inscripción
	if err := services.EditInscription(inscripcionID, nueva, tokenUserID); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inscripción actualizada correctamente"})
}

// maneja la eliminación de una inscripción
func DeleteInscription(c *gin.Context) {
	tokenUserIDRaw, exists := c.Get("usuarioID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"mensaje": "Token inválido"})
		return
	}
	tokenUserID := uint(tokenUserIDRaw.(float64))

	// Obtener ID de la inscripción
	idStr := c.Param("id")
	idParsed, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensaje": "ID inválido"})
		return
	}
	inscripcionID := uint(idParsed)

	if err := services.DeleteInscription(inscripcionID, tokenUserID); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"mensaje": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Inscripción eliminada exitosamente"})
}
