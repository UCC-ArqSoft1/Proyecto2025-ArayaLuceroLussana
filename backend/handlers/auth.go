package handlers

import (
	"alua/models"
	"alua/services"
	"alua/utils/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("clave")

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Datos inválidos"})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al hashear la contraseña"})
		return
	}
	user.Password = hashedPassword

	if err := services.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al registrar el usuario"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuario registrado correctamente"})
}

func Login(c *gin.Context) {
	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Datos inválidos"})
		return
	}

	user, err := services.GetUserByEmail(data.Email)
	if err != nil || !utils.CheckPasswordHash(data.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Usuario o contraseña incorrectos"})
		return
	}

	// Forma más simple de crear token
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"rol":   user.Rol,
		"exp":   time.Now().Add(72 * time.Hour).Unix(),
	}

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al generar el token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString,
		"userId": user.ID,
	})
}
