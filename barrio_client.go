package barrio

import (
	"ucc-gorm/model"

	"gorm.io/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetBarrioById(id int) model.Barrio {
	var barrio model.Barrio

	Db.Where("id = ?", id).First(&barrio)
	log.Debug("Barrio: ", barrio)

	return barrio
}

func GetBarrioSensoresById(id int) model.Barrio {
	var barrio model.Barrio

	Db.Where("id = ?", id).Preload("Sensores").First(&barrio)
	log.Debug("Barrio: ", barrio)

	return barrio
}

func InsertarBarrio(barrio model.Barrio) model.Barrio {
	result := Db.Create(&barrio)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Barrio Registrado: ", barrio)
	return barrio
}

func ActualizarBarrio(barrio model.Barrio) model.Barrio {
	
	result := Db.Model(&barrio).Update("descripcion", barrio.Descripcion)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Barrio Actualizado: ", barrio)
	return barrio
}

func GetBarrios() model.Barrios {
	var barrios model.Barrios
	Db.Order("descripcion").Find(&barrios)

	log.Debug("Barrios: ", barrios)

	return barrios
}
