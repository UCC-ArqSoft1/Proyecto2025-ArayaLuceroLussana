package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashPassword aplica SHA256 a la contraseña string recibida
func HashPassword(password string) (string, error) {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:]), nil
}

// Compara la contraseña en texto plano con el hash SHA256
func CheckPasswordHash(password, hash string) bool {
	hashedPassword, _ := HashPassword(password) //hashea de nuevo la contra en string, compara el nuevo hash con el almacenado en la bd
	return hashedPassword == hash               //validar credenciales en el login
}
