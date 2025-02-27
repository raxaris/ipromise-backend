package repositories

import (
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/config"
	"github.com/raxaris/ipromise-backend/internal/models"
)

// CreateUser – создание пользователя в БД
func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

// GetUserByID – получение пользователя по ID
func GetUserByID(userID uuid.UUID) (*models.User, error) {
	var user models.User
	err := config.DB.First(&user, "id = ?", userID).Error
	if err != nil {
		return nil, err // Если пользователь не найден, возвращаем nil и ошибку
	}
	return &user, nil
}

// GetUserByUsername – получение пользователя по ID
func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := config.DB.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err // Если пользователь не найден, возвращаем nil и ошибку
	}
	return &user, nil
}

// GetAllUsers – получение списка всех пользователей
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := config.DB.Find(&users).Error
	return users, err
}

// UpdateUser – обновление пользователя
func UpdateUser(user *models.User) error {
	return config.DB.Save(user).Error
}

// DeleteUser – удаление пользователя
func DeleteUser(userID uuid.UUID) error {
	return config.DB.Delete(&models.User{}, "id = ?", userID).Error
}

// IsEmailExists – проверка, есть ли email в БД
func IsEmailExists(email string) bool {
	var count int64
	config.DB.Model(&models.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsUsernameExists – проверка, есть ли username в БД
func IsUsernameExists(username string) bool {
	var count int64
	config.DB.Model(&models.User{}).Where("username = ?", username).Count(&count)
	return count > 0
}
