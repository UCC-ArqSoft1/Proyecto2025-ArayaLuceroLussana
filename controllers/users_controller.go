package controllers

import (
	"backend/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	usersService UsersService
}

type UsersService interface {
	Login(username string, password string) (int, string, error)
}

func NewUserController(usersService UsersService) *UserController {
	return &UserController{
		usersService: usersService,
	}
}

func (c *UserController) Login(ctx *gin.Context) {
	//Recibe el usuario y la password desde el body de la request
	var request domain.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", err.Error()})

		return
	}
	//Llama al service para loguear al usuario
	userID, token, err := c.usersService.Login(request.Username, request.Password)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Invalid credentials", err.Error()})
		return
	}
	//Devuelve el ID y el token al cliente, en al request
	ctx.JSON(http.StatusOK, domain.loginResponse{
		UserID: userID,
		Token:  token,
	})
}
