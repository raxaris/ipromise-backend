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

	user.Email = strings.TrimSpace(user.Email)
	user.Username = strings.TrimSpace(user.Username)

	if len(user.Username) < 3 {
		return errors.New("Username must be at least 3 characters long")
	}

	emailExists, err := s.userRepo.IsEmailExists(ctx, user.Email)
	if err != nil {
		return err
	}
	if emailExists {
		return errors.New("Email already taken")
	}

	usernameExists, err := s.userRepo.IsUsernameExists(ctx, user.Username)
	if err != nil {
		return err
	}
	if usernameExists {
		return errors.New("Username already taken")
	}

	if err := user.HashPassword(); err != nil {
		return err
	}

	user.ID = uuid.New()

	return s.userRepo.CreateUser(ctx, user)
}

func (s *userService) GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return s.userRepo.GetUserByID(ctx, id)
}

func (s *userService) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	return s.userRepo.GetUserByUsername(ctx, username)
}

func (s *userService) UpdateUser(ctx context.Context, requesterID, userID uuid.UUID, req dto.UpdateUserRequest, isAdmin bool) error {
	existingUser, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	if requesterID != userID && !isAdmin {
		return errors.New("Not enough permission")
	}

	if req.Username != nil {
		newName := strings.TrimSpace(*req.Username)
		if newName != existingUser.Username {
			exists, err := s.userRepo.IsUsernameExists(ctx, newName)
			if err != nil {
				return err
			}
			if exists {
				return errors.New("Username already taken")
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

	return s.userRepo.UpdateUser(ctx, existingUser)
}

func (s *userService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return s.userRepo.DeleteUser(ctx, userID)
}

func (s *userService) ListAllUsers(ctx context.Context) ([]models.User, error) {
	return s.userRepo.GetAllUsers(ctx)
}
