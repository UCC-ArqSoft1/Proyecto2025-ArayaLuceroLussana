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
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"mensaje": "Token no proporcionado"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(tokenString, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"mensaje": "Token inválido"})
			c.Abort()
			return
		}

		var claims jwt.MapClaims
		if parsedClaims, ok := token.Claims.(jwt.MapClaims); ok {
			claims = parsedClaims
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"mensaje": "Token inválido"})
			c.Abort()
			return
		}

		rol, ok := claims["rol"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"mensaje": "Rol no válido en el token"})
			c.Abort()
			return
		}

		c.Set("usuarioID", claims["usuarioID"]) // si lo usás
		c.Set("rol", rol)

		c.Next()
	}
}

// permitir solo usuarios con rol "admin"
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rolValue, exists := c.Get("rol")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Acceso denegado: rol no encontrado"})
			c.Abort()
			return
		}

		rolStr, ok := rolValue.(string)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Acceso denegado: rol inválido"})
			c.Abort()
			return
		}

		if strings.ToLower(rolStr) != "Admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Acceso denegado: rol no autorizado"})
			c.Abort()
			return
		}

		c.Next()
	}
}
