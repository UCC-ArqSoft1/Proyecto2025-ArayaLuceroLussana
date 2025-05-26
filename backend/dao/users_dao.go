//Acceso a la base de datos
//Ejecutan consultas SQL, obtienen o escriben datos en la DB
//No deben tener logica de negocio ni saber como se usan los datos

package dao

type User struct {
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"unique"`
	PasswordHash string `gorm:"not null"` //Diferencia con el main, la password es HASH
}
