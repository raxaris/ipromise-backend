package services

import (
	"context"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/repositories/microtask"
	"github.com/raxaris/ipromise-backend/internal/repositories/post"
	"github.com/raxaris/ipromise-backend/internal/repositories/promise"
	"github.com/raxaris/ipromise-backend/internal/repositories/user"
)

type AdminService interface {
	ListAllUsers(ctx context.Context) ([]models.User, error)
	ListAllPosts(ctx context.Context) ([]models.Post, error)
	ListAllMicrotasks(ctx context.Context) ([]models.Microtask, error)
	ListAllPromises(ctx context.Context) ([]models.Promise, error)
}

type adminService struct {
	userRepo      user.UserRepository
	postRepo      post.PostRepository
	microtaskRepo microtask.MicrotaskRepository
	promiseRepo   promise.PromiseRepository
}

func NewAdminService(
	userRepo user.UserRepository,
	postRepo post.PostRepository,
	microtaskRepo microtask.MicrotaskRepository,
	promiseRepo promise.PromiseRepository,
) AdminService {
	return &adminService{
		userRepo:      userRepo,
		postRepo:      postRepo,
		microtaskRepo: microtaskRepo,
		promiseRepo:   promiseRepo,
	}
}

func (s *adminService) ListAllUsers(ctx context.Context) ([]models.User, error) {
	return s.userRepo.GetAllUsers(ctx)
}

func (s *adminService) ListAllPosts(ctx context.Context) ([]models.Post, error) {
	return s.postRepo.GetAllPosts(ctx)
}

func (s *adminService) ListAllMicrotasks(ctx context.Context) ([]models.Microtask, error) {
	return s.microtaskRepo.GetAllMicrotasks(ctx)
}

func (s *adminService) ListAllPromises(ctx context.Context) ([]models.Promise, error) {
	return s.promiseRepo.GetAllPromises(ctx)
}
