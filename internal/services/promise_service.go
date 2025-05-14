package services

import (
	"context"
	"errors"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/repositories/microtask"
	"github.com/raxaris/ipromise-backend/internal/repositories/post"
	"github.com/raxaris/ipromise-backend/internal/repositories/user"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/repositories/follower"
	"github.com/raxaris/ipromise-backend/internal/repositories/promise"
)

type PromiseService interface {
	CreatePromise(ctx context.Context, userID uuid.UUID, title, description string, deadline time.Time, isPrivate bool) error
	CreatePromiseWithMicrotasks(ctx context.Context, userID uuid.UUID, req *dto.CreatePromiseWithMicrotasksRequest) error
	GetUserPromisesWithMicrotasksProgress(ctx context.Context, viewerID uuid.UUID, username string, limit int, after *time.Time) ([]dto.PromiseWithMicrotasksProgressResponse, error)
	GetPromiseByID(ctx context.Context, viewerID uuid.UUID, promiseID uuid.UUID) (*models.Promise, error)
	UpdatePromise(ctx context.Context, userID uuid.UUID, promiseID uuid.UUID, title, description *string, deadline *time.Time, isPrivate *bool) error
	DeletePromise(ctx context.Context, userID uuid.UUID, promiseID uuid.UUID) error

	ListProfilePromises(ctx context.Context, viewerID uuid.UUID, profileUserID uuid.UUID, limit int, afterCreatedAt *time.Time) ([]models.Promise, error)
	ListFeedPromises(ctx context.Context, viewerID uuid.UUID, limit int, afterCreatedAt *time.Time) ([]models.Promise, error)
	ListPublicPromises(ctx context.Context, limit int, afterCreatedAt *time.Time) ([]models.Promise, error)
}

type promiseService struct {
	promiseRepo   promise.PromiseRepository
	microtaskRepo microtask.MicrotaskRepository
	followerRepo  follower.FollowerRepository
	userRepo      user.UserRepository
	postRepo      post.PostRepository
}

func NewPromiseService(promiseRepo promise.PromiseRepository, microtaskRepo microtask.MicrotaskRepository, followerRepo follower.FollowerRepository, userRepo user.UserRepository, postRepo post.PostRepository) PromiseService {
	return &promiseService{
		promiseRepo:   promiseRepo,
		followerRepo:  followerRepo,
		microtaskRepo: microtaskRepo,
		userRepo:      userRepo,
		postRepo:      postRepo,
	}
}

func (s *promiseService) CreatePromise(ctx context.Context, userID uuid.UUID, title, description string, deadline time.Time, isPrivate bool) error {
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

func (s *promiseService) CreatePromiseWithMicrotasks(ctx context.Context, userID uuid.UUID, req *dto.CreatePromiseWithMicrotasksRequest) error {
	newPromise := &models.Promise{
		ID:          uuid.New(),
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		Deadline:    req.Deadline,
		IsPrivate:   req.IsPrivate,
		Status:      "in_progress",
	}

	if err := s.promiseRepo.CreatePromise(ctx, newPromise); err != nil {
		return err
	}

	if len(req.Microtasks) > 0 {
		var microtasks []models.Microtask
		for i, m := range req.Microtasks {
			microtasks = append(microtasks, models.Microtask{
				ID:             uuid.New(),
				PromiseID:      newPromise.ID,
				Title:          m.Title,
				Status:         m.Status,
				MicrotaskOrder: i,
			})
		}

		if err := s.microtaskRepo.CreateManyMicrotasks(ctx, microtasks); err != nil {
			return err
		}
	}

	return nil
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

func (s *promiseService) GetUserPromisesWithMicrotasksProgress(ctx context.Context, viewerID uuid.UUID, username string, limit int, after *time.Time) ([]dto.PromiseWithMicrotasksProgressResponse, error) {
	existingUser, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	var promises []models.Promise

	if viewerID == existingUser.ID {
		promises, err = s.promiseRepo.ListAllPromisesByUserID(ctx, existingUser.ID, limit, after)
	} else {
		isFollowing, err := s.followerRepo.IsFollowing(ctx, viewerID, existingUser.ID)
		if err != nil {
			return nil, err
		}
		if isFollowing {
			promises, err = s.promiseRepo.ListAllPromisesByUserID(ctx, existingUser.ID, limit, after)
		} else {
			promises, err = s.promiseRepo.ListPublicPromisesByUserID(ctx, existingUser.ID, limit, after)
		}
	}
	if err != nil {
		return nil, err
	}

	var result []dto.PromiseWithMicrotasksProgressResponse

	for _, promise := range promises {
		microtasks, err := s.microtaskRepo.ListMicrotasksByPromiseID(ctx, promise.ID)
		if err != nil {
			return nil, err
		}

		mtResponses := make([]dto.MicrotaskProgress, 0) // ✅ всегда будет []

		for _, mt := range microtasks {
			postCount, err := s.postRepo.CountRootPostsByMicrotaskID(ctx, mt.ID)
			if err != nil {
				return nil, err
			}

			var percent float64 = 0
			if mt.StepsPlanned > 0 {
				percent = float64(postCount) / float64(mt.StepsPlanned) * 100
			}

			mtResponses = append(mtResponses, dto.MicrotaskProgress{
				ID:              mt.ID.String(),
				Title:           mt.Title,
				StepsPlanned:    mt.StepsPlanned,
				PostsCount:      postCount,
				ProgressPercent: percent,
				Status:          mt.Status,
				Order:           mt.MicrotaskOrder,
			})
		}

		result = append(result, dto.PromiseWithMicrotasksProgressResponse{
			ID:          promise.ID.String(),
			Title:       promise.Title,
			Description: promise.Description,
			Deadline:    promise.Deadline,
			IsPrivate:   promise.IsPrivate,
			Status:      promise.Status,
			Microtasks:  mtResponses,
		})
	}

	return result, nil
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
