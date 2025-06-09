package prediction

import (
	"context"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"gorm.io/gorm"
)

type PredictionRepository interface {
	GetByPromiseID(ctx context.Context, promiseID uuid.UUID) (*models.Prediction, error)
	Create(ctx context.Context, prediction *models.Prediction) error
	CreateTx(ctx context.Context, tx *gorm.DB, prediction *models.Prediction) error
	Update(ctx context.Context, prediction *models.Prediction) error
	DeleteByPromiseID(ctx context.Context, promiseID uuid.UUID) error
	DeleteByPromiseIDTx(ctx context.Context, tx *gorm.DB, promiseID uuid.UUID) error
}
