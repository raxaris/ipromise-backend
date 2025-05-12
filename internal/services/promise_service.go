package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/repositories/follower"
	"github.com/raxaris/ipromise-backend/internal/repositories/promise"
)

type PromiseService interface {
	CreatePromise(ctx context.Context, userID uuid.UUID, title, description string, deadline time.Time, isPrivate bool) error
	GetPromiseByID(ctx context.Context, viewerID uuid.UUID, promiseID uuid.UUID) (*models.Promise, error)
	UpdatePromise(ctx context.Context, userID uuid.UUID, promiseID uuid.UUID, title, description *string, deadline *time.Time, isPrivate *bool) error
	DeletePromise(ctx context.Context, userID uuid.UUID, promiseID uuid.UUID) error

	ListProfilePromises(ctx context.Context, viewerID uuid.UUID, profileUserID uuid.UUID, limit int, afterCreatedAt *time.Time) ([]models.Promise, error)
	ListFeedPromises(ctx context.Context, viewerID uuid.UUID, limit int, afterCreatedAt *time.Time) ([]models.Promise, error)
	ListPublicPromises(ctx context.Context, limit int, afterCreatedAt *time.Time) ([]models.Promise, error)
}

type promiseService struct {
	promiseRepo  promise.PromiseRepository
	followerRepo follower.FollowerRepository
}

func NewPromiseService(promiseRepo promise.PromiseRepository, followerRepo follower.FollowerRepository) PromiseService {
	return &promiseService{
		promiseRepo:  promiseRepo,
		followerRepo: followerRepo,
	}
}

func (s *promiseService) CreatePromise(ctx context.Context, userID uuid.UUID, title, description string, deadline time.Time, isPrivate bool) error {
	// валидируем
	if err := validatePromiseInput(&title, &description, &deadline); err != nil {
		return err
	}

	newPromise := &models.Promise{
		ID:          uuid.New(),
		UserID:      userID,
		Title:       title,
		Description: description,
		Deadline:    deadline,
		IsPrivate:   isPrivate,
		Status:      "in_progress",
	}

	return s.promiseRepo.CreatePromise(ctx, newPromise)
}

func (s *promiseService) GetPromiseByID(ctx context.Context, viewerID, promiseID uuid.UUID) (*models.Promise, error) {
	existingPromise, err := s.promiseRepo.GetPromiseByID(ctx, promiseID)
	if err != nil {
		return nil, err
	}
	if viewerID == existingPromise.UserID || !existingPromise.IsPrivate {
		return existingPromise, nil
	}
	isFollowing, err := s.followerRepo.IsFollowing(ctx, viewerID, existingPromise.UserID)
	if err != nil {
		return nil, err
	}
	if isFollowing {
		return existingPromise, nil
	}
	return nil, errors.New("обещание недоступно")
}

func (s *promiseService) UpdatePromise(ctx context.Context, userID, promiseID uuid.UUID, title, description *string, deadline *time.Time, isPrivate *bool) error {
	existing, err := s.promiseRepo.GetPromiseByID(ctx, promiseID)
	if err != nil {
		return err
	}
	if err := checkOwnership(userID, existing); err != nil {
		return err
	}

	// Валидация входных данных
	if err := validatePromiseInput(title, description, deadline); err != nil {
		return err
	}

	if title != nil {
		existing.Title = strings.TrimSpace(*title)
	}
	if description != nil {
		existing.Description = strings.TrimSpace(*description)
	}
	if deadline != nil {
		existing.Deadline = *deadline
	}
	if isPrivate != nil {
		existing.IsPrivate = *isPrivate
	}

	return s.promiseRepo.UpdatePromise(ctx, existing)
}

func (s *promiseService) DeletePromise(ctx context.Context, userID, promiseID uuid.UUID) error {
	existingPromise, err := s.promiseRepo.GetPromiseByID(ctx, promiseID)
	if err != nil {
		return err
	}
	if err := checkOwnership(userID, existingPromise); err != nil {
		return err
	}
	return s.promiseRepo.DeletePromise(ctx, promiseID)
}

func (s *promiseService) ListProfilePromises(ctx context.Context, viewerID, profileUserID uuid.UUID, limit int, afterCreatedAt *time.Time) ([]models.Promise, error) {
	if viewerID == profileUserID {
		return s.promiseRepo.ListAllPromisesByUserID(ctx, profileUserID, limit, afterCreatedAt)
	}
	isFollowing, err := s.followerRepo.IsFollowing(ctx, viewerID, profileUserID)
	if err != nil {
		return nil, err
	}
	if isFollowing {
		return s.promiseRepo.ListAllPromisesByUserID(ctx, profileUserID, limit, afterCreatedAt)
	}
	return s.promiseRepo.ListPublicPromisesByUserID(ctx, profileUserID, limit, afterCreatedAt)
}

func (s *promiseService) ListFeedPromises(ctx context.Context, viewerID uuid.UUID, limit int, afterCreatedAt *time.Time) ([]models.Promise, error) {
	following, err := s.followerRepo.ListFollowing(ctx, viewerID)
	if err != nil {
		return nil, err
	}
	var userIDs []uuid.UUID
	for _, follow := range following {
		userIDs = append(userIDs, follow.FollowingID)
	}
	if len(userIDs) == 0 {
		return []models.Promise{}, nil
	}
	return s.promiseRepo.ListFeedPromises(ctx, userIDs, limit, afterCreatedAt)
}

func (s *promiseService) ListPublicPromises(ctx context.Context, limit int, afterCreatedAt *time.Time) ([]models.Promise, error) {
	return s.promiseRepo.ListPublicPromises(ctx, limit, afterCreatedAt)
}

func validatePromiseInput(title, description *string, deadline *time.Time) error {
	if title != nil {
		trimmed := strings.TrimSpace(*title)
		if trimmed == "" {
			return errors.New("название не может быть пустым")
		}
		if len(trimmed) > 100 {
			return errors.New("название слишком длинное (макс 100 символов)")
		}
	}

	if description != nil {
		trimmed := strings.TrimSpace(*description)
		if len(trimmed) > 2000 {
			return errors.New("описание слишком длинное (макс 2000 символов)")
		}
	}

	if deadline != nil && time.Now().After(*deadline) {
		return errors.New("дедлайн не может быть в прошлом")
	}

	return nil
}

func checkOwnership(userID uuid.UUID, promise *models.Promise) error {
	if promise.UserID != userID {
		return errors.New("вы не владелец этого обещания")
	}
	return nil
}
