package main

import (
	"alua/config"
	"alua/handlers"
	"alua/middleware"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	r := gin.Default()

	// Configuración personalizada de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Rutas públicas
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.GET("/activities", handlers.ShowActivities)
	r.GET("/activities/:id", handlers.GetActivityByID)

	// Rutas de administrador
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.POST("/activity", handlers.AddActivity)
		admin.PUT("/activity/:id", handlers.UpdateActivity)
		admin.DELETE("/activity/:id", handlers.DeleteActivity)
	}

	// Rutas para socios
	socio := r.Group("/socio")
	socio.Use(middleware.AuthMiddleware())
	{
		socio.POST("/enroll/:UserID/:ActivityID", handlers.CreateInscription)
		socio.GET("/users/:id/activities", handlers.GetActivitiesByUser)
		socio.PUT("/inscription/:id", handlers.EditInscription)
		socio.DELETE("/inscription/:id", handlers.DeleteInscription)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
