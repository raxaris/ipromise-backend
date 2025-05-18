package follower

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"gorm.io/gorm"
	"time"
)

type followerRepository struct {
	db *gorm.DB
}

func NewFollowerRepository(db *gorm.DB) FollowerRepository {
	return &followerRepository{db: db}
}

func (r *followerRepository) RequestFollow(ctx context.Context, followerID, followingID uuid.UUID) error {
	follower := models.Follower{
		FollowerID:  followerID,
		FollowingID: followingID,
		Status:      "pending",
	}
	return r.db.WithContext(ctx).Create(&follower).Error
}

func (r *followerRepository) AcceptFollowRequest(ctx context.Context, followerID, followingID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Model(&models.Follower{}).
		Where("follower_id = ? AND following_id = ? AND status = ?", followerID, followingID, "pending").
		Update("status", "accepted").Error
}

func (r *followerRepository) DeclineFollowRequest(ctx context.Context, followerID, followingID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("follower_id = ? AND following_id = ? AND status = ?", followerID, followingID, "pending").
		Delete(&models.Follower{}).Error
}

func (r *followerRepository) Unfollow(ctx context.Context, followerID, followingID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("follower_id = ? AND following_id = ?", followerID, followingID).
		Delete(&models.Follower{}).Error
}

func (r *followerRepository) IsFollowing(ctx context.Context, followerID, followingID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Follower{}).
		Where("follower_id = ? AND following_id = ? AND status = ?", followerID, followingID, "accepted").
		Count(&count).Error
	return count > 0, err
}

func (r *followerRepository) ListFollowers(ctx context.Context, userID uuid.UUID) ([]models.Follower, error) {
	var followers []models.Follower
	err := r.db.WithContext(ctx).
		Where("following_id = ? AND status = ?", userID, "accepted").
		Find(&followers).Error
	return followers, err
}

func (r *followerRepository) ListFollowing(ctx context.Context, userID uuid.UUID) ([]models.Follower, error) {
	var following []models.Follower
	err := r.db.WithContext(ctx).
		Where("follower_id = ? AND status = ?", userID, "accepted").
		Find(&following).Error
	return following, err
}

func (r *followerRepository) ListPendingFollowRequests(ctx context.Context, userID uuid.UUID) ([]models.Follower, error) {
	var requests []models.Follower
	err := r.db.WithContext(ctx).
		Where("following_id = ? AND status = ?", userID, "pending").
		Find(&requests).Error
	return requests, err
}

func (r *followerRepository) ListSentFollowRequests(ctx context.Context, userID uuid.UUID) ([]models.Follower, error) {
	var requests []models.Follower
	err := r.db.WithContext(ctx).
		Where("follower_id = ? AND status = ?", userID, "pending").
		Find(&requests).Error
	return requests, err
}

func (r *followerRepository) CountFollowers(ctx context.Context, userID uuid.UUID) (int, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Follower{}).
		Where("following_id = ?", userID).
		Count(&count).Error
	return int(count), err
}

func (r *followerRepository) CountFollowing(ctx context.Context, userID uuid.UUID) (int, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Follower{}).
		Where("follower_id = ?", userID).
		Count(&count).Error
	return int(count), err
}

func (r *followerRepository) ListMutualFollowers(ctx context.Context, userID uuid.UUID) ([]models.Follower, error) {
	var mutuals []models.Follower
	err := r.db.WithContext(ctx).Raw(`
		SELECT f1.*
		FROM followers f1
		JOIN followers f2
		ON f1.follower_id = f2.following_id AND f1.following_id = f2.follower_id
		WHERE f1.status = 'accepted' AND f2.status = 'accepted'
		AND f1.following_id = ?
	`, userID).Scan(&mutuals).Error

	return mutuals, err
}

func (r *followerRepository) IsMutualFollower(ctx context.Context, user1, user2 uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Raw(`
		SELECT COUNT(*) FROM followers f1
		JOIN followers f2 ON f1.follower_id = f2.following_id AND f1.following_id = f2.follower_id
		WHERE f1.follower_id = ? AND f1.following_id = ?
	`, user1, user2).Scan(&count).Error

	return count > 0, err
}

func (r *followerRepository) ListUsersWithMutualFriendPrioritized(
	ctx context.Context,
	userID uuid.UUID,
	limit int,
	afterCreatedAt *time.Time,
) ([]models.User, error) {
	query := r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id != ?", userID).
		Where("id NOT IN (?)",
			r.db.Model(&models.Follower{}).
				Select("following_id").
				Where("follower_id = ?", userID),
		)

	if afterCreatedAt != nil {
		query = query.Where("created_at < ?", *afterCreatedAt)
	}

	orderClause := fmt.Sprintf(`
	(
		SELECT COUNT(*) FROM followers f1
		JOIN followers f2 ON f1.follower_id = f2.follower_id
		WHERE f1.following_id = users.id AND f2.following_id = '%s'
		AND f1.status = 'accepted' AND f2.status = 'accepted'
	) DESC, users.created_at DESC
`, userID.String())

	query = query.Order(orderClause).Limit(limit)

	var users []models.User
	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *followerRepository) GetFollowRecord(ctx context.Context, followerID, followingID uuid.UUID) (*models.Follower, error) {
	var follower *models.Follower
	err := r.db.WithContext(ctx).
		Where("follower_id = ? AND following_id = ?", followerID, followingID).
		First(&follower).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return follower, err
}

func (r *followerRepository) CancelFollowRequest(ctx context.Context, followerID, followingID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("follower_id = ? AND following_id = ? AND status = ?", followerID, followingID, "pending").
		Delete(&models.Follower{}).Error
}
