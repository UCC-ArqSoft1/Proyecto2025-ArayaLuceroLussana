//Representa el dominio del problema, con logica o interfaces de uso
//Puede contener interfaces que definen como se usan los datos
package domain

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}
