package handler

import (
	"database/sql"
	"food_delivery/repository/db"
	"food_delivery/server/response"
	"net/http"
)

type ProductHandler struct {
	db *sql.DB
}

func NewProductHandler(db *sql.DB) *ProductHandler {
	return &ProductHandler{
		db: db,
	}
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := db.NewProductRepository(h.db).GetAll()
	if err != nil {
		response.SendServerError(w, err)
		return
	}
	response.SendOK(w, products)
}
