package user

import (
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(userID uuid.UUID) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(userID uuid.UUID) error
	IsEmailExists(email string) (bool, error)
	IsUsernameExists(username string) (bool, error)
}
