package promise

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"gorm.io/gorm"
)

type promiseRepository struct {
	db *gorm.DB
}

func NewPromiseRepository(db *gorm.DB) PromiseRepository {
	return &promiseRepository{db: db}
}

func (r *promiseRepository) CreatePromise(ctx context.Context, promise *models.Promise) error {
	return r.db.WithContext(ctx).Create(promise).Error
}

func (r *promiseRepository) GetPromiseByID(ctx context.Context, id uuid.UUID) (*models.Promise, error) {
	var promise models.Promise
	if err := r.db.WithContext(ctx).First(&promise, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &promise, nil
}

func (r *promiseRepository) GetAllPromises(ctx context.Context) ([]models.Promise, error) {
	var promises []models.Promise
	err := r.db.WithContext(ctx).Find(&promises).Error
	return promises, err
}

func (r *promiseRepository) UpdatePromise(ctx context.Context, promise *models.Promise) error {
	return r.db.WithContext(ctx).Save(promise).Error
}

func (r *promiseRepository) DeletePromise(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.Promise{}, "id = ?", id).Error
}

func (r *promiseRepository) ListAllPromisesByUserID(
	ctx context.Context,
	userID uuid.UUID,
	limit int,
	afterCreatedAt *time.Time,
) ([]models.Promise, error) {
	var promises []models.Promise

	query := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit)

	if afterCreatedAt != nil {
		query = query.Where("created_at < ?", *afterCreatedAt)
	}

	if err := query.Find(&promises).Error; err != nil {
		return nil, err
	}

	return promises, nil
}

func (r *promiseRepository) ListPublicPromisesByUserID(
	ctx context.Context,
	userID uuid.UUID,
	limit int,
	afterCreatedAt *time.Time,
) ([]models.Promise, error) {
	var promises []models.Promise

	query := r.db.WithContext(ctx).
		Where("user_id = ? AND is_private = false", userID).
		Order("created_at DESC").
		Limit(limit)

	if afterCreatedAt != nil {
		query = query.Where("created_at < ?", *afterCreatedAt)
	}

	if err := query.Find(&promises).Error; err != nil {
		return nil, err
	}

	return promises, nil
}

func (r *promiseRepository) ListFeedPromises(
	ctx context.Context,
	userIDs []uuid.UUID,
	limit int,
	afterCreatedAt *time.Time,
) ([]models.Promise, error) {
	var promises []models.Promise

	if len(userIDs) == 0 {
		return promises, nil
	}

	query := r.db.WithContext(ctx).
		Where("user_id IN ?", userIDs).
		Order("created_at DESC").
		Limit(limit)

	if afterCreatedAt != nil {
		query = query.Where("created_at < ?", *afterCreatedAt)
	}

	if err := query.Find(&promises).Error; err != nil {
		return nil, err
	}

	return promises, nil
}

func (r *promiseRepository) ListPublicPromises(
	ctx context.Context,
	limit int,
	afterCreatedAt *time.Time,
) ([]models.Promise, error) {
	var promises []models.Promise

	query := r.db.WithContext(ctx).
		Where("is_private = false").
		Order("created_at DESC").
		Limit(limit)

	if afterCreatedAt != nil {
		query = query.Where("created_at < ?", *afterCreatedAt)
	}

	if err := query.Find(&promises).Error; err != nil {
		return nil, err
	}

	return promises, nil
}

func (r *promiseRepository) CountAllPromisesByUserID(
	ctx context.Context,
	userID uuid.UUID,
) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Promise{}).
		Where("user_id = ?", userID).
		Count(&count).Error
	return count, err
}

func (r *promiseRepository) CountPublicPromisesByUserID(
	ctx context.Context,
	userID uuid.UUID,
) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Promise{}).
		Where("user_id = ? AND is_private = false", userID).
		Count(&count).Error
	return count, err
}

func (r *promiseRepository) ListPromisesBeforeDeadline(
	ctx context.Context,
	before time.Time,
) ([]models.Promise, error) {
	var promises []models.Promise

	err := r.db.WithContext(ctx).
		Where("deadline <= ? AND status = ?", before, "in_progress").
		Order("deadline ASC").
		Find(&promises).Error

	if err != nil {
		return nil, err
	}

	return promises, nil
}

func (r *promiseRepository) WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return r.db.WithContext(ctx).Transaction(fn)
}
