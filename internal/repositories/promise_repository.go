package repositories

import (
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/config"
	"github.com/raxaris/ipromise-backend/internal/models"
)

// CreatePromise – создаёт обещание в БД
func CreatePromise(promise *models.Promise) error {
	return config.DB.Create(promise).Error
}

// GetPromiseByID – получает обещание по ID
func GetPromiseByID(id uuid.UUID) (*models.Promise, error) {
	var promise models.Promise
	if err := config.DB.First(&promise, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &promise, nil
}

// GetPromisesByUserID – получает все обещания конкретного пользователя
func GetPromisesByUserID(userID uuid.UUID) ([]models.Promise, error) {
	var promises []models.Promise
	if err := config.DB.Where("user_id = ?", userID).Find(&promises).Error; err != nil {
		return nil, err
	}
	return promises, nil
}

// GetAllPromises – получение всех обещаний
func GetAllPromises() ([]models.Promise, error) {
	var promises []models.Promise
	err := config.DB.Find(&promises).Error
	return promises, err
}

// GetPublicPromises – возвращает только публичные обещания
func GetPublicPromises() ([]models.Promise, error) {
	var promises []models.Promise
	err := config.DB.Where("is_private = ?", false).Find(&promises).Error
	return promises, err
}

// UpdatePromise – обновляет обещание (например, меняет статус)
func UpdatePromise(promise *models.Promise) error {
	return config.DB.Save(promise).Error
}

// DeletePromise – мягкое удаление обещания (soft-delete)
func DeletePromise(id uuid.UUID) error {
	return config.DB.Delete(&models.Promise{}, "id = ?", id).Error
}
