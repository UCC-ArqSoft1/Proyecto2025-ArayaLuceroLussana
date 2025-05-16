package clients

import (
	"backend/dao"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLClient struct {
	DB *gorm.DB //Conexion propia con la base de datos
}

func NewMySQLClient() *MySQLClient {
	dsnFormat := "%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4&loc=Local" //Genera el string de la conexion
	dsn := fmt.Sprintf(dsnFormat, "root", "127.0.0.1", 3306, "backend")
	//Abre la conexion
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("Error connecting to database: %w", err))
	}
	//Llama la funcion autoMigrate para crear las tablas
	for _, table := range []interface{}{
		dao.User{},
		dao.Activity{},
		dao.Inscription{},
		dao.TimeSlot{},
	} {
		if err := db.AutoMigrate(&table); err != nil {
			panic(fmt.Errorf("Error migrating table %w", err))
		}
	}

	return &MySQLClient{
		DB: db,
	}

}

func (c *MySQLClient) GetUserByUsername(username string) (dao.User, error) { //Uso en el service para poder traer el usuario en base al username
	var userDAO dao.User
	//Es equivalente a SELECT * FROM users WHERE username = "admin" LIMIT 1;
	txn := c.DB.First(&userDAO, "username = ?", username)
	if txn.Error != nil {
		return dao.User{}, fmt.Errorf("Error getting user by username: %w", txn.Error)
	}
	return userDAO, nil
}
