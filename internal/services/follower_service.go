package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/repositories/follower"
	"github.com/raxaris/ipromise-backend/internal/repositories/user"
	"gorm.io/gorm"
	"time"
)

type FollowerService interface {
	RequestFollow(ctx context.Context, followerID, followingID uuid.UUID) error
	AcceptFollowRequest(ctx context.Context, followerID, followingID uuid.UUID) error
	DeclineFollowRequest(ctx context.Context, followerID, followingID uuid.UUID) error
	Unfollow(ctx context.Context, followerID, followingID uuid.UUID) error
	CancelFollowRequest(ctx context.Context, followerID, followingID uuid.UUID) error
	IsFollowing(ctx context.Context, followerID, followingID uuid.UUID) (bool, error)

	ListFriendsByUsername(ctx context.Context, username string) ([]dto.UserLiteResponse, error)
	ListFollowers(ctx context.Context, userID uuid.UUID) ([]dto.UserLiteResponse, error)
	ListFollowing(ctx context.Context, userID uuid.UUID) ([]dto.UserLiteResponse, error)
	ListPendingRequests(ctx context.Context, userID uuid.UUID) ([]dto.UserLiteResponse, error)
	ListSentRequests(ctx context.Context, userID uuid.UUID) ([]dto.UserLiteResponse, error)
	GetRecommendedUsers(ctx context.Context, userID uuid.UUID, limit int, afterCreatedAt *time.Time) ([]dto.UserLiteResponse, error)
}

type followerService struct {
	followerRepo        follower.FollowerRepository
	userRepo            user.UserRepository
	notificationService NotificationService
}

func NewFollowerService(followerRepo follower.FollowerRepository, userRepo user.UserRepository, notificationService NotificationService) FollowerService {
	return &followerService{
		followerRepo:        followerRepo,
		userRepo:            userRepo,
		notificationService: notificationService}
}

func (s *followerService) RequestFollow(ctx context.Context, followerID, followingID uuid.UUID) error {
	if followerID == followingID {
		return errors.New("cannot follow yourself")
	}

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

	if err := s.followerRepo.RequestFollow(ctx, followerID, followingID); err != nil {
		return err
	}

	user, err := s.userRepo.GetUserByID(ctx, followerID)
	if err == nil {
		notification := &models.Notification{
			Type:      "follow_request",
			Message:   fmt.Sprintf("%s wants to follow you", user.Username),
			RelatedID: &followerID,
		}
		_ = s.notificationService.SendNotification(ctx, followingID, notification)
	}

	return nil
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
	if err := s.followerRepo.AcceptFollowRequest(ctx, fromUserID, userID); err != nil {
		return err
	}

	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err == nil {
		notification := &models.Notification{
			Type:      "follow_accepted",
			Message:   fmt.Sprintf("%s accepted your follow request.", user.Username),
			RelatedID: &userID,
		}
		_ = s.notificationService.SendNotification(ctx, fromUserID, notification)
	}

	return nil
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

	if err := s.followerRepo.DeclineFollowRequest(ctx, fromUserID, userID); err != nil {
		return err
	}

	_ = s.notificationService.DeleteByTypeAndRelatedID(ctx, userID, "follow_request", fromUserID)

	return nil
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

	if err := s.followerRepo.Unfollow(ctx, followerID, followingID); err != nil {
		return err
	}

	_ = s.notificationService.DeleteByTypeAndRelatedID(ctx, followingID, "follow_accepted", followerID)

	return nil
}

func (s *followerService) CancelFollowRequest(ctx context.Context, followerID, followingID uuid.UUID) error {
	follow, err := s.followerRepo.GetFollowRecord(ctx, followerID, followingID)
	if err != nil {
		return err
	}
	if follow == nil {
		return errors.New("follow request not found")
	}

	if follow.Status != "pending" {
		return errors.New("the request has already been accepted or declined and cannot be canceled")
	}

	if err := s.followerRepo.CancelFollowRequest(ctx, followerID, followingID); err != nil {
		return err
	}

	_ = s.notificationService.DeleteByTypeAndRelatedID(ctx, followingID, "follow_request", followerID)

	return nil
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

func (s *followerService) ListSentRequests(ctx context.Context, userID uuid.UUID) ([]dto.UserLiteResponse, error) {
	records, err := s.followerRepo.ListSentFollowRequests(ctx, userID)
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

func (s *followerService) GetRecommendedUsers(
	ctx context.Context,
	userID uuid.UUID,
	limit int,
	afterCreatedAt *time.Time,
) ([]dto.UserLiteResponse, error) {

	users, err := s.followerRepo.ListUsersWithMutualFriendPrioritized(ctx, userID, limit, afterCreatedAt)
	if err != nil {
		return nil, err
	}

	result := make([]dto.UserLiteResponse, 0)
	for _, u := range users {
		result = append(result, dto.UserLiteResponse{
			ID:        u.ID.String(),
			Username:  u.Username,
			AvatarURL: u.AvatarURL,
			Bio:       u.Bio,
		})
	}

	return result, nil
}
