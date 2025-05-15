package follower

import (
	"context"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"time"
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
	ListSentFollowRequests(ctx context.Context, userID uuid.UUID) ([]models.Follower, error)
	CountFollowing(ctx context.Context, userID uuid.UUID) (int, error)
	CountFollowers(ctx context.Context, userID uuid.UUID) (int, error)
	ListMutualFollowers(ctx context.Context, userID uuid.UUID) ([]models.Follower, error)
	IsMutualFollower(ctx context.Context, user1, user2 uuid.UUID) (bool, error)
	GetFollowRecord(ctx context.Context, followerID, followingID uuid.UUID) (*models.Follower, error)
	ListUsersWithMutualFriendPrioritized(ctx context.Context, userID uuid.UUID, limit int, afterCreatedAt *time.Time) ([]models.User, error)
}
