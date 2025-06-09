package prediction

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"gorm.io/gorm"
)

type predictionRepository struct {
	db *gorm.DB
}

func NewPredictionRepository(db *gorm.DB) PredictionRepository {
	return &predictionRepository{db: db}
}

func (r *predictionRepository) GetByPromiseID(ctx context.Context, promiseID uuid.UUID) (*models.Prediction, error) {
	var prediction models.Prediction
	err := r.db.WithContext(ctx).First(&prediction, "promise_id = ?", promiseID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &prediction, nil
}

func (r *predictionRepository) Create(ctx context.Context, prediction *models.Prediction) error {
	return r.db.WithContext(ctx).Create(prediction).Error
}

func (r *predictionRepository) Update(ctx context.Context, prediction *models.Prediction) error {
	return r.db.WithContext(ctx).Save(prediction).Error
}

func (r *predictionRepository) DeleteByPromiseID(ctx context.Context, promiseID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("promise_id = ?", promiseID).
		Delete(&models.Prediction{}).Error
}

func (r *predictionRepository) CreateTx(ctx context.Context, tx *gorm.DB, prediction *models.Prediction) error {
	return tx.WithContext(ctx).Create(prediction).Error
}

func (r *predictionRepository) DeleteByPromiseIDTx(ctx context.Context, tx *gorm.DB, promiseID uuid.UUID) error {
	return tx.WithContext(ctx).
		Where("promise_id = ?", promiseID).
		Delete(&models.Prediction{}).Error
}
