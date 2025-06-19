//controladores HTTP que reciben las peticiones del cliente, llaman a los servicios correspondientes y devuelven respuestas. Es el punto de entrada del backend a cada funcionalidad.

package handlers

import (
	"alua/models"
	"alua/services"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateInscription(c *gin.Context) {
	activityParam := strings.TrimSpace(c.Param("ActivityID"))
	userParam := strings.TrimSpace(c.Param("UserID"))

	activityID, err := strconv.Atoi(activityParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Activity ID", "received": activityParam})
		return
	}

	userID, err := strconv.Atoi(userParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid User ID", "received": userParam})
		return
	}

	err = services.CreateInscription(uint(userID), uint(activityID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Inscription created successfully",
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
	role := c.GetHeader("Role") // Verifica rol
	if role != "socio" {
		c.JSON(http.StatusForbidden, gin.H{"message": "You do not have permission to perform this action"})
		return
	}

	// Obtener ID de la inscripción
	idStr := c.Param("id")
	idParsed, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}
	inscripcionID := uint(idParsed)

	// Obtener UserID del parámetro URL
	userIDStr := c.Param("UserID")
	if userIDStr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "UserID no proporcionado"})
		return
	}
	userIDParsed, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "UserID inválido"})
		return
	}
	userID := uint(userIDParsed)

	// Llamar al servicio para eliminar inscripción con validación
	if err := services.DeleteInscription(inscripcionID, userID); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inscripción eliminada exitosamente"})
}
