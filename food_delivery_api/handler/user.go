package handler

import (
	"database/sql"
	"food_delivery/repository/db"
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

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := db.NewUserRepository(h.db).GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.SendOK(w, users)
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user request.RegisterRequest

	err := request.ParseBody(r, &user)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	userResponse, err := db.NewUserRepository(h.db).RegisterUser(user)
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, userResponse)
}
