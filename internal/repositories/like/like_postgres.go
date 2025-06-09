package like

import (
	"context"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"gorm.io/gorm"
)

type likeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) LikeRepository {
	return &likeRepository{db: db}
}

// ✅ Лайкнуть пост
func (r *likeRepository) LikePost(ctx context.Context, userID, postID uuid.UUID) error {
	like := models.Like{
		UserID: userID,
		PostID: postID,
	}
	return r.db.WithContext(ctx).Create(&like).Error
}

// ✅ Убрать лайк
func (r *likeRepository) UnlikePost(ctx context.Context, userID, postID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("user_id = ? AND post_id = ?", userID, postID).
		Delete(&models.Like{}).Error
}

// ✅ Проверить, лайкал ли юзер пост
func (r *likeRepository) IsPostLikedByUser(ctx context.Context, userID, postID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Like{}).
		Where("user_id = ? AND post_id = ?", userID, postID).
		Count(&count).Error
	return count > 0, err
}

// ✅ Посчитать количество лайков у поста
func (r *likeRepository) CountLikesByPostID(ctx context.Context, postID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Like{}).
		Where("post_id = ?", postID).
		Count(&count).Error
	return count, err
}

// ✅ Посчитать лайки + проверить, лайкал ли юзер (оптимизация для фронта)
func (r *likeRepository) CountLikesAndIsLiked(ctx context.Context, postID, userID uuid.UUID) (int64, bool, error) {
	var count int64
	var liked bool

	// Считаем количество лайков
	err := r.db.WithContext(ctx).
		Model(&models.Like{}).
		Where("post_id = ?", postID).
		Count(&count).Error
	if err != nil {
		return 0, false, err
	}

	// Проверяем, лайкал ли пользователь
	liked, err = r.IsPostLikedByUser(ctx, userID, postID)
	if err != nil {
		return count, false, err
	}

	return count, liked, nil
}
