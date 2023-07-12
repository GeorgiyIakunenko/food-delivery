package handler

import (
	"encoding/json"
	"food_delivery/config"
	"food_delivery/server/request"
	"food_delivery/server/response"
	"food_delivery/service"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AuthHandler struct {
	cfg          *config.Config
	UserServiceI service.UserServiceI
}

func NewAuthHandler(UserServiceI service.UserServiceI, cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		cfg:          cfg,
		UserServiceI: UserServiceI,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	req := new(request.LoginRequest)

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	user, err := h.UserServiceI.GetUserByEmail(req.Email)
	if err != nil {
		response.SendInvalidCredentials(w)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		response.SendInvalidCredentials(w)
		return
	}

	accessString, err := service.GenerateToken(int(user.ID), h.cfg.AccessLifetimeMinutes, h.cfg.AccessSecret)
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	refreshString, err := service.GenerateToken(int(user.ID), h.cfg.RefreshLifetimeMinutes, h.cfg.RefreshSecret)
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	resp := response.TokenResponse{
		AccessToken:  accessString,
		RefreshToken: refreshString,
	}

	response.SendOK(w, resp)
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	req := new(request.RegisterRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	_, err := h.UserServiceI.RegisterUser(*req)
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, "user was successfully registered")
}

func (h *AuthHandler) GetTokenPair(w http.ResponseWriter, r *http.Request) {
	UserRefreshToken := r.Header.Get("Authorization")
	user, err := service.ValidateToken(UserRefreshToken, h.cfg.RefreshSecret)
	if err != nil {
		response.SendInvalidCredentials(w)
		return
	}

	accessToken, err := service.GenerateToken(int(user.ID), h.cfg.AccessLifetimeMinutes, h.cfg.AccessSecret)
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	refreshToken, err := service.GenerateToken(int(user.ID), h.cfg.RefreshLifetimeMinutes, h.cfg.RefreshSecret)
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	refreshResponse := response.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	response.SendOK(w, refreshResponse)
}
