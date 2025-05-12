package services

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/repositories/badge"
	"github.com/raxaris/ipromise-backend/internal/repositories/follower"
	"github.com/raxaris/ipromise-backend/internal/repositories/promise"
	"github.com/raxaris/ipromise-backend/internal/repositories/user"
)

type ProfileService interface {
	GetMyProfile(ctx context.Context, userID uuid.UUID) (*dto.MyProfileResponse, error)
	GetPublicProfile(ctx context.Context, viewerID uuid.UUID, username string) (*dto.PublicProfileResponse, error)
	UpdateProfile(ctx context.Context, userID uuid.UUID, req dto.UpdateProfileRequest) error
}

type profileService struct {
	userRepo     user.UserRepository
	badgeRepo    badge.BadgeRepository
	promiseRepo  promise.PromiseRepository
	followerRepo follower.FollowerRepository
}

func NewProfileService(
	userRepo user.UserRepository,
	badgeRepo badge.BadgeRepository,
	promiseRepo promise.PromiseRepository,
	followerRepo follower.FollowerRepository,
) ProfileService {
	return &profileService{
		userRepo:     userRepo,
		badgeRepo:    badgeRepo,
		promiseRepo:  promiseRepo,
		followerRepo: followerRepo,
	}
}

func (s *profileService) GetMyProfile(ctx context.Context, userID uuid.UUID) (*dto.MyProfileResponse, error) {
	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	promisesCount, err := s.promiseRepo.CountAllPromisesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	followersCount, err := s.followerRepo.CountFollowers(ctx, userID)
	if err != nil {
		return nil, err
	}

	followingCount, err := s.followerRepo.CountFollowing(ctx, userID)
	if err != nil {
		return nil, err
	}

	badges, err := s.badgeRepo.GetUserBadges(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &dto.MyProfileResponse{
		ID:          user.ID.String(),
		Username:    user.Username,
		Email:       user.Email,
		AvatarURL:   user.AvatarURL,
		Role:        user.Role,
		Followers:   followersCount,
		Following:   followingCount,
		Promises:    int(promisesCount),
		BadgesCount: len(badges),
		Bio:         user.Bio,
	}, nil
}

func (s *profileService) GetPublicProfile(ctx context.Context, viewerID uuid.UUID, username string) (*dto.PublicProfileResponse, error) {
	user, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	isFollowing, err := s.followerRepo.IsFollowing(ctx, viewerID, user.ID)
	if err != nil {
		return nil, err
	}

	canViewAll := viewerID == user.ID || isFollowing || user.Role == "admin"

	var promisesCount int64
	if canViewAll {
		promisesCount, err = s.promiseRepo.CountAllPromisesByUserID(ctx, user.ID)
	} else {
		promisesCount, err = s.promiseRepo.CountPublicPromisesByUserID(ctx, user.ID)
	}
	if err != nil {
		return nil, err
	}

	followersCount, err := s.followerRepo.CountFollowers(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	followingCount, err := s.followerRepo.CountFollowing(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	badges, err := s.badgeRepo.GetUserBadges(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return &dto.PublicProfileResponse{
		ID:          user.ID.String(),
		Username:    user.Username,
		AvatarURL:   user.AvatarURL,
		Bio:         user.Bio,
		Followers:   followersCount,
		Following:   followingCount,
		Promises:    int(promisesCount),
		BadgesCount: len(badges),
		IsFollowing: isFollowing,
	}, nil
}

func (s *profileService) UpdateProfile(ctx context.Context, userID uuid.UUID, req dto.UpdateProfileRequest) error {
	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	if req.Username != nil {
		username := strings.TrimSpace(*req.Username)
		if username != "" && username != user.Username {
			exists, err := s.userRepo.IsUsernameExists(ctx, username)
			if err != nil {
				return err
			}
			if exists {
				return errors.New("это имя пользователя уже занято")
			}
			user.Username = username
		}
	}

	if req.AvatarURL != nil {
		avatar := strings.TrimSpace(*req.AvatarURL)
		if avatar != "" {
			user.AvatarURL = avatar
		}
	}

	if req.Bio != nil {
		bio := strings.TrimSpace(*req.Bio)
		if bio != "" {
			user.Bio = bio
		}
	}

	return s.userRepo.UpdateUser(ctx, user)
}
