package service

import (
	"food_delivery/config"
	"food_delivery/server/request"
	"food_delivery/server/response"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserServiceI UserServiceI
	cfg          *config.Config
}

type AuthServiceI interface {
	Login(req request.LoginRequest) (*response.TokenResponse, error)
	Register(req request.RegisterRequest) error
	GetTokenPair(userID int) (*response.TokenResponse, error)
}

func NewAuthService(UserServiceI UserServiceI, cfg *config.Config) AuthServiceI {
	return &AuthService{
		UserServiceI: UserServiceI,
		cfg:          cfg,
	}
}

func (h *AuthService) Login(req request.LoginRequest) (*response.TokenResponse, error) {
	user, err := h.UserServiceI.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, err
	}

	accessString, err := GenerateToken(int(user.ID), h.cfg.AccessLifetimeMinutes, h.cfg.AccessSecret)
	if err != nil {
		return nil, err
	}

	refreshString, err := GenerateToken(int(user.ID), h.cfg.RefreshLifetimeMinutes, h.cfg.RefreshSecret)
	if err != nil {
		return nil, err
	}

	resp := response.TokenResponse{
		AccessToken:  accessString,
		RefreshToken: refreshString,
	}

	return &resp, nil
}

func (h *AuthService) Register(req request.RegisterRequest) error {

	_, err := h.UserServiceI.RegisterUser(req)
	if err != nil {
		return err
	}

	return nil
}

func (h *AuthService) GetTokenPair(userID int) (*response.TokenResponse, error) {
	accessToken, err := GenerateToken(userID, h.cfg.AccessLifetimeMinutes, h.cfg.AccessSecret)
	if err != nil {
		return nil, err
	}

	refreshToken, err := GenerateToken(userID, h.cfg.RefreshLifetimeMinutes, h.cfg.RefreshSecret)
	if err != nil {
		return nil, err
	}

	refreshResponse := response.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &refreshResponse, nil
}
