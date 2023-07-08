package handler

import (
	"database/sql"
	"food_delivery/repository/db"
	"food_delivery/server/response"
	"net/http"
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
