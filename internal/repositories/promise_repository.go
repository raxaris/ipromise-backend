package repositories

import (
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"gorm.io/gorm"
)

type PromiseRepository interface {
	Create(promise *models.Promise) error
	GetByID(id uuid.UUID) (*models.Promise, error)
	Update(promise *models.Promise) error
	Delete(id uuid.UUID) error

	GetAll() ([]models.Promise, error)
	GetAllPublic() ([]models.Promise, error)

	GetByUserID(userID uuid.UUID) ([]models.Promise, error)
	GetPublicByUserID(userID uuid.UUID) ([]models.Promise, error)

	GetChildren(parentID uuid.UUID) ([]models.Promise, error)
}

type promiseRepo struct {
	db *gorm.DB
}

func NewPromiseRepository(db *gorm.DB) PromiseRepository {
	return &promiseRepo{db: db}
}

// Create – создать новое обещание
func (r *promiseRepo) Create(promise *models.Promise) error {
	return r.db.Create(promise).Error
}

// GetByID – получить обещание по ID
func (r *promiseRepo) GetByID(id uuid.UUID) (*models.Promise, error) {
	var promise models.Promise
	err := r.db.First(&promise, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &promise, nil
}

// Update – обновить обещание
func (r *promiseRepo) Update(promise *models.Promise) error {
	return r.db.Save(promise).Error
}

// Delete – удалить обещание
func (r *promiseRepo) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Promise{}, "id = ?", id).Error
}

// GetAll – все обещания (для админов)
func (r *promiseRepo) GetAll() ([]models.Promise, error) {
	var promises []models.Promise
	err := r.db.Order("created_at DESC").Find(&promises).Error
	return promises, err
}

// GetAllPublic – все публичные обещания
func (r *promiseRepo) GetAllPublic() ([]models.Promise, error) {
	var promises []models.Promise
	err := r.db.Where("is_private = false").Order("created_at DESC").Find(&promises).Error
	return promises, err
}

// GetByUserID – обещания конкретного пользователя
func (r *promiseRepo) GetByUserID(userID uuid.UUID) ([]models.Promise, error) {
	var promises []models.Promise
	err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&promises).Error
	return promises, err
}

// GetPublicByUserID – публичные обещания пользователя
func (r *promiseRepo) GetPublicByUserID(userID uuid.UUID) ([]models.Promise, error) {
	var promises []models.Promise
	err := r.db.Where("user_id = ? AND is_private = false", userID).Order("created_at DESC").Find(&promises).Error
	return promises, err
}

// GetChildren – получить вложенные обещания (прогресс)
func (r *promiseRepo) GetChildren(parentID uuid.UUID) ([]models.Promise, error) {
	var promises []models.Promise
	err := r.db.Where("parent_id = ?", parentID).Order("created_at ASC").Find(&promises).Error
	return promises, err
}
