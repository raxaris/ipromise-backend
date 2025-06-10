package post

import (
	"context"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"time"
)

type PostRepository interface {
	CreatePost(ctx context.Context, post *models.Post) error
	GetPostByID(ctx context.Context, id uuid.UUID) (*models.Post, error)
	UpdatePost(ctx context.Context, post *models.Post) error
	DeletePost(ctx context.Context, id uuid.UUID) error
	GetAllPosts(ctx context.Context) ([]models.Post, error)
	ListPostsWithRepliesByMicrotaskID(ctx context.Context, microtaskID uuid.UUID, limit int, after *time.Time, afterID *uuid.UUID) (rootPosts []*models.Post, allReplies []*models.Post, err error)
	ListPostsWithRepliesByPromiseID(ctx context.Context, promiseID uuid.UUID, limit int, after *time.Time, afterID *uuid.UUID) (rootPosts []*models.Post, allReplies []*models.Post, err error)
	CountUserPosts(ctx context.Context, userID uuid.UUID) (int, error)
	ListRepliesByPostID(
		ctx context.Context,
		postID uuid.UUID,
		limit int,
		after *time.Time,
		afterID *uuid.UUID,
	) ([]*models.Post, error)

	ListUserPostsWithReplies(
		ctx context.Context,
		username string,
		viewerID uuid.UUID,
		limit int,
		afterCreatedAt *time.Time,
		afterID *uuid.UUID,
	) ([]*models.Post, []*models.Post, error)

	GetPostWithRepliesTree(
		ctx context.Context,
		postID uuid.UUID,
	) (*models.Post, []*models.Post, error)

	ListPublicPosts(
		ctx context.Context,
		limit int,
		afterCreatedAt *time.Time,
		afterID *uuid.UUID,
	) ([]models.Post, error)

	ListFeedPosts(
		ctx context.Context,
		userID uuid.UUID,
		limit int,
		afterCreatedAt *time.Time,
		afterID *uuid.UUID,
	) ([]models.Post, error)

	ListPublicPostsWithReplies(
		ctx context.Context,
		limit int,
		afterCreatedAt *time.Time,
		afterID *uuid.UUID,
	) ([]*models.Post, []*models.Post, error)

	ListFeedPostsWithReplies(
		ctx context.Context,
		viewerID uuid.UUID,
		limit int,
		afterCreatedAt *time.Time,
		afterID *uuid.UUID,
	) ([]*models.Post, []*models.Post, error)

	CountRepliesByPostID(ctx context.Context, postID uuid.UUID) (int64, error)
	CountRootPostsByMicrotaskID(ctx context.Context, microtaskID uuid.UUID) (int64, error)
}
