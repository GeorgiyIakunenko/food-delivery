package handler

import (
	"database/sql"
	"encoding/json"
	"food_delivery/repository/db"
	"food_delivery/repository/models"
	"food_delivery/server/request"
	"food_delivery/server/response"
	"net/http"
)

type UserHandler struct {
	db *sql.DB
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{
		db: db,
	}
}

func (u *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := db.NewUserRepository(u.db).GetAll()
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, users)
}

func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var registerRequest request.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&registerRequest)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	user := models.User{
		FirstName: registerRequest.FirstName,
		LastName:  registerRequest.LastName,
		Username:  registerRequest.Username,
		Age:       registerRequest.Age,
		Email:     registerRequest.Email,
		Phone:     registerRequest.Phone,
		Password:  registerRequest.Password,
		CreatedAt: registerRequest.CreatedAt,
	}

	// Create the address
	address := request.Address{
		City:         registerRequest.Address.City,
		PostalCode:   registerRequest.Address.PostalCode,
		AddressLine1: registerRequest.Address.AddressLine1,
		AddressLine2: registerRequest.Address.AddressLine2,
		Country:      registerRequest.Address.Country,
	}

	addressID, err := db.NewAddressRepository(u.db).Create(address)
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	user.AddressID = addressID

	err = db.NewUserRepository(u.db).Create(&user)
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, "user added successfully")

}
