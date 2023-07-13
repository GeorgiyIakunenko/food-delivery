package handler

import (
	"encoding/json"
	"food_delivery/config"
	"food_delivery/server/request"
	"food_delivery/server/response"
	"food_delivery/service"
	"net/http"
)

type contextKey string

const (
	refreshTokenContextKey contextKey = "refreshToken"
)

type AuthHandler struct {
	cfg *config.Config
	//UserServiceI service.UserServiceI
	AuthServiceI service.AuthServiceI
}

func NewAuthHandler(AuthServiceI service.AuthServiceI, cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		cfg:          cfg,
		AuthServiceI: AuthServiceI,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	req := new(request.LoginRequest)

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	resp, err := h.AuthServiceI.Login(*req)
	if err != nil {
		response.SendInvalidCredentials(w)
		return
	}

	response.SendOK(w, resp)
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	req := new(request.RegisterRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	err := h.AuthServiceI.Register(*req)
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, "user was successfully registered")
}

func (h *AuthHandler) GetTokenPair(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(config.NewConfig().RefreshSecret).(*service.JwtCustomClaims)

	refreshResponse, err := h.AuthServiceI.GetTokenPair(claims.ID)
	if err != nil {
		response.SendInvalidCredentials(w)
		return
	}

	response.SendOK(w, refreshResponse)
}
