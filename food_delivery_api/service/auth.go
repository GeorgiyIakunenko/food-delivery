package service

import (
	"context"
	"errors"
	"fmt"
	"food_delivery/config"
	"food_delivery/server/request"
	"food_delivery/server/response"
	"food_delivery/utils"
	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type AuthService struct {
	UserServiceI UserServiceI
	redisClient  *redis.Client
	cfg          *config.Config
}

type AuthServiceI interface {
	Login(req request.LoginRequest) (*response.TokenResponse, error)
	Register(req request.RegisterRequest) error
	GetTokenPair(userID int) (*response.TokenResponse, error)
	StoreResetCode(email string, resetCode string) error
	GetResetCodeByUserEmail(email string) (string, error)
	InitiatePasswordReset(req request.PasswordResetRequest) error
	SubmitResetCode(req request.PasswordResetRequest) error
}

func NewAuthService(UserServiceI UserServiceI, redisClient *redis.Client, cfg *config.Config) AuthServiceI {
	return &AuthService{
		UserServiceI: UserServiceI,
		redisClient:  redisClient,
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

func (h *AuthService) StoreResetCode(email string, resetCode string) error {
	ctx := context.Background()

	key := fmt.Sprintf("reset_code:%s", email)
	value := resetCode

	err := h.redisClient.Set(ctx, key, value, time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func (h *AuthService) GetResetCodeByUserEmail(email string) (string, error) {
	key := fmt.Sprintf("reset_code:%s", email)

	resetCode, err := h.redisClient.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", errors.New("Invalid or expired reset code")
		}
		return "", err
	}

	return resetCode, nil
}

func (h *AuthService) InitiatePasswordReset(req request.PasswordResetRequest) error {

	user, err := h.UserServiceI.GetUserByEmail(req.Email)
	if err != nil {
		log.Printf("Failed to get user by email: %s", err)
		return fmt.Errorf("failed to initiate password reset: user not found")
	}

	resetCode := utils.GenerateResetCode()

	err = h.StoreResetCode(user.Email, resetCode)
	if err != nil {
		log.Printf("Failed to store reset code: %s", err)
		return fmt.Errorf("failed to initiate password reset: internal server error")
	}

	err = utils.SendResetCodeEmail(req.Email, resetCode, h.cfg)
	if err != nil {
		log.Printf("Failed to send reset code email: %s", err)
		return fmt.Errorf("failed to initiate password reset: internal server error")
	}

	return nil
}

func (h *AuthService) SubmitResetCode(req request.PasswordResetRequest) error {
	resetCode, err := h.GetResetCodeByUserEmail(req.Email)
	if err != nil {
		return fmt.Errorf("failed to retrieve reset code: %w", err)
	}

	if resetCode != req.ResetCode {
		return fmt.Errorf("invalid reset code: expected '%s' but received '%s'", resetCode, req.ResetCode)
	}

	user, err := h.UserServiceI.GetUserByEmail(req.Email)
	if err != nil {
		return fmt.Errorf("failed to retrieve user: %w", err)
	}

	err = h.UserServiceI.UpdateUserPasswordById(int(user.ID), req.NewPassword)
	if err != nil {
		return fmt.Errorf("failed to update user password: %w", err)
	}

	return nil
}
