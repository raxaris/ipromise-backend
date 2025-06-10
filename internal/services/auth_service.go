package services

import (
	"context"
	"errors"
	"github.com/raxaris/ipromise-backend/internal/utils"
	"time"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/repositories/token"
	"github.com/raxaris/ipromise-backend/internal/repositories/user"
)

type AuthService interface {
	Signup(ctx context.Context, req dto.SignupRequest) error
	Login(ctx context.Context, req dto.LoginRequest) (string, string, error)
	Refresh(ctx context.Context, refreshToken string) (string, error)
	Logout(ctx context.Context, refreshToken string) error
}

type authService struct {
	userRepo  user.UserRepository
	tokenRepo token.TokenRepository
}

func NewAuthService(userRepo user.UserRepository, tokenRepo token.TokenRepository) AuthService {
	return &authService{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
	}
}

func (s *authService) Signup(ctx context.Context, req dto.SignupRequest) error {
	if req.Password != req.ConfirmPassword {
		return errors.New("password do not match")
	}

	emailExists, err := s.userRepo.IsEmailExists(ctx, req.Email)
	if err != nil {
		return errors.New("email validation error: " + err.Error())
	}
	if emailExists {
		return errors.New("email already exists")
	}

	usernameExists, err := s.userRepo.IsUsernameExists(ctx, req.Username)
	if err != nil {
		return errors.New("username validation error: " + err.Error())
	}
	if usernameExists {
		return errors.New("username already exists")
	}

	newUser := &models.User{
		ID:        uuid.New(),
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		AvatarURL: "",
		Bio:       "",
	}

	if err := newUser.HashPassword(); err != nil {
		return errors.New("password hashing error: " + err.Error())
	}

	return s.userRepo.CreateUser(ctx, newUser)
}

func (s *authService) Login(ctx context.Context, req dto.LoginRequest) (string, string, error) {
	existingUser, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil || !existingUser.CheckPassword(req.Password) {
		return "", "", errors.New("invalid email or password")
	}

	accessToken, err := utils.GenerateAccessToken(existingUser.ID.String(), existingUser.Role)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.GenerateRefreshToken(existingUser.ID.String(), existingUser.Role)
	if err != nil {
		return "", "", err
	}

	token := &models.RefreshToken{
		ID:        uuid.New(),
		UserID:    existingUser.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	if err := s.tokenRepo.Save(ctx, token); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *authService) Refresh(ctx context.Context, token string) (string, error) {
	rt, err := s.tokenRepo.FindValid(ctx, token)
	if err != nil {
		return "", errors.New("invalid refresh-token")
	}

	existingUser, err := s.userRepo.GetUserByID(ctx, rt.UserID)
	if err != nil {
		return "", errors.New("user not found")
	}

	return utils.GenerateAccessToken(existingUser.ID.String(), existingUser.Role)
}

func (s *authService) Logout(ctx context.Context, refreshToken string) error {
	rt, err := s.tokenRepo.FindValid(ctx, refreshToken)
	if err != nil {
		return errors.New("invalid refresh-token")
	}
	return s.tokenRepo.Delete(ctx, rt)
}
