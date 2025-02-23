package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/repositories"
)

// Ошибки
var (
	ErrUserNotFound   = errors.New("Пользователь не найден")
	ErrForbidden      = errors.New("У вас нет прав для выполнения этого действия")
	ErrEmailExists    = errors.New("Email уже используется")
	ErrUsernameExists = errors.New("Username уже используется")
)

// CreateUser – создание пользователя
func CreateUser(username, email, password string) (*models.User, error) {
	if repositories.IsEmailExists(email) {
		return nil, ErrEmailExists
	}
	if repositories.IsUsernameExists(username) {
		return nil, ErrUsernameExists
	}

	user := &models.User{
		ID:       uuid.New(),
		Username: username,
		Email:    email,
		Password: password,
	}

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	err := repositories.CreateUser(user)
	return user, err
}

// GetAllUsers – получение всех пользователей
func GetAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}

// UpdateUser – обновление пользователя
func UpdateUser(userID string, newUserData *models.User) error {
	existingUser, err := repositories.GetUserByID(userID)
	if err != nil {
		return ErrUserNotFound
	}

	existingUser.Username = newUserData.Username
	existingUser.Email = newUserData.Email

	return repositories.UpdateUser(existingUser)
}

// DeleteUser – удаление пользователя
func DeleteUser(userID string) error {
	return repositories.DeleteUser(userID)
}
