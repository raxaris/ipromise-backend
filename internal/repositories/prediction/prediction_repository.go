package prediction

import (
	"context"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
)

type PredictionRepository interface {
	GetByPromiseID(ctx context.Context, promiseID uuid.UUID) (*models.Prediction, error)
	Create(ctx context.Context, prediction *models.Prediction) error
	Update(ctx context.Context, prediction *models.Prediction) error
	DeleteByPromiseID(ctx context.Context, promiseID uuid.UUID) error
}
