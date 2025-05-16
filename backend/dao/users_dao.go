package dao

type User struct {
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"unique"`
	PasswordHash string `gorm:"not null"` //Diferencia con el main, la password es HASH
}
