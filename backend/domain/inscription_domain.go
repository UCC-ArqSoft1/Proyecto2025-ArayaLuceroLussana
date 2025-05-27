package domain

// Representa el dominio del problema, con logica o interfaces de uso
// Puede contener interfaces que definen como se usan los datos

type Inscription struct {
	ID         int    `json:"id"`
	ActivityID int    `json:"activity_id"`
	UserID     int    `json:"user_id"`
	Status     string `json:"status"`     // e.g., "pending", "approved", "rejected"
	CreatedAt  string `json:"created_at"` // ISO 8601 format
	UpdatedAt  string `json:"updated_at"` // ISO 8601 format
}
