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

func (r *postRepository) CreatePost(ctx context.Context, post *models.Post) error {
	return r.db.WithContext(ctx).Create(post).Error
}

func (r *postRepository) GetPostByID(ctx context.Context, id uuid.UUID) (*models.Post, error) {
	var post models.Post
	if err := r.db.WithContext(ctx).First(&post, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) GetAllPosts(ctx context.Context) ([]models.Post, error) {
	var posts []models.Post
	err := r.db.WithContext(ctx).Find(&posts).Error
	return posts, err
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

func (r *postRepository) DeletePost(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.Post{}, "id = ?", id).Error
}

func (r *postRepository) CountUserPosts(ctx context.Context, userID uuid.UUID) (int, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Post{}).
		Where("user_id = ?", userID).
		Count(&count).Error

	return int(count), err
}

func (r *postRepository) ListPostsWithRepliesByMicrotaskID(
	ctx context.Context,
	microtaskID uuid.UUID,
	limit int,
	after *time.Time,
	afterID *uuid.UUID,
) ([]*models.Post, []*models.Post, error) {
	var roots []*models.Post
	var replies []*models.Post

	query := r.db.WithContext(ctx).Where("microtask_id = ? AND parent_id IS NULL", microtaskID)
	if after != nil {
		query = query.Where("created_at < ? OR (created_at = ? AND id < ?)", *after, *after, *afterID)
	}
	if limit > 0 {
		query = query.Order("created_at DESC").Limit(limit)
	}
	if err := query.Find(&roots).Error; err != nil {
		return nil, nil, err
	}

	rootIDs := make([]uuid.UUID, 0, len(roots))
	for _, r := range roots {
		rootIDs = append(rootIDs, r.ID)
	}

	if len(rootIDs) > 0 {
		if err := r.db.WithContext(ctx).
			Where("root_id IN ?", rootIDs).
			Find(&replies).Error; err != nil {
			return nil, nil, err
		}
	}

	return roots, replies, nil
}

