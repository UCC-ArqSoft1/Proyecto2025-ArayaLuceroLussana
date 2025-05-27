package actividad

import (
	"ucc-gorm/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetActividadById(id int) model.Actividad {
	var actividad model.Actividad

	Db.Where("id = ?", id).Preload("Categoria").First(&actividad)
	log.Debug("Act: ", actividad)

	return actividad
}

func InsertActividad(actividad model.Actividad) model.Actividad {
	result := Db.Create(&actividad)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Actividad Created: ", actividad.Id)
	return actividad
}
