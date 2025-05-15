package services

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/repositories/follower"
	"github.com/raxaris/ipromise-backend/internal/repositories/user"
	"gorm.io/gorm"
)

type FollowerService interface {
	RequestFollow(ctx context.Context, followerID, followingID uuid.UUID) error
	AcceptFollowRequest(ctx context.Context, followerID, followingID uuid.UUID) error
	DeclineFollowRequest(ctx context.Context, followerID, followingID uuid.UUID) error
	Unfollow(ctx context.Context, followerID, followingID uuid.UUID) error
	IsFollowing(ctx context.Context, followerID, followingID uuid.UUID) (bool, error)

	ListFriendsByUsername(ctx context.Context, username string) ([]dto.UserLiteResponse, error)
	ListFollowers(ctx context.Context, userID uuid.UUID) ([]dto.UserLiteResponse, error)
	ListFollowing(ctx context.Context, userID uuid.UUID) ([]dto.UserLiteResponse, error)
	ListPendingRequests(ctx context.Context, userID uuid.UUID) ([]dto.UserLiteResponse, error)
}

type followerService struct {
	followerRepo follower.FollowerRepository
	userRepo     user.UserRepository
}

func NewFollowerService(followerRepo follower.FollowerRepository, userRepo user.UserRepository) FollowerService {
	return &followerService{followerRepo: followerRepo, userRepo: userRepo}
}

func (s *followerService) RequestFollow(ctx context.Context, followerID, followingID uuid.UUID) error {
	if followerID == followingID {
		return errors.New("cannot follow yourself")
	}

	// Проверка: есть ли уже запись вообще
	existing, err := s.followerRepo.GetFollowRecord(ctx, followerID, followingID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if existing != nil {
		switch existing.Status {
		case "pending":
			return errors.New("follow request already sent")
		case "accepted":
			return errors.New("already following this user")
		default:
			return errors.New("follow request already exists")
		}
	}

	return s.followerRepo.RequestFollow(ctx, followerID, followingID)
}

func (s *followerService) AcceptFollowRequest(ctx context.Context, userID, fromUserID uuid.UUID) error {
	requests, err := s.followerRepo.ListPendingFollowRequests(ctx, userID)
	if err != nil {
		return err
	}
	found := false
	for _, req := range requests {
		if req.FollowerID == fromUserID {
			found = true
			break
		}
	}
	if !found {
		return errors.New("no pending request from this user")
	}
	return s.followerRepo.AcceptFollowRequest(ctx, fromUserID, userID)
}

func (s *followerService) DeclineFollowRequest(ctx context.Context, userID, fromUserID uuid.UUID) error {
	requests, err := s.followerRepo.ListPendingFollowRequests(ctx, userID)
	if err != nil {
		return err
	}
	found := false
	for _, req := range requests {
		if req.FollowerID == fromUserID {
			found = true
			break
		}
	}
	if !found {
		return errors.New("no pending request from this user")
	}
	return s.followerRepo.DeclineFollowRequest(ctx, fromUserID, userID)
}

func (s *followerService) Unfollow(ctx context.Context, followerID, followingID uuid.UUID) error {
	if followerID == followingID {
		return errors.New("cannot unfollow yourself")
	}

	isFollowing, err := s.followerRepo.IsFollowing(ctx, followerID, followingID)
	if err != nil {
		return err
	}
	if !isFollowing {
		return errors.New("you are not following this user")
	}

	return s.followerRepo.Unfollow(ctx, followerID, followingID)
}

func (s *followerService) IsFollowing(ctx context.Context, followerID, followingID uuid.UUID) (bool, error) {
	return s.followerRepo.IsFollowing(ctx, followerID, followingID)
}

func (s *followerService) ListFriendsByUsername(ctx context.Context, username string) ([]dto.UserLiteResponse, error) {
	user, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	records, err := s.followerRepo.ListMutualFollowers(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	result := make([]dto.UserLiteResponse, 0)
	for _, f := range records {
		friendID := f.FollowerID
		if f.FollowerID == user.ID {
			friendID = f.FollowingID
		}
		friend, err := s.userRepo.GetUserByID(ctx, friendID)
		if err == nil {
			result = append(result, dto.UserLiteResponse{
				ID:        friend.ID.String(),
				Username:  friend.Username,
				AvatarURL: friend.AvatarURL,
				Bio:       friend.Bio,
			})
		}
	}

	return result, nil
}

func (s *followerService) ListFollowers(ctx context.Context, userID uuid.UUID) ([]dto.UserLiteResponse, error) {
	records, err := s.followerRepo.ListFollowers(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make([]dto.UserLiteResponse, 0)
	for _, f := range records {
		u, err := s.userRepo.GetUserByID(ctx, f.FollowerID)
		if err == nil {
			result = append(result, dto.UserLiteResponse{
				ID:        u.ID.String(),
				Username:  u.Username,
				AvatarURL: u.AvatarURL,
				Bio:       u.Bio,
			})
		}
	}
	return result, nil
}

func (s *followerService) ListFollowing(ctx context.Context, userID uuid.UUID) ([]dto.UserLiteResponse, error) {
	records, err := s.followerRepo.ListFollowing(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make([]dto.UserLiteResponse, 0)
	for _, f := range records {
		u, err := s.userRepo.GetUserByID(ctx, f.FollowingID)
		if err == nil {
			result = append(result, dto.UserLiteResponse{
				ID:        u.ID.String(),
				Username:  u.Username,
				AvatarURL: u.AvatarURL,
				Bio:       u.Bio,
			})
		}
	}
	return result, nil
}

func (s *followerService) ListPendingRequests(ctx context.Context, userID uuid.UUID) ([]dto.UserLiteResponse, error) {
	records, err := s.followerRepo.ListPendingFollowRequests(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make([]dto.UserLiteResponse, 0)
	for _, f := range records {
		u, err := s.userRepo.GetUserByID(ctx, f.FollowerID)
		if err == nil {
			result = append(result, dto.UserLiteResponse{
				ID:        u.ID.String(),
				Username:  u.Username,
				AvatarURL: u.AvatarURL,
				Bio:       u.Bio,
			})
		}
	}
	return result, nil
}
