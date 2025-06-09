package microtask

import (
	"context"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"gorm.io/gorm"
)

type microtaskRepository struct {
	db *gorm.DB
}

func NewMicrotaskRepository(db *gorm.DB) MicrotaskRepository {
	return &microtaskRepository{db: db}
}

func (r *microtaskRepository) CreateMicrotask(ctx context.Context, microtask *models.Microtask) error {
	return r.db.WithContext(ctx).Create(microtask).Error
}

func (r *microtaskRepository) CreateManyMicrotasks(ctx context.Context, microtasks []models.Microtask) error {
	return r.db.WithContext(ctx).Create(&microtasks).Error
}

func (r *microtaskRepository) GetMicrotaskByID(ctx context.Context, id uuid.UUID) (*models.Microtask, error) {
	var microtask models.Microtask
	if err := r.db.WithContext(ctx).First(&microtask, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &microtask, nil
}

func (r *microtaskRepository) GetAllMicrotasks(ctx context.Context) ([]models.Microtask, error) {
	var microtasks []models.Microtask
	err := r.db.WithContext(ctx).Find(&microtasks).Error
	return microtasks, err
}

func (r *microtaskRepository) UpdateMicrotask(ctx context.Context, microtask *models.Microtask) error {
	return r.db.WithContext(ctx).Save(microtask).Error
}

func (r *microtaskRepository) DeleteMicrotask(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.Microtask{}, "id = ?", id).Error
}

func (r *microtaskRepository) ListMicrotasksByPromiseID(
	ctx context.Context,
	promiseID uuid.UUID,
) ([]models.Microtask, error) {
	var microtasks []models.Microtask
	err := r.db.WithContext(ctx).
		Where("promise_id = ?", promiseID).
		Order("microtask_order ASC").
		Find(&microtasks).Error
	return microtasks, err
}

func (r *microtaskRepository) ListMicrotasksByStatus(
	ctx context.Context,
	promiseID uuid.UUID,
	status string,
) ([]models.Microtask, error) {
	var microtasks []models.Microtask
	err := r.db.WithContext(ctx).
		Where("promise_id = ? AND status = ?", promiseID, status).
		Order("microtask_order ASC").
		Find(&microtasks).Error
	return microtasks, err
}

func (r *microtaskRepository) AreAllMicrotasksCompleted(
	ctx context.Context,
	promiseID uuid.UUID,
) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Microtask{}).
		Where("promise_id = ? AND status != ?", promiseID, "completed").
		Count(&count).Error

	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func (r *microtaskRepository) CountMicrotasksByPromiseID(
	ctx context.Context,
	promiseID uuid.UUID,
) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Microtask{}).
		Where("promise_id = ?", promiseID).
		Count(&count).Error
	return count, err
}

func (r *microtaskRepository) GetMaxOrderByPromiseID(ctx context.Context, promiseID uuid.UUID) (int, error) {
	var maxOrder int
	err := r.db.WithContext(ctx).
		Model(&models.Microtask{}).
		Where("promise_id = ?", promiseID).
		Select("COALESCE(MAX(microtask_order), 0)"). // если нет — вернёт 0
		Scan(&maxOrder).Error
	return maxOrder, err
}

func (r *microtaskRepository) ReorderMicrotasks(
	ctx context.Context,
	promiseID uuid.UUID,
	orders map[uuid.UUID]int,
) error {
	tx := r.db.WithContext(ctx).Begin()

	for id, order := range orders {
		if err := tx.Model(&models.Microtask{}).
			Where("id = ? AND promise_id = ?", id, promiseID).
			Update("microtask_order", order).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (r *microtaskRepository) CreateManyMicrotasksTx(ctx context.Context, tx *gorm.DB, microtasks []models.Microtask) error {
	return tx.WithContext(ctx).Create(&microtasks).Error
}
