//controladores HTTP que reciben las peticiones del cliente, llaman a los servicios correspondientes y devuelven respuestas. Es el punto de entrada del backend a cada funcionalidad.

package handlers

import (
	"Proyecto2025-ArayaLuceroLussana/backend/models"
	"Proyecto2025-ArayaLuceroLussana/backend/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("clave") // Clave secreta para firmar el JWT

// Registro de usuario  (recibe un json con los datos del usuario, hashea la contraseña, llama a crearusuario para guardar en la bd )
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	// Hash of the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Password hashing failed"})
		return
	}

	user.Password = hashedPassword // Reemplaza la contraseña en el modelo con la versión hasheada

	if err := services.createUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error registering user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login de usuario (recibe datos, busca al usuario en BD, verifica la contra hasheada y si es correccto genera un JWT )
func Login(c *gin.Context) {
	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	user, err := services.getUserByEmail(data.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
		return
	}

	if !utils.CheckPasswordHash(data.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
		return
	}

	// Crea el JWT (token de autorizacion permite que mantenga su sesion sin guardar datos en el servidor)
	// El token contiene el ID del usuario, su email y rol, y una fecha de expiración
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserID":   user.ID,
		"Username": user.Email,
		"Rol":      user.Rol,
		"Exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString(jwtKey) // Firma el token con la clave secreta
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Token generation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Token": tokenString})
}
