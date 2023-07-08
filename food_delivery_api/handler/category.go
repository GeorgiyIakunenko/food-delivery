package handler

import (
	"database/sql"
	"food_delivery/repository/db"
	"food_delivery/server/response"
	"net/http"
)

type CategoryHandler struct {
	db *sql.DB
}

func NewCategoryHandler(db *sql.DB) *CategoryHandler {
	return &CategoryHandler{
		db: db,
	}
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := db.NewCategoryRepository(h.db).GetAll()
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, categories)
}
