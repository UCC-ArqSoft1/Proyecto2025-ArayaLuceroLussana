package services

import (
	"Proyecto2025-ArayaLuceroLussana/backend/config"
	"Proyecto2025-ArayaLuceroLussana/backend/models"
)

func createUser(user *models.User) error {
	return config.DB.Create(user).Error //Insert a new user in the DB
}

// GetUserByID retrieves a user by their ID from the database.
func getUserByEmail(Email string) (*models.User, error) {
	var user models.User
	if err := config.DB.Where("Email = ?", Email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// func (s *UsersService) Login(username string, password string) (int, string, error) {
// 	userDAO, err := s.usersClient.GetUserByUsername(username) //Llama al usersClient para crear el usuario por username
// 	if err != nil {
// 		return 0, "", fmt.Errorf("Error getting user by username: %w", err)
// 	}
// 	if utils.HashSHA256(password) != userDAO.PasswordHash { //Accede al password hash, lo compara con el HASH de la password que estoy recibiendo
// 		return 0, "", fmt.Errorf("Invalid password")
// 	}
// 	token, err := utils.GenerateJWT(userDAO.ID) //Genera el token en base al ID del usuario
// 	if err != nil {
// 		return 0, "", fmt.Errorf("Error generating token: %w", err)
// 	}
// 	return userDAO.ID, token, nil
// }
