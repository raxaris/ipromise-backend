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

func (r *postRepository) GetPostWithRepliesTree(
	ctx context.Context, postID uuid.UUID,
) (*models.Post, []*models.Post, error) {
	root := &models.Post{}
	if err := r.db.WithContext(ctx).First(root, "id = ?", postID).Error; err != nil {
		return nil, nil, err
	}

	var allReplies []*models.Post
	err := r.db.WithContext(ctx).
		Raw(`
			WITH RECURSIVE reply_tree AS (
				SELECT * FROM posts WHERE parent_id = ?
				UNION ALL
				SELECT p.* FROM posts p
				INNER JOIN reply_tree rt ON p.parent_id = rt.id
			)
			SELECT * FROM reply_tree ORDER BY created_at ASC;
		`, postID).Scan(&allReplies).Error

	if err != nil {
		return nil, nil, err
	}

	return root, allReplies, nil
}

func (r *postRepository) ListPublicPosts(
	ctx context.Context, limit int, afterCreatedAt *time.Time, afterID *uuid.UUID,
) ([]models.Post, error) {
	var posts []models.Post
	query := r.db.WithContext(ctx).
		Joins("JOIN microtasks ON posts.microtask_id = microtasks.id").
		Joins("JOIN promises ON microtasks.promise_id = promises.id").
		Where("promises.is_private = FALSE").
		Order("posts.created_at DESC, posts.id DESC").
		Limit(limit)

	if afterCreatedAt != nil && afterID != nil {
		query = query.Where(
			"(posts.created_at < ?) OR (posts.created_at = ? AND posts.id < ?)",
			*afterCreatedAt, *afterCreatedAt, *afterID,
		)
	}

	err := query.Find(&posts).Error
	return posts, err
}

func (r *postRepository) ListFeedPosts(
	ctx context.Context, userID uuid.UUID, limit int, afterCreatedAt *time.Time, afterID *uuid.UUID,
) ([]models.Post, error) {
	var posts []models.Post
	query := r.db.WithContext(ctx).
		Joins("JOIN microtasks ON posts.microtask_id = microtasks.id").
		Joins("JOIN promises ON microtasks.promise_id = promises.id").
		Joins("JOIN followers ON promises.user_id = followers.following_id").
		Where("followers.follower_id = ? AND promises.is_private = FALSE", userID).
		Order("posts.created_at DESC, posts.id DESC").
		Limit(limit)

	if afterCreatedAt != nil && afterID != nil {
		query = query.Where(
			"(posts.created_at < ?) OR (posts.created_at = ? AND posts.id < ?)",
			*afterCreatedAt, *afterCreatedAt, *afterID,
		)
	}

	err := query.Find(&posts).Error
	return posts, err
}

func (r *postRepository) ListPostsByPromiseID(
	ctx context.Context, promiseID uuid.UUID,
) ([]models.Post, error) {
	var posts []models.Post
	err := r.db.WithContext(ctx).
		Joins("JOIN microtasks ON posts.microtask_id = microtasks.id").
		Where("microtasks.promise_id = ?", promiseID).
		Find(&posts).Error
	return posts, err
}

func (r *postRepository) CountRepliesByPostID(
	ctx context.Context, postID uuid.UUID,
) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Post{}).
		Where("parent_id = ?", postID).
		Count(&count).Error
	return count, err
}
