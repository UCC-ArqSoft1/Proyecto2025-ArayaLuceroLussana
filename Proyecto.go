package main

import (
	"net/http"
	"time"
)

type Usuario struct {
	ID            string    `json:"id" gorm:"primarykey;autoIncrement"`
	Nombre        string    `json:"nombre"`
	Apellido      string    `json:"apellido"`
	DNI           uint      `json:"DNI"`
	Email         string    `json:"email"`
	Contraseña    string    `json:"contraseña"`
	Rol           string    `json:"rol"`
	FechaCreacion time.Time `json:"fechaCreacion"`
}

type Acividades struct {
	ID            string    `json:"id" gorm:"primarykey;autoIncrement"`
	Dia           string    `json:"dia"`
	Cupo          int       `json:"cupo"`
	Horario       time.Time `json:"horario"`
	Categoria     string    `json:"categoria"`
	Instructor    string    `json:"instructor"`
	Titulo        string    `json:"titulo"`
	Descripcion   string    `json:"descripcion"`
	Resena        string    `json:"resena"`
	FechaCreacion time.Time `json:"fechaCreacion"`
}

type Inscripcion struct {
	ID               string    `json:"id" gorm:"primarykey;autoIncrement"`
	UsuarioID        string    `json:"usuarioID"`
	ActividadID      string    `json:"actividadID"`
	FechaInscripcion time.Time `json:"fechaInscripcion"`
	IsActivo         bool      `json:"estado"`
}

func mostrarActividades(ctx *gin.Context) {
	var actividades []Acividades
	if err := db.Find(&actividades).Error; err != nil {
		ctx.IndentendJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, actividades)
}

func mostrarActividById(ctx *gin.Context) {
	id := ctx.Param("id")

	var actividad Acividades
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
	var modifyUsuarioo Usuario
	var usuario Usuario
	if err := ctx.BindJSON(&modifyUsuarioo); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Error al leer el cuerpo de la solicitud"})
		return
	}
	if err := db.First(&usuario, id).Error; err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}
	usuario.Nombre = modifyUsuarioo.Nombre
	usuario.Apellido = modifyUsuarioo.Apellido
	usuario.DNI = modifyUsuarioo.DNI
	usuario.Email = modifyUsuarioo.Email
	usuario.Contraseña = modifyUsuarioo.Contraseña
	usuario.Rol = modifyUsuarioo.Rol
	db.Save(&usuario)
	ctx.IntendedJSON(http.StatusOK, usuario)
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
	var newActividad Acividades

	if err := ctx.BindJSON(&actividad); err != nil {
		return
	}

	db.Create(&newActividad)
	ctx.IndentedJSON(http.StatusCreated, newActividad)
}

func actualizarActividad(ctx *gin.Context) {
	id := ctx.Param("id")
	var modifyActividad Acividades
	var actividad Acividades
	if err := ctx.BindJSON(&modifyActividad); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Error al leer el cuerpo de la solicitud"})
		return
	}
	if err := db.First(&actividad, id).Error; err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Actividad no encontrada"})
		return
	}
	actividad.Dia = modifyActividad.Dia
	actividad.Cupo = modifyActividad.Cupo
	actividad.Horario = modifyActividad.Horario
	actividad.Categoria = modifyActividad.Categoria
	actividad.Instructor = modifyActividad.Instructor
	actividad.Titulo = modifyActividad.Titulo
	actividad.Descripcion = modifyActividad.Descripcion
	db.Save(&actividad)
	ctx.IntendedJSON(http.StatusOK, actividad)
}

func main() {
	var err error
	// db, err = gorm.Open("sqlite3", "test.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// router.POST("/usuarios", crearUsuario)

	router := gin.Default()
	if Usuario.Rol == "admin" {
		router.GET("/actividades", mostrarActividades)
		router.GET("/actividades/:id", mostrarActividById)
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
