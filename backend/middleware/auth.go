package middleware

import (
	"net/http" // para las respuestas HTTP
	"strings"  // para manipular cadenas de texto

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt" //libreria para manejar JWT
)

var jwtKey = []byte("clave") // Clave secreta para firmar y verificar los tokens (no se comparte). se convierte en bte para la libreria jwt

// middleware para autenticar usuarios verificando que haya un token JWT válido en la cabecera HTTP
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization") //extrae el token del encabezado http "Authorization"

		//si no hay token, devuelve error 401 (no autorizado) y frena la ejecucion
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"mensaje": "Token no proporcionado"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(tokenString, "Bearer ") //elimina el prefijo "Token " para obtener solo el jwt real (ej: "token abc123" -> "abc123")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil //intenta parsear y verificar el token usando la clave secreta
		})

		// manejar error de parseo del token (devuelve 401)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"mensaje": "Token inválido"})
			c.Abort()
			return
		}

		//si el token es válido, se accede a los claims (datos del usuario dentro del token)
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("usuarioID", claims["usuarioID"]) // Guardar el ID del usuario en el contexto c para que pueda usarse en los handlers (ej usuarioID := c.Get("usuarioID") rol := c.Get("rol"))
			c.Set("rol", claims["rol"])             // Guardar el rol del usuario
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"mensaje": "Token inválido"})
			c.Abort()
			return
		}

	}
}

// permitir solo usuarios con rol "admin"
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rol, exists := c.Get("rol")
		rolStr, ok := rol.(string) // asegura que sea string

		if !exists || !ok || (rolStr != "admin" && rolStr != "Admin") {
			c.JSON(http.StatusForbidden, gin.H{"error": "Acceso denegado"})
			c.Abort()
			return
		}

		c.Next()
	}
}
