package main

import (
	"alua/config"
	"alua/handlers"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	r := gin.Default()

	// Configuración de CORS para permitir el frontend (React por ejemplo)
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

	// Rutas de administrador (requiere autenticación y rol admin)
	// admin := r.Group("/admin")
	// admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	// {
	r.POST("/admin/activity", handlers.AddActivity)
	r.PUT("/admin/activity/:id", handlers.UpdateActivity)
	r.DELETE("admin/activity/:id", handlers.DeleteActivity)
	//}

	// Rutas para socios autenticados
	// socio := r.Group("/socio")
	// socio.Use(middleware.AuthMiddleware())
	// {
	r.POST("socio/enroll/:UserID/:ActivityID", handlers.CreateInscription)
	r.GET("socio/users/:id/activities", handlers.GetActivitiesByUser)
	r.PUT("socio/inscription/:id", handlers.EditInscription)
	r.DELETE("socio/inscription/:id", handlers.DeleteInscription)
	//}

	// Arranca el servidor en el puerto 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
