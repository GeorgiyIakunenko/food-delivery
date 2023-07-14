package handler

import (
	"encoding/json"
	"food_delivery/config"
	"food_delivery/server/request"
	"food_delivery/server/response"
	"food_delivery/service"
	"net/http"
)

type UserHandler struct {
	cfg          *config.Config
	UserServiceI service.UserServiceI
}

func NewUserHandler(UserServiceI service.UserServiceI, cfg *config.Config) *UserHandler {
	return &UserHandler{
		cfg:          cfg,
		UserServiceI: UserServiceI,
	}
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserServiceI.GetAll()
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, users)
}

func (h *UserHandler) GetUserProfile(w http.ResponseWriter, r *http.Request) {

	claims := r.Context().Value(config.NewConfig().AccessSecret).(*service.JwtCustomClaims)

	user, err := h.UserServiceI.GetUserByID(claims.ID)
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, user)
}

func (h *UserHandler) UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	req := new(request.UpdateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	claims := r.Context().Value(config.NewConfig().AccessSecret).(*service.JwtCustomClaims)

	err := h.UserServiceI.UpdateUserProfile(claims.ID, *req)
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, "user profile was successfully updated")
}
