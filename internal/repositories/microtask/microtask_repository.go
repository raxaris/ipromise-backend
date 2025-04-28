package microtask

import (
	"context"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
)

type MicrotaskRepository interface {
	CreateMicrotask(ctx context.Context, microtask *models.Microtask) error
	GetMicrotaskByID(ctx context.Context, id uuid.UUID) (*models.Microtask, error)
	UpdateMicrotask(ctx context.Context, microtask *models.Microtask) error
	DeleteMicrotask(ctx context.Context, id uuid.UUID) error

	ListMicrotasksByPromiseID(ctx context.Context, promiseID uuid.UUID) ([]models.Microtask, error)

	ListMicrotasksByStatus(ctx context.Context, promiseID uuid.UUID, status string) ([]models.Microtask, error)

	AreAllMicrotasksCompleted(ctx context.Context, promiseID uuid.UUID) (bool, error)

	CountMicrotasksByPromiseID(ctx context.Context, promiseID uuid.UUID) (int64, error)

	ReorderMicrotasks(ctx context.Context, promiseID uuid.UUID, orders map[uuid.UUID]int) error
}
