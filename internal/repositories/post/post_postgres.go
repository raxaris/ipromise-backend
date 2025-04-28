package post

import (
	"context"
	"github.com/google/uuid"
	"time"

	"github.com/raxaris/ipromise-backend/internal/models"
	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

// ✅ CreatePost — создать новый пост
func (r *postRepository) CreatePost(ctx context.Context, post *models.Post) error {
	return r.db.WithContext(ctx).Create(post).Error
}

// ✅ Получить пост по ID
func (r *postRepository) GetPostByID(ctx context.Context, id uuid.UUID) (*models.Post, error) {
	var post models.Post
	if err := r.db.WithContext(ctx).First(&post, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) UpdatePost(ctx context.Context, post *models.Post) error {
	return r.db.WithContext(ctx).
		Model(&models.Post{}).
		Where("id = ?", post.ID).
		Updates(map[string]interface{}{
			"content":    post.Content,
			"updated_at": gorm.Expr("NOW()"),
		}).Error
}

// ✅ Удалить пост по ID (soft delete, если есть DeletedAt)
func (r *postRepository) DeletePost(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.Post{}, "id = ?", id).Error
}

func (r *postRepository) ListRootPostsByMicrotaskID(
	ctx context.Context,
	microtaskID uuid.UUID,
	limit int,
	afterCreatedAt *time.Time,
	afterID *uuid.UUID,
) ([]models.Post, error) {
	var posts []models.Post

	query := r.db.WithContext(ctx).
		Where("microtask_id = ? AND parent_id IS NULL", microtaskID).
		Order("created_at DESC, id DESC").
		Limit(limit)

	if afterCreatedAt != nil && afterID != nil {
		query = query.Where(
			"(created_at < ?) OR (created_at = ? AND id < ?)",
			*afterCreatedAt, *afterCreatedAt, *afterID,
		)
	}

	err := query.Find(&posts).Error
	return posts, err
}

func (r *postRepository) ListRepliesByPostID(
	ctx context.Context,
	parentID uuid.UUID,
	limit int,
	afterCreatedAt *time.Time,
	afterID *uuid.UUID,
) ([]models.Post, error) {
	var posts []models.Post

	query := r.db.WithContext(ctx).
		Where("parent_id = ?", parentID).
		Order("created_at DESC, id DESC").
		Limit(limit)

	if afterCreatedAt != nil && afterID != nil {
		query = query.Where(
			"(created_at < ?) OR (created_at = ? AND id < ?)",
			*afterCreatedAt, *afterCreatedAt, *afterID,
		)
	}

	err := query.Find(&posts).Error
	return posts, err
}

func (r *postRepository) HasReplies(
	ctx context.Context,
	postID uuid.UUID,
) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Post{}).
		Where("parent_id = ?", postID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
