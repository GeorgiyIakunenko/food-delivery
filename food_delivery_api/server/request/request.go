package request

import "time"

type RegisterRequest struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	Address   Address   `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

type Address struct {
	City         string `json:"city"`
	PostalCode   int    `json:"postal_code"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	Country      string `json:"country"`
}
