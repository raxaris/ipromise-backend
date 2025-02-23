package services

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
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
	// Убираем пробелы
	username = strings.TrimSpace(username)
	email = strings.TrimSpace(email)

	// Проверяем уникальность email и username
	if repositories.IsEmailExists(email) {
		return nil, ErrEmailTaken
	}
	if repositories.IsUsernameExists(username) {
		return nil, ErrUsernameTaken
	}

	// Проверяем длину username (не менее 3 символов)
	if len(username) < 3 {
		return nil, errors.New("имя пользователя должно содержать минимум 3 символа")
	}

	user := &models.User{
		ID:       uuid.New(),
		Username: username,
		Email:    email,
		Password: password,
	}

	// Хешируем пароль
	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	// Создаём пользователя в БД
	err := repositories.CreateUser(user)
	return user, err
}

// GetAllUsers – получение всех пользователей
func GetAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}

// GetUserByID – получение пользователя по ID
func GetUserByID(userID uuid.UUID) (*models.User, error) {
	user, err := repositories.GetUserByID(userID)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

// GetUserByUsername – получение пользователя по username
func GetUserByUsername(username string) (*models.User, error) {
	user, err := repositories.GetUserByUsername(username)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return user, nil
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
	if req.Username != nil {
		newUsername := strings.TrimSpace(*req.Username)
		if newUsername != existingUser.Username {
			if repositories.IsUsernameExists(newUsername) {
				return ErrUsernameTaken
			}
			existingUser.Username = newUsername
		}
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