func (r *postRepository) ListPostsWithRepliesByPromiseID(
	ctx context.Context,
	promiseID uuid.UUID,
	limit int,
	after *time.Time,
	afterID *uuid.UUID,
) ([]*models.Post, []*models.Post, error) {
	var roots []*models.Post
	var replies []*models.Post

	query := r.db.WithContext(ctx).Where("promise_id = ? AND parent_id IS NULL", promiseID)
	if after != nil {
		query = query.Where("created_at < ? OR (created_at = ? AND id < ?)", *after, *after, *afterID)
	}
	if limit > 0 {
		query = query.Order("created_at DESC").Limit(limit)
	}
	if err := query.Find(&roots).Error; err != nil {
		return nil, nil, err
	}

	rootIDs := make([]uuid.UUID, 0, len(roots))
	for _, r := range roots {
		rootIDs = append(rootIDs, r.ID)
	}

	if len(rootIDs) > 0 {
		if err := r.db.WithContext(ctx).
			Where("root_id IN ?", rootIDs).
			Find(&replies).Error; err != nil {
			return nil, nil, err
		}
	}

	return roots, replies, nil
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

func (r *postRepository) ListPublicPostsWithReplies(
	ctx context.Context,
	limit int,
	afterCreatedAt *time.Time,
	afterID *uuid.UUID,
) ([]*models.Post, []*models.Post, error) {
	var rootPosts []*models.Post

	query := r.db.WithContext(ctx).
		Raw(`
			SELECT posts.*
			FROM posts
			JOIN microtasks ON posts.microtask_id = microtasks.id
			JOIN promises ON microtasks.promise_id = promises.id
			WHERE promises.is_private = FALSE AND posts.parent_id IS NULL
			ORDER BY posts.created_at DESC, posts.id DESC
			LIMIT ?
		`, limit)

	if afterCreatedAt != nil && afterID != nil {
		query = r.db.WithContext(ctx).Raw(`
			SELECT posts.*
			FROM posts
			JOIN microtasks ON posts.microtask_id = microtasks.id
			JOIN promises ON microtasks.promise_id = promises.id
			WHERE promises.is_private = FALSE AND posts.parent_id IS NULL
				AND ((posts.created_at < ?) OR (posts.created_at = ? AND posts.id < ?))
			ORDER BY posts.created_at DESC, posts.id DESC
			LIMIT ?
		`, *afterCreatedAt, *afterCreatedAt, *afterID, limit)
	}

	if err := query.Scan(&rootPosts).Error; err != nil {
		return nil, nil, err
	}

	// Если root пустой — возвращаем сразу
	if len(rootPosts) == 0 {
		return []*models.Post{}, []*models.Post{}, nil
	}

	// Собираем все root IDs
	var rootIDs []uuid.UUID
	for _, p := range rootPosts {
		rootIDs = append(rootIDs, p.ID)
	}

	// Получаем все replies по root-ам через CTE
	var replies []*models.Post
	err := r.db.WithContext(ctx).Raw(`
		WITH RECURSIVE reply_tree AS (
			SELECT * FROM posts WHERE parent_id IN ?
			UNION ALL
			SELECT p.* FROM posts p
			INNER JOIN reply_tree rt ON p.parent_id = rt.id
		)
		SELECT * FROM reply_tree ORDER BY created_at ASC
	`, rootIDs).Scan(&replies).Error

	if err != nil {
		return nil, nil, err
	}

	return rootPosts, replies, nil
}

func (r *postRepository) ListFeedPostsWithReplies(
	ctx context.Context,
	viewerID uuid.UUID,
	limit int,
	afterCreatedAt *time.Time,
	afterID *uuid.UUID,
) ([]*models.Post, []*models.Post, error) {
	var rootPosts []*models.Post

	query := r.db.WithContext(ctx).Raw(`
		SELECT posts.*
		FROM posts
		JOIN microtasks ON posts.microtask_id = microtasks.id
		JOIN promises ON microtasks.promise_id = promises.id
		JOIN followers ON promises.user_id = followers.following_id
		WHERE followers.follower_id = ? AND promises.is_private = FALSE AND posts.parent_id IS NULL
		ORDER BY posts.created_at DESC, posts.id DESC
		LIMIT ?
	`, viewerID, limit)

	if afterCreatedAt != nil && afterID != nil {
		query = r.db.WithContext(ctx).Raw(`
			SELECT posts.*
			FROM posts
			JOIN microtasks ON posts.microtask_id = microtasks.id
			JOIN promises ON microtasks.promise_id = promises.id
			JOIN followers ON promises.user_id = followers.following_id
			WHERE followers.follower_id = ? AND promises.is_private = FALSE AND posts.parent_id IS NULL
				AND ((posts.created_at < ?) OR (posts.created_at = ? AND posts.id < ?))
			ORDER BY posts.created_at DESC, posts.id DESC
			LIMIT ?
		`, viewerID, *afterCreatedAt, *afterCreatedAt, *afterID, limit)
	}

	if err := query.Scan(&rootPosts).Error; err != nil {
		return nil, nil, err
	}

	if len(rootPosts) == 0 {
		return []*models.Post{}, []*models.Post{}, nil
	}

	var rootIDs []uuid.UUID
	for _, p := range rootPosts {
		rootIDs = append(rootIDs, p.ID)
	}

	var replies []*models.Post
	err := r.db.WithContext(ctx).Raw(`
		WITH RECURSIVE reply_tree AS (
			SELECT * FROM posts WHERE parent_id IN ?
			UNION ALL
			SELECT p.* FROM posts p
			INNER JOIN reply_tree rt ON p.parent_id = rt.id
		)
		SELECT * FROM reply_tree ORDER BY created_at ASC
	`, rootIDs).Scan(&replies).Error

	if err != nil {
		return nil, nil, err
	}

	return rootPosts, replies, nil
}

func (r *postRepository) ListRepliesByPostID(
	ctx context.Context,
	postID uuid.UUID,
	limit int,
	after *time.Time,
	afterID *uuid.UUID,
) ([]*models.Post, error) {
	var replies []*models.Post

	query := r.db.WithContext(ctx).
		Where("parent_id = ?", postID).
		Order("created_at ASC")

	if after != nil && afterID != nil {
		query = query.Where("created_at > ? OR (created_at = ? AND id > ?)", *after, *after, *afterID)
	}

	if limit > 0 {
		query = query.Limit(limit)
	}

	if err := query.Find(&replies).Error; err != nil {
		return nil, err
	}

	return replies, nil
}

func (r *postRepository) ListUserPostsWithReplies(
	ctx context.Context,
	username string,
	viewerID uuid.UUID, // для проверки приватности
	limit int,
	afterCreatedAt *time.Time,
	afterID *uuid.UUID,
) ([]*models.Post, []*models.Post, error) {
	var rootPosts []*models.Post

	baseQuery := `
		SELECT posts.*
		FROM posts
		JOIN users ON posts.user_id = users.id
		JOIN microtasks ON posts.microtask_id = microtasks.id
		JOIN promises ON microtasks.promise_id = promises.id
		WHERE users.username = ? AND posts.parent_id IS NULL
		AND (promises.is_private = FALSE OR promises.user_id = ?)`

	args := []interface{}{username, viewerID}

	if afterCreatedAt != nil && afterID != nil {
		baseQuery += `
			AND ((posts.created_at < ?) OR (posts.created_at = ? AND posts.id < ?))`
		args = append(args, *afterCreatedAt, *afterCreatedAt, *afterID)
	}

	baseQuery += `
		ORDER BY posts.created_at DESC, posts.id DESC
		LIMIT ?`

	args = append(args, limit)

	err := r.db.WithContext(ctx).Raw(baseQuery, args...).Scan(&rootPosts).Error
	if err != nil {
		return nil, nil, err
	}

	if len(rootPosts) == 0 {
		return []*models.Post{}, []*models.Post{}, nil
	}

	// replies
	var rootIDs []uuid.UUID
	for _, p := range rootPosts {
		rootIDs = append(rootIDs, p.ID)
	}

	var replies []*models.Post
	err = r.db.WithContext(ctx).Raw(`
		WITH RECURSIVE reply_tree AS (
			SELECT * FROM posts WHERE parent_id IN ?
			UNION ALL
			SELECT p.* FROM posts p
			INNER JOIN reply_tree rt ON p.parent_id = rt.id
		)
		SELECT * FROM reply_tree ORDER BY created_at ASC
	`, rootIDs).Scan(&replies).Error

	return rootPosts, replies, err
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

func (r *postRepository) CountRootPostsByMicrotaskID(ctx context.Context, microtaskID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Post{}).
		Where("microtask_id = ? AND parent_id IS NULL", microtaskID).
		Count(&count).Error
	return count, err
}
