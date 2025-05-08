package services

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/repositories/user"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	UpdateUser(ctx context.Context, requesterID, userID uuid.UUID, req dto.UpdateUserRequest, isAdmin bool) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	ListAllUsers(ctx context.Context) ([]models.User, error)
}

type userService struct {
	userRepo user.UserRepository
}

func NewUserService(userRepo user.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(ctx context.Context, user *models.User) error {
	// Trim input
	user.Email = strings.TrimSpace(user.Email)
	user.Username = strings.TrimSpace(user.Username)

	// Валидация
	if len(user.Username) < 3 {
		return errors.New("имя пользователя должно быть не короче 3 символов")
	}

	// Проверка на уникальность
	emailExists, err := s.userRepo.IsEmailExists(user.Email)
	if err != nil {
		return err
	}
	if emailExists {
		return errors.New("email уже используется")
	}

	usernameExists, err := s.userRepo.IsUsernameExists(user.Username)
	if err != nil {
		return err
	}
	if usernameExists {
		return errors.New("имя пользователя уже занято")
	}

	// Хеширование пароля
	if err := user.HashPassword(); err != nil {
		return err
	}

	// Присваиваем ID
	user.ID = uuid.New()

	return s.userRepo.CreateUser(user)
}

func (s *userService) GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *userService) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	return s.userRepo.GetUserByUsername(username)
}

func (s *userService) UpdateUser(ctx context.Context, requesterID, userID uuid.UUID, req dto.UpdateUserRequest, isAdmin bool) error {
	existingUser, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return err
	}

	if requesterID != userID && !isAdmin {
		return errors.New("нет прав на редактирование")
	}

	if req.Username != nil {
		newName := strings.TrimSpace(*req.Username)
		if newName != existingUser.Username {
			exists, err := s.userRepo.IsUsernameExists(newName)
			if err != nil {
				return err
			}
			if exists {
				return errors.New("имя уже занято")
			}
			existingUser.Username = newName
		}
	}

	if req.AvatarURL != nil {
		trimmed := strings.TrimSpace(*req.AvatarURL)
		if len(trimmed) > 0 {
			existingUser.AvatarURL = trimmed
		}
	}

	if isAdmin && req.Role != nil {
		existingUser.Role = *req.Role
	}

	return s.userRepo.UpdateUser(existingUser)
}

func (s *userService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return s.userRepo.DeleteUser(userID)
}

func (s *userService) ListAllUsers(ctx context.Context) ([]models.User, error) {
	return s.userRepo.GetAllUsers()
}
