package repositories

import (
	"context"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
)

type FollowerRepository interface {
	RequestFollow(ctx context.Context, followerID, followingID uuid.UUID) error
	AcceptFollowRequest(ctx context.Context, followerID, followingID uuid.UUID) error
	DeclineFollowRequest(ctx context.Context, followerID, followingID uuid.UUID) error
	Unfollow(ctx context.Context, followerID, followingID uuid.UUID) error

	IsFollowing(ctx context.Context, followerID, followingID uuid.UUID) (bool, error)

	ListFollowers(ctx context.Context, userID uuid.UUID) ([]models.Follower, error)
	ListFollowing(ctx context.Context, userID uuid.UUID) ([]models.Follower, error)
	ListPendingFollowRequests(ctx context.Context, userID uuid.UUID) ([]models.Follower, error)
}
