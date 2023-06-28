package models

import "time"

type User struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	AddressID int64     `json:"address_id"`
	CreatedAt time.Time `json:"created_at"`
}
