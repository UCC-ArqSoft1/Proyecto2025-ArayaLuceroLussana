package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Usuario struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	LastName      string    `json:"lastName"`
	DNI           uint      `json:"DNI"`
	Email         string    `json:"email"`
	Contraseña    string    `json:"contraseña"`
	Rol           string    `json:"rol"`
	FechaCreacion time.Time `json:"fechaCreacion"`
}

type Actividades struct {
	ID           string    `json:"id"`
	Day          string    `json:"day"`
	Cupo         int       `json:"cupo"`
	Horario      time.Time `json:"horario"`
	Category     string    `json:"category"`
	Instructor   string    `json:"instructor"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Review       string    `json:"review"`
	CreationDate string    `json:"creationDate"`
}

type Inscription struct {
	ID         string `json:"id"`
	UserID     string `json:"user_ID"`
	ActivityID string `json:"activity_ID"`
	Date       string `json:"date"` // Fecha especifica de la sesion
	//IsActivo         bool      `json:"estado"`
	StartTime string `json:"start_time"` //Hora de inicio  Ej: "18:00"
	EndTime   string `json:"end_time"`   //Hora de finalizacion  Ej: "19:00"

}

func mostrarActividades(ctx *gin.Context) {
	var actividades []Actividades
	if err := db.Find(&actividades).Error; err != nil {
		ctx.IndentendJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, actividades)
}

func mostrarActividadById(ctx *gin.Context) {
	id := ctx.Param("id")

	var actividad Actividades
	if err := db.First(&actividad, id).Error; err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Actividad no encontrada"})
		return
	}
	ctx.JSON(http.StatusOK, actividad)
}

func mostrarUsuarios(ctx *gin.Context) {
	var usuarios []Usuario
	if err := db.Find(&usuarios).Error; err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, usuarios)
}

func actualizarUsuarioById(ctx *gin.Context) {
	id := ctx.Param("id")
	var modifyUsuario Usuario
	var usuario Usuario
	if err := ctx.BindJSON(&modifyUsuario); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Error al leer el cuerpo de la solicitud"})
		return
	}
	if err := db.First(&usuario, id).Error; err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}
	usuario.Name = modifyUsuario.Name
	usuario.LastName = modifyUsuario.LastName
	usuario.DNI = modifyUsuario.DNI
	usuario.Email = modifyUsuario.Email
	usuario.Contraseña = modifyUsuario.Contraseña
	usuario.Rol = modifyUsuario.Rol
	db.Save(&usuario)
	ctx.IndentedJSON(http.StatusOK, usuario)
}

func borrarUsuario(ctx *gin.Context) {
	id := ctx.Param("id")

	result := db.Delete(&Usuario{}, "id = ?", id)
	if result.RowsAffected == 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Usuario no encontrado"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Usuario eliminado"})
}

func crearActividad(ctx *gin.Context) {
	var newActividad Actividades

	if err := ctx.BindJSON(&actividad); err != nil {
		return
	}

	db.Create(&newActividad)
	ctx.IndentedJSON(http.StatusCreated, newActividad)
}

func actualizarActividad(ctx *gin.Context) {
	id := ctx.Param("id")
	var modifyActividad Actividades
	var actividad Actividades
	if err := ctx.BindJSON(&modifyActividad); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Error al leer el cuerpo de la solicitud"})
		return
	}
	if err := db.First(&actividad, id).Error; err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Actividad no encontrada"})
		return
	}
	actividad.Day = modifyActividad.Day
	actividad.Cupo = modifyActividad.Cupo
	actividad.Horario = modifyActividad.Horario
	actividad.Category = modifyActividad.Category
	actividad.Instructor = modifyActividad.Instructor
	actividad.Title = modifyActividad.Title
	actividad.Description = modifyActividad.Description
	db.Save(&actividad)
	ctx.IndentedJSON(http.StatusOK, actividad)
}

func eliminarActividad(ctx *gin.Context) {
	id := ctx.Param("id")

	result := db.Delete(&Actividades{}, "id = ?", id)
	if result.RowsAffected == 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Actividad no encontrada"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Actividad eliminada"})
}

func crearInscripcion(ctx *gin.Context) {
	var newInscripcion Inscription

	if err := ctx.BindJSON(&newInscripcion); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Error al leer el cuerpo de la solicitud"})
		return
	}

	db.Create(&newInscripcion)
	ctx.IndentedJSON(http.StatusCreated, newInscripcion)
}

var db *gorm.DB

func main() {
	// db, err := gorm.Open("sqlite3", "test.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// router.POST("/usuarios", crearUsuario)

	router := gin.Default()
	if Usuario.Rol == "admin" {
		router.GET("/actividades", mostrarActividades)
		router.GET("/actividades/:id", mostrarActividadById)
		router.GET("/usuarios", mostrarUsuarios)
		router.PUT("/usuarios/:id", actualizarUsuarioById)
		router.DELETE("/usuarios/:id", borrarUsuario)
		router.POST("/actividades", crearActividad)
		router.PUT("/actividades/:id", actualizarActividad)
		router.DELETE("/actividades/:id", eliminarActividad)
	} else if Usuario.Rol == "socio" {
		router.GET("/actividades", mostrarActividades)
		router.GET("/actividades/:id", mostrarActividadById)
		router.POST("/inscripcion", crearInscripcion)
	}
}
