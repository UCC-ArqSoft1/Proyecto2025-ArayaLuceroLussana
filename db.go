package db

import (
	barrioClient "ucc-gorm/clients/barrio"
	sensorClient "ucc-gorm/clients/sensor"
	medicionClient "ucc-gorm/clients/medicion"
	"ucc-gorm/model"
	"gorm.io/gorm/logger"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	log "github.com/sirupsen/logrus"
)
var (
	dba  *gorm.DB
	err error
)

func init() {
	
	// Using a temporary SqlLite database in system memory
	dba, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	//dba, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
	//	Logger: logger.Default.LogMode(logger.Info),
	//})
	

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}
	// Adding all CLients that we build.
	barrioClient.Db = dba
	sensorClient.Db = dba
	medicionClient.Db = dba
}


func StartDbEngine() {
	// Migrating all classes model.
	dba.AutoMigrate(&model.Barrio{})
	dba.AutoMigrate(&model.Sensor{})
	dba.AutoMigrate(&model.Medicion{})
	log.Info("Finishing Migration Database Tables")
	testGorm()
}

// For Class
func testGorm() {
	newBarrio := model.Barrio{
        Descripcion: "Nueva Córdoba",
    }
	barrioClient.InsertarBarrio(newBarrio)

	newBarrio = barrioClient.GetBarrioById(1)

	newBarrio.Descripcion = "Los Naranjos"
	

	barrioClient.ActualizarBarrio(newBarrio)

	barrioClient.GetBarrioById(1)

}