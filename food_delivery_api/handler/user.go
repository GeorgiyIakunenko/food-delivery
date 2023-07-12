package handler

import (
	"food_delivery/config"
	"food_delivery/server/response"
	"food_delivery/service"
	"net/http"
)

type UserHandler struct {
	UserServiceI service.UserServiceI
	cfg          *config.Config
}

func NewUserHandler(UserServiceI service.UserServiceI, cfg *config.Config) *UserHandler {
	return &UserHandler{
		UserServiceI: UserServiceI,
		cfg:          cfg,
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

	claims, err := service.ValidateToken(service.GetTokenFromBearerString(r.Header.Get("Authorization")), h.cfg.AccessSecret)
	if err != nil {
		response.SendInvalidCredentials(w)
		return
	}

	user, err := h.UserServiceI.GetUserByID(claims.ID)
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, user)
}
