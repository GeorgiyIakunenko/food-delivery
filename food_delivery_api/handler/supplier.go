package handler

import (
	"database/sql"
	"food_delivery/repository/db"
	"food_delivery/server/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type SupplierHandler struct {
	db *sql.DB
}

func NewSupplierHandler(db *sql.DB) *SupplierHandler {
	return &SupplierHandler{
		db: db,
	}
}

func (h *SupplierHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	suppliers, err := db.NewSupplierRepository(h.db).GetAll()
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, suppliers)
}

func (h *SupplierHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)["id"]

	id, err := strconv.ParseInt(vars, 10, 64)
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	supplier, err := db.NewSupplierRepository(h.db).GetByID(id)
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, supplier)
}
