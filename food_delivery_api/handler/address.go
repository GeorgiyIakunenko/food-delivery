package handler

import (
	"database/sql"
	"food_delivery/repository/db"
	"food_delivery/server/response"
	"net/http"
)

type AddressHandler struct {
	db *sql.DB
}

func NewAddressHandler(db *sql.DB) *AddressHandler {
	return &AddressHandler{
		db: db,
	}
}

func (a *AddressHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	addresses, err := db.NewAddressRepository(a.db).GetAll()
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, addresses)
}
