package handler

import (
	"database/sql"
	"food_delivery/repository/db"
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
