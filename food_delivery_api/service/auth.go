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
	StoreTokensByUserID(userID int, tokens *response.TokenResponse) error
	GetTokensByUserID(userID int) (*response.TokenResponse, error)
	DeleteTokensByUserID(userID int) error
	Logout(userID int) error
	ChangePassword(req request.ChangePasswordRequest, cfg *config.Config) (*response.TokenResponse, error)
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

	tokens := response.TokenResponse{
		AccessToken:  accessString,
		RefreshToken: refreshString,
	}

	// Store the tokens in Redis
	err = h.StoreTokensByUserID(int(user.ID), &tokens)
	if err != nil {
		return nil, err
	}

	return &tokens, nil

}

func (h *AuthService) Register(req request.RegisterRequest) error {

	_, err := h.UserServiceI.RegisterUser(req)
	if err != nil {
		return err
	}

	return nil
}

func (h *AuthService) GetTokenPair(userID int) (*response.TokenResponse, error) {

	err := h.DeleteTokensByUserID(userID)
	if err != nil {
		return nil, err
	}

	accessToken, err := GenerateToken(userID, h.cfg.AccessLifetimeMinutes, h.cfg.AccessSecret)
	if err != nil {
		return nil, err
	}

	refreshToken, err := GenerateToken(userID, h.cfg.RefreshLifetimeMinutes, h.cfg.RefreshSecret)
	if err != nil {
		return nil, err
	}

	tokens := response.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	err = h.StoreTokensByUserID(userID, &tokens)
	if err != nil {
		return nil, err
	}

	return &tokens, nil
}

func (h *AuthService) StoreResetCode(email string, resetCode string) error {
	ctx := context.Background()

	key := fmt.Sprintf("reset_code:%s", email)
	value := resetCode

	err := h.redisClient.Set(ctx, key, value, time.Minute*10).Err()
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
		return err
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
		return fmt.Errorf("failed to retrieve reset code")
	}

	if resetCode != req.ResetCode {
		return fmt.Errorf("invalid reset code")
	}

	user, err := h.UserServiceI.GetUserByEmail(req.Email)
	if err != nil {
		return fmt.Errorf("failed to retrieve user")
	}

	err = h.UserServiceI.UpdateUserPasswordById(int(user.ID), req.NewPassword)
	if err != nil {
		return err
	}

	return nil
}

func (h *AuthService) StoreTokensByUserID(userID int, tokens *response.TokenResponse) error {
	accessTokenKey := fmt.Sprintf("access_token:%d", userID)
	refreshTokenKey := fmt.Sprintf("refresh_token:%d", userID)

	ctx := context.Background()

	err := h.redisClient.Set(ctx, accessTokenKey, tokens.AccessToken, time.Duration(h.cfg.AccessLifetimeMinutes)*time.Minute).Err()
	if err != nil {
		return err
	}

	err = h.redisClient.Set(ctx, refreshTokenKey, tokens.RefreshToken, time.Duration(h.cfg.RefreshLifetimeMinutes)*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

func (h *AuthService) GetTokensByUserID(userID int) (*response.TokenResponse, error) {
	accessTokenKey := fmt.Sprintf("access_token:%d", userID)
	refreshTokenKey := fmt.Sprintf("refresh_token:%d", userID)

	ctx := context.Background()

	accessToken, err := h.redisClient.Get(ctx, accessTokenKey).Result()
	if err != nil {
		return nil, err
	}

	refreshToken, err := h.redisClient.Get(ctx, refreshTokenKey).Result()
	if err != nil {
		return nil, err
	}

	resp := response.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &resp, nil
}

func (h *AuthService) DeleteTokensByUserID(userID int) error {
	accessTokenKey := fmt.Sprintf("access_token:%d", userID)
	refreshTokenKey := fmt.Sprintf("refresh_token:%d", userID)

	ctx := context.Background()

	err := h.redisClient.Del(ctx, accessTokenKey).Err()
	if err != nil {
		return err
	}

	err = h.redisClient.Del(ctx, refreshTokenKey).Err()
	if err != nil {
		return err
	}

	return nil
}

func (h *AuthService) Logout(userID int) error {
	err := h.DeleteTokensByUserID(userID)
	if err != nil {
		return err
	}

	return nil
}

func (h *AuthService) ChangePassword(req request.ChangePasswordRequest, cfg *config.Config) (*response.TokenResponse, error) {
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

	accessString, err := GenerateToken(int(user.ID), cfg.AccessLifetimeMinutes, cfg.AccessSecret)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	refreshString, err := GenerateToken(int(user.ID), cfg.RefreshLifetimeMinutes, cfg.RefreshSecret)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	tokens := response.TokenResponse{
		AccessToken:  accessString,
		RefreshToken: refreshString,
	}

	// Delete the existing token pair
	err = h.DeleteTokensByUserID(int(user.ID))
	if err != nil {
		log.Printf("Failed to delete existing token pair: %s", err)
		return nil, errors.New("failed to delete existing token pair")
	}

	// Store the new token pair
	err = h.StoreTokensByUserID(int(user.ID), &tokens)
	if err != nil {
		return nil, errors.New("failed to store token pair")
	}

	return &tokens, nil
}
