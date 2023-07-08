package handler

import (
	"database/sql"
	"encoding/json"
	"food_delivery/repository/db"
	"food_delivery/server/request"
	"food_delivery/server/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

func (a *AddressHandler) Create(w http.ResponseWriter, r *http.Request) {
	var address request.Address

	err := json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	id, err := db.NewAddressRepository(a.db).Create(address)
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendJson(w, 200, id)

}

func (a *AddressHandler) Delete(w http.ResponseWriter, r *http.Request) {
	varsId := mux.Vars(r)["id"]

	id, err := strconv.ParseInt(varsId, 10, 64)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	res, err := db.NewAddressRepository(a.db).DeleteById(id)
	if err != nil {

		response.SendServerError(w, err)
		return
	}

	response.SendJson(w, 200, res)

}
