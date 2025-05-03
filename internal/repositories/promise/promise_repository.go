package promise

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
)

type PromiseRepository interface {
	CreatePromise(ctx context.Context, promise *models.Promise) error
	GetPromiseByID(ctx context.Context, id uuid.UUID) (*models.Promise, error)
	UpdatePromise(ctx context.Context, promise *models.Promise) error
	DeletePromise(ctx context.Context, id uuid.UUID) error

	ListAllPromisesByUserID(
		ctx context.Context,
		userID uuid.UUID,
		limit int,
		afterCreatedAt *time.Time,
	) ([]models.Promise, error)

	ListPublicPromisesByUserID(
		ctx context.Context,
		userID uuid.UUID,
		limit int,
		afterCreatedAt *time.Time,
	) ([]models.Promise, error)

	ListFeedPromises(
		ctx context.Context,
		userIDs []uuid.UUID, // подписки
		limit int,
		afterCreatedAt *time.Time, // пагинация
	) ([]models.Promise, error)

	ListPublicPromises(
		ctx context.Context,
		limit int,
		afterCreatedAt *time.Time,
	) ([]models.Promise, error)

	CountAllPromisesByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) (int64, error)

	CountPublicPromisesByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) (int64, error)

	ListPromisesBeforeDeadline(
		ctx context.Context,
		before time.Time,
	) ([]models.Promise, error)
}
