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
	HasChild(promiseID uuid.UUID) (bool, error)
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

func (r *promiseRepo) GetAllDescendants(parentID uuid.UUID) ([]models.Promise, error) {
	var descendants []models.Promise

	var fetchChildren func(uuid.UUID) error
	fetchChildren = func(currentID uuid.UUID) error {
		children, err := r.GetChildren(currentID)
		if err != nil {
			return err
		}
		for _, child := range children {
			descendants = append(descendants, child)
			if err := fetchChildren(child.ID); err != nil {
				return err
			}
		}
		return nil
	}

	err := fetchChildren(parentID)
	return descendants, err
}

// Update – обновить обещание
func (r *promiseRepo) Update(promise *models.Promise) error {
	return r.db.Save(promise).Error
}

// Delete – удалить обещание
func (r *promiseRepo) Delete(id uuid.UUID) error {
	if err := r.db.Where("parent_id = ?", id).Delete(&models.Promise{}).Error; err != nil {
		return err
	}

	return r.db.Where("id = ?", id).Delete(&models.Promise{}).Error
}

func (r *promiseRepo) HasChild(promiseID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&models.Promise{}).Where("parent_id = ?", promiseID).Count(&count).Error
	return count > 0, err
}
