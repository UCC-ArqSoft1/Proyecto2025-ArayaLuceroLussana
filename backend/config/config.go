//Configuraciones del proyecto, conexiones, puertos, etc
//Leer variables de entorno, inicializar conexion a la DB
//No contiene logica de negocios, solo setup

package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("The file .env could not be charged, using system variables")
	}

	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	DB = db
	fmt.Println("Database connected successfully")
}
