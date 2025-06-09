// punto de entrada de la aplicación
package main

import (
	"alua/config"
	"alua/handlers"
	"alua/middleware"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()    // Inicializar la conexión a la base de datos (primero al iniciar programa)
	r := gin.Default() // Crear una nueva instancia del framework web gin

	//rutas publicas (disponibles sin autenticación)
	// Rutas públicas
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.GET("/activities", handlers.ShowActivities)
	r.GET("/activities/:id", handlers.GetActivityByID)

	// Rutas para administradores
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.POST("/activity", handlers.AddActivity)
		admin.PUT("/activity/:id", handlers.UpdateActivity)
		admin.DELETE("/activity/:id", handlers.DeleteActivity)
	}

	// Routes for authenticated users
	socio := r.Group("/socio")
	socio.Use(middleware.AuthMiddleware())
	{
		socio.POST("/enroll/:UserID/:ActivityID", handlers.CreateInscription)
		socio.GET("/users/:id/activities", handlers.GetActivitiesByUser)
		socio.PUT("/inscription/:id", handlers.EditInscription)
		socio.DELETE("/inscription/:id", handlers.DeleteInscription)
	}

	// Iniciar el servidor en el puerto 80
	if err := r.Run(":80"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
