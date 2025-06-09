package like

import (
	"context"
	"github.com/google/uuid"
)

type LikeRepository interface {
	LikePost(ctx context.Context, userID, postID uuid.UUID) error
	UnlikePost(ctx context.Context, userID, postID uuid.UUID) error
	IsPostLikedByUser(ctx context.Context, userID, postID uuid.UUID) (bool, error)
	CountLikesByPostID(ctx context.Context, postID uuid.UUID) (int64, error)
	CountLikesAndIsLiked(ctx context.Context, postID, userID uuid.UUID) (count int64, liked bool, err error)
}
