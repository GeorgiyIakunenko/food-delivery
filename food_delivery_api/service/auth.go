package service

import (
	"errors"
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
	GetTokenPair(UserRefreshToken string) (*response.TokenResponse, error)
	ResetPassword(req request.ResetPasswordRequest) (*response.TokenResponse, error)
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

func (h *AuthService) GetTokenPair(UserRefreshToken string) (*response.TokenResponse, error) {
	user, err := ValidateToken(UserRefreshToken, h.cfg.RefreshSecret)
	if err != nil {
		return nil, err
	}

	accessToken, err := GenerateToken(int(user.ID), h.cfg.AccessLifetimeMinutes, h.cfg.AccessSecret)
	if err != nil {
		return nil, err
	}

	refreshToken, err := GenerateToken(int(user.ID), h.cfg.RefreshLifetimeMinutes, h.cfg.RefreshSecret)
	if err != nil {
		return nil, err
	}

	refreshResponse := response.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &refreshResponse, nil
}

func (h *AuthService) ResetPassword(req request.ResetPasswordRequest) (*response.TokenResponse, error) {
	user, err := h.UserServiceI.GetUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("failed to retrieve user")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		return nil, errors.New("invalid current password")
	}

	if err = h.UserServiceI.UpdateUserPasswordById(int(user.ID), req.NewPassword); err != nil {
		return nil, errors.New("failed to update password")
	}

	accessString, err := GenerateToken(int(user.ID), h.cfg.AccessLifetimeMinutes, h.cfg.AccessSecret)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	refreshString, err := GenerateToken(int(user.ID), h.cfg.RefreshLifetimeMinutes, h.cfg.RefreshSecret)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	resp := response.TokenResponse{
		AccessToken:  accessString,
		RefreshToken: refreshString,
	}

	return &resp, nil
}
