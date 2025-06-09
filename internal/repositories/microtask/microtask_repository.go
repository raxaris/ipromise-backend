package microtask

import (
	"context"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"gorm.io/gorm"
)

type MicrotaskRepository interface {
	CreateMicrotask(ctx context.Context, microtask *models.Microtask) error
	CreateManyMicrotasks(ctx context.Context, microtasks []models.Microtask) error
	GetMicrotaskByID(ctx context.Context, id uuid.UUID) (*models.Microtask, error)
	GetAllMicrotasks(ctx context.Context) ([]models.Microtask, error)
	UpdateMicrotask(ctx context.Context, microtask *models.Microtask) error
	DeleteMicrotask(ctx context.Context, id uuid.UUID) error
	ListMicrotasksByPromiseID(ctx context.Context, promiseID uuid.UUID) ([]models.Microtask, error)
	ListMicrotasksByStatus(ctx context.Context, promiseID uuid.UUID, status string) ([]models.Microtask, error)
	AreAllMicrotasksCompleted(ctx context.Context, promiseID uuid.UUID) (bool, error)
	CountMicrotasksByPromiseID(ctx context.Context, promiseID uuid.UUID) (int64, error)
	ReorderMicrotasks(ctx context.Context, promiseID uuid.UUID, orders map[uuid.UUID]int) error
	GetMaxOrderByPromiseID(ctx context.Context, promiseID uuid.UUID) (int, error)
	CreateManyMicrotasksTx(ctx context.Context, tx *gorm.DB, microtasks []models.Microtask) error
}
