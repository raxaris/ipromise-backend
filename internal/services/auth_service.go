package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/repositories"
)

type AuthService interface {
	Signup(req dto.SignupRequest) error
	Login(req dto.LoginRequest) (string, string, error)
	Refresh(refreshToken string) (string, error)
	Logout(refreshToken string) error
}

type authService struct {
	userRepo  repositories.UserRepository
	tokenRepo repositories.TokenRepository
}

func NewAuthService(userRepo repositories.UserRepository, tokenRepo repositories.TokenRepository) AuthService {
	return &authService{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
	}
}

func (s *authService) Signup(req dto.SignupRequest) error {
	if req.Password != req.ConfirmPassword {
		return errors.New("пароли не совпадают")
	}

	if s.userRepo.IsEmailExists(req.Email) {
		return errors.New("email уже используется")
	}

	if s.userRepo.IsUsernameExists(req.Username) {
		return errors.New("username уже используется")
	}

	user := &models.User{
		ID:       uuid.New(),
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := user.HashPassword(); err != nil {
		return errors.New("ошибка хеширования пароля")
	}

	return s.userRepo.CreateUser(user)
}

func (s *authService) Login(req dto.LoginRequest) (string, string, error) {
	user, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil || !user.CheckPassword(req.Password) {
		return "", "", errors.New("неверный email или пароль")
	}

	accessToken, err := GenerateAccessToken(user.ID.String(), user.Role)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := GenerateRefreshToken(user.ID.String(), user.Role)
	if err != nil {
		return "", "", err
	}

	token := &models.RefreshToken{
		ID:        uuid.New(),
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	if err := s.tokenRepo.Save(token); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *authService) Refresh(token string) (string, error) {
	rt, err := s.tokenRepo.FindValid(token)
	if err != nil {
		return "", errors.New("refresh-токен недействителен или истёк")
	}

	user, err := s.userRepo.GetUserByID(rt.UserID)
	if err != nil {
		return "", errors.New("пользователь не найден")
	}

	return GenerateAccessToken(user.ID.String(), user.Role)
}

func (s *authService) Logout(refreshToken string) error {
	rt, err := s.tokenRepo.FindValid(refreshToken)
	if err != nil {
		return errors.New("refresh-токен не найден или истёк")
	}
	return s.tokenRepo.Delete(rt)
}
