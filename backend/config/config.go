//Configuraciones del proyecto, conexiones, puertos, etc
//Leer variables de entorno, inicializar conexion a la DB
//No contiene logica de negocios, solo setup

package config

import (
	"fmt"
	"log"
	"os"

	"alua/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("The file .env could not be charged, using system variables")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	DB = db
	fmt.Println("Database connected successfully")

	//Migracion de los modelos a la base de datos
	err = db.AutoMigrate(&models.User{}, &models.Activity{}, &models.Inscription{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate models: %v", err)
	}
}
