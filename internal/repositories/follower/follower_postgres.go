package follower

import (
	"context"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"gorm.io/gorm"
)

type followerRepository struct {
	db *gorm.DB
}

func NewFollowerRepository(db *gorm.DB) FollowerRepository {
	return &followerRepository{db: db}
}

// ✅ Запрос на подписку (pending)
func (r *followerRepository) RequestFollow(ctx context.Context, followerID, followingID uuid.UUID) error {
	follower := models.Follower{
		FollowerID:  followerID,
		FollowingID: followingID,
		Status:      "pending",
	}
	return r.db.WithContext(ctx).Create(&follower).Error
}

// ✅ Принять запрос (accepted)
func (r *followerRepository) AcceptFollowRequest(ctx context.Context, followerID, followingID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Model(&models.Follower{}).
		Where("follower_id = ? AND following_id = ? AND status = ?", followerID, followingID, "pending").
		Update("status", "accepted").Error
}

// ✅ Отклонить запрос (удаляем)
func (r *followerRepository) DeclineFollowRequest(ctx context.Context, followerID, followingID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("follower_id = ? AND following_id = ? AND status = ?", followerID, followingID, "pending").
		Delete(&models.Follower{}).Error
}

// ✅ Отписаться (удалить подписку)
func (r *followerRepository) Unfollow(ctx context.Context, followerID, followingID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("follower_id = ? AND following_id = ?", followerID, followingID).
		Delete(&models.Follower{}).Error
}

// ✅ Проверить, подписан ли я (accepted)
func (r *followerRepository) IsFollowing(ctx context.Context, followerID, followingID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Follower{}).
		Where("follower_id = ? AND following_id = ? AND status = ?", followerID, followingID, "accepted").
		Count(&count).Error
	return count > 0, err
}

// ✅ Получить список подписчиков (accepted)
func (r *followerRepository) ListFollowers(ctx context.Context, userID uuid.UUID) ([]models.Follower, error) {
	var followers []models.Follower
	err := r.db.WithContext(ctx).
		Where("following_id = ? AND status = ?", userID, "accepted").
		Find(&followers).Error
	return followers, err
}

// ✅ Получить список подписок (accepted)
func (r *followerRepository) ListFollowing(ctx context.Context, userID uuid.UUID) ([]models.Follower, error) {
	var following []models.Follower
	err := r.db.WithContext(ctx).
		Where("follower_id = ? AND status = ?", userID, "accepted").
		Find(&following).Error
	return following, err
}

// ✅ Получить pending-запросы (ждут моего одобрения)
func (r *followerRepository) ListPendingFollowRequests(ctx context.Context, userID uuid.UUID) ([]models.Follower, error) {
	var requests []models.Follower
	err := r.db.WithContext(ctx).
		Where("following_id = ? AND status = ?", userID, "pending").
		Find(&requests).Error
	return requests, err
}

// CountFollowers — количество фолловеров (тех, кто подписан на userID)
func (r *followerRepository) CountFollowers(ctx context.Context, userID uuid.UUID) (int, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Follower{}).
		Where("following_id = ?", userID).
		Count(&count).Error
	return int(count), err
}

// CountFollowing — количество подписок (на кого подписан userID)
func (r *followerRepository) CountFollowing(ctx context.Context, userID uuid.UUID) (int, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Follower{}).
		Where("follower_id = ?", userID).
		Count(&count).Error
	return int(count), err
}
