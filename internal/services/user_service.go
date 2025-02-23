package services

import (
	"errors"
	"github.com/raxaris/ipromise-backend/internal/dto"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/repositories"
)

// Ошибки
var (
	ErrUserNotFound     = errors.New("пользователь не найден")
	ErrUsernameTaken    = errors.New("это имя пользователя уже занято")
	ErrEmailTaken       = errors.New("этот email уже используется")
	ErrNotAllowedToEdit = errors.New("у вас нет прав для редактирования этого пользователя")
)

// CreateUser – создание пользователя
func CreateUser(username, email, password string) (*models.User, error) {
	if repositories.IsEmailExists(email) {
		return nil, ErrEmailTaken
	}
	if repositories.IsUsernameExists(username) {
		return nil, ErrUsernameTaken
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
func UpdateUser(requesterID uuid.UUID, userID uuid.UUID, req *dto.UpdateUserRequest, isAdmin bool) error {
	// Получаем существующего пользователя
	existingUser, err := repositories.GetUserByID(userID)
	if err != nil {
		return ErrUserNotFound
	}

	// Проверяем, имеет ли пользователь право редактировать
	if requesterID != userID && !isAdmin {
		return ErrNotAllowedToEdit
	}

	// Проверяем уникальность username, если его меняют
	if req.Username != nil && *req.Username != existingUser.Username {
		if repositories.IsUsernameExists(*req.Username) {
			return ErrUsernameTaken
		}
		existingUser.Username = *req.Username
	}

	// Админ может менять роль
	if isAdmin && req.Role != nil {
		existingUser.Role = *req.Role
	}

	// Обновляем пользователя в БД
	return repositories.UpdateUser(existingUser)
}

// DeleteUser – удаление пользователя
func DeleteUser(userID uuid.UUID) error {
	return repositories.DeleteUser(userID)
}
