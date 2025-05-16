package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	jwtDuration = 24 * time.Hour
	jwtSecret   = "LuzPaulaMariano" // Cambia esto por una clave secreta más segura
)

func HashSHA256(value string) string { //Provee variable string y devuelve variable string
	hash := sha256.Sum256([]byte(value))
	return hex.EncodeToString(hash[:])
}

func GenerateJWT(userID int) (string, error) { //Recibe el usuario y devuelve un string
	//Setear expiracion del token
	expirationTime := time.Now().Add(jwtDuration)

	//Construir los claims
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime), //Datos que viajan en el token
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "backennd",
		Subject:   "auth", //Para que el token es valido
		ID:        fmt.Sprintf("%d", userID),
	}

	//Crear el token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Firmar el token
	tokenString, err := token.SignedString([]byte(jwtSecret)) //Encriptar el token en base a una contraseña
	if err != nil {
		return "", fmt.Errorf("Error generating token: %w", err)
	}
	return tokenString, nil

}
