// punto de entrada de la aplicación
package main

import (
	"Proyecto2025-ArayaLuceroLussana/backend/config"
	"Proyecto2025-ArayaLuceroLussana/backend/handlers"
	"Proyecto2025-ArayaLuceroLussana/backend/middlewares"
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
	r.GET("/activities", handlers.GetActividades)
	r.GET("/activities/:id", handlers.getActivityByID)

	// Rutas para administradores
	admin := r.Group("/admin")
	admin.Use(middlewares.AuthMiddleware(), middlewares.AdminMiddleware())
	{
		admin.POST("/activity", handlers.createActivity)
		admin.PUT("/activity/:id", handlers.updateActivity)
		admin.DELETE("/activity/:id", handlers.deleteActivity)
	}

	// Rutas para socios autenticados
	socio := r.Group("/socio")
	socio.Use(middlewares.AuthMiddleware())
	{
		socio.POST("/enroll/:UserID/:ActivityID", handlers.createInscription)
		socio.GET("/users/:id/activities", handlers.getActivitiesByUser)
		socio.PUT("/inscription/:id", handlers.editInscription)
		socio.DELETE("/inscription/:id", handlers.deleteInscription)
	}

	// Iniciar el servidor en el puerto 80
	if err := r.Run(":80"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
