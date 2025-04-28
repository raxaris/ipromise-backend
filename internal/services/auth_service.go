package services

import (
	"errors"
	"github.com/raxaris/ipromise-backend/internal/repositories/token"
	"github.com/raxaris/ipromise-backend/internal/repositories/user"
	"time"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/models"
)

type AuthService interface {
	Signup(req dto.SignupRequest) error
	Login(req dto.LoginRequest) (string, string, error)
	Refresh(refreshToken string) (string, error)
	Logout(refreshToken string) error
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

func (s *authService) Signup(req dto.SignupRequest) error {
	if req.Password != req.ConfirmPassword {
		return errors.New("пароли не совпадают")
	}

	// Проверка email
	emailExists, err := s.userRepo.IsEmailExists(req.Email)
	if err != nil {
		return errors.New("ошибка проверки email: " + err.Error())
	}
	if emailExists {
		return errors.New("email уже используется")
	}

	// Проверка username
	usernameExists, err := s.userRepo.IsUsernameExists(req.Username)
	if err != nil {
		return errors.New("ошибка проверки username: " + err.Error())
	}
	if usernameExists {
		return errors.New("username уже используется")
	}

	newUser := &models.User{
		ID:       uuid.New(),
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := newUser.HashPassword(); err != nil {
		return errors.New("ошибка хеширования пароля")
	}

	return s.userRepo.CreateUser(newUser)
}

func (s *authService) Login(req dto.LoginRequest) (string, string, error) {
	existingUser, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil || !existingUser.CheckPassword(req.Password) {
		return "", "", errors.New("неверный email или пароль")
	}

	accessToken, err := GenerateAccessToken(existingUser.ID.String(), existingUser.Role)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := GenerateRefreshToken(existingUser.ID.String(), existingUser.Role)
	if err != nil {
		return "", "", err
	}

	token := &models.RefreshToken{
		ID:        uuid.New(),
		UserID:    existingUser.ID,
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

	existingUser, err := s.userRepo.GetUserByID(rt.UserID)
	if err != nil {
		return "", errors.New("пользователь не найден")
	}

	return GenerateAccessToken(existingUser.ID.String(), existingUser.Role)
}

func (s *authService) Logout(refreshToken string) error {
	rt, err := s.tokenRepo.FindValid(refreshToken)
	if err != nil {
		return errors.New("refresh-токен не найден или истёк")
	}
	return s.tokenRepo.Delete(rt)
}
