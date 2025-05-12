package post

import (
	"context"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"time"
)

type PostRepository interface {
	// CRUD
	CreatePost(ctx context.Context, post *models.Post) error
	GetPostByID(ctx context.Context, id uuid.UUID) (*models.Post, error)
	UpdatePost(ctx context.Context, post *models.Post) error
	DeletePost(ctx context.Context, id uuid.UUID) error

	// Получить root посты для microtask (parent_id IS NULL)
	ListRootPostsByMicrotaskID(
		ctx context.Context,
		microtaskID uuid.UUID,
		limit int,
		afterCreatedAt *time.Time,
		afterID *uuid.UUID,
	) ([]models.Post, error)

	// Получить replies (child posts) для конкретного поста
	ListRepliesByPostID(
		ctx context.Context,
		parentID uuid.UUID,
		limit int,
		afterCreatedAt *time.Time,
		afterID *uuid.UUID,
	) ([]models.Post, error)

	// 🔥 Новое: дерево комментариев (вся глубина)
	GetPostWithRepliesTree(ctx context.Context, postID uuid.UUID) (*models.Post, []*models.Post, error)

	// 🔥 Новое: публичная лента
	ListPublicPosts(ctx context.Context, limit int, afterCreatedAt *time.Time, afterID *uuid.UUID) ([]models.Post, error)

	// 🔥 Новое: лента от подписок
	ListFeedPosts(ctx context.Context, userID uuid.UUID, limit int, afterCreatedAt *time.Time, afterID *uuid.UUID) ([]models.Post, error)

	// 🔥 Новое: все посты по PromiseID (для статистики/ленты)
	ListPostsByPromiseID(ctx context.Context, promiseID uuid.UUID) ([]models.Post, error)

	CountRepliesByPostID(ctx context.Context, postID uuid.UUID) (int64, error)
}
