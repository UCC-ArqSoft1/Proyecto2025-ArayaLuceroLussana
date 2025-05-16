package services

import (
	"backend/dao"
	"backend/utils"
	"fmt"
)

type UsersClient interface {
	GetUserByUsername(username string) (dao.User, error)
}

type UsersService struct {
	usersClient UsersClient
}

func NewUsersService(usersClient UsersClient) *UsersService {
	return &UsersService{
		usersClient: usersClient,
	}
}

func (s *UsersService) Login(username string, password string) (int, string, error) {
	userDAO, err := s.usersClient.GetUserByUsername(username) //Llama al usersClient para crear el usuario por username
	if err != nil {
		return 0, "", fmt.Errorf("Error getting user by username: %w", err)
	}
	if utils.HashSHA256(password) != userDAO.PasswordHash { //Accede al password hash, lo compara con el HASH de la password que estoy recibiendo
		return 0, "", fmt.Errorf("Invalid password")
	}
	token, err := utils.GenerateJWT(userDAO.ID) //Genera el token en base al ID del usuario
	if err != nil {
		return 0, "", fmt.Errorf("Error generating token: %w", err)
	}
	return userDAO.ID, token, nil
}
