package handlers

import (
	"net/http"

	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	LastName     string    `json:"lastName"`
	DNI          uint      `json:"DNI"`
	Email        string    `json:"email"`
	Password     string    `json:"**********"`
	Rol          string    `json:"rol"`
	DateCreation time.Time `json:"fechaCreacion" gorm:"autoCreateTime"`
}

func createUser(ctx *gin.Context) {
	var newUser Usuario //Como vincular con el struct de la base de datos

	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	//Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Password could not be saved"})
		return
	}
	newUser.Password = string(hash)

	//Save user
	db.Create(&newUser)
	ctx.IndentedJSON(http.StatusCreated, newUser)
}

func getAllUsers(ctx *gin.Context) {
	var users []Usuario
	if err := db.Find(&users).Error; err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func updateUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	var modifyUser Usuario
	var user Usuario
	if err := ctx.BindJSON(&modifyUser); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Error reading request body"})
		return
	}
	if err := db.First(&user, id).Error; err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	user.Name = modifyUser.Name
	user.LastName = modifyUser.LastName
	user.DNI = modifyUser.DNI
	user.Email = modifyUser.Email
	user.Password = modifyUser.Password
	user.Rol = modifyUser.Rol
	db.Save(&user)
	ctx.IndentedJSON(http.StatusOK, user)
}

func deleteUserByID(ctx *gin.Context) {
	id := ctx.Param("id")

	result := db.Delete(&Usuario{}, "id = ?", id)
	if result.RowsAffected == 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "User deleted"})
}
