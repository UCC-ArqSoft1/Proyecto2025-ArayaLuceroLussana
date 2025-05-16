// package main

// import (
// 	"log"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// var db *gorm.DB

// func main() {
// 	var err error
// 	//db, err := gorm.Open(sqlite.Open("gimnasio.db"), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Error conectando con base de datos")
// 	}

// 	db.AutoMigrate(&Usuario{}, &Activity{}, &Inscription{})
// 	// router.POST("/usuarios", crearUsuario)

// 	router := gin.Default()
// 	//if Usuario.Rol == "admin" {
// 	router.GET("/actividades", showActivities)
// 	router.GET("/actividades/:id", mostrarActividadById)
// 	router.GET("/usuarios", getAllUsers)
// 	router.PUT("/usuarios/:id", updateUserByID)
// 	router.DELETE("/usuarios/:id", deleteUser)
// 	router.POST("/actividades", createActivity)
// 	router.PUT("/actividades/:id", actualizarActividad)
// 	router.DELETE("/actividades/:id", deleteActivity)
// 	//} else if Usuario.Rol == "socio" {
// 	router.GET("/actividades", showActivities)
// 	router.GET("/actividades/:id", mostrarActividadById)
// 	router.POST("/inscripcion", crearInscripcion)
// }

// //}

package main

import (
	"ucc-gorm/app"
	"ucc-gorm/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
