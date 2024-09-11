package models

type Users struct {
	ID        int    `json:"id"`
	LastName  string `json:"lastname"`
	FirstName string `json:"firstname"`
	Adresss   string `json:"address"`
	City      string `json:"city"`
	Email     string `json:"email"`
	Password  string `json:"pass"`
	Role      string `json:"role"`
}
