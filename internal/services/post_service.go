package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/repositories/post"
)

type PostService interface {
	CreatePost(ctx context.Context, userID, promiseID, microtaskID uuid.UUID, content string, parentID *uuid.UUID) error
	UpdatePost(ctx context.Context, userID, postID uuid.UUID, content string) error
	DeletePost(ctx context.Context, userID, postID uuid.UUID) error
	GetPostThread(ctx context.Context, postID uuid.UUID) ([]models.Post, error)
	ListRootPosts(ctx context.Context, microtaskID uuid.UUID, limit int, afterCreatedAt *time.Time, afterID *uuid.UUID) ([]models.Post, error)
	ListReplies(ctx context.Context, postID uuid.UUID, limit int, afterCreatedAt *time.Time, afterID *uuid.UUID) ([]models.Post, error)
}

type postService struct {
	postRepo post.PostRepository
}

func NewPostService(postRepo post.PostRepository) PostService {
	return &postService{postRepo: postRepo}
}

func (s *postService) CreatePost(ctx context.Context, userID, promiseID, microtaskID uuid.UUID, content string, parentID *uuid.UUID) error {
	content = strings.TrimSpace(content)
	if len(content) == 0 {
		return errors.New("контент не может быть пустым")
	}
	post := &models.Post{
		ID:          uuid.New(),
		UserID:      userID,
		PromiseID:   promiseID,
		MicrotaskID: microtaskID,
		ParentID:    parentID,
		Content:     content,
	}
	return s.postRepo.CreatePost(ctx, post)
}

func (s *postService) UpdatePost(ctx context.Context, userID, postID uuid.UUID, content string) error {
	post, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		return err
	}
	if post.UserID != userID {
		return errors.New("недостаточно прав для редактирования")
	}
	post.Content = strings.TrimSpace(content)
	return s.postRepo.UpdatePost(ctx, post)
}

func (s *postService) DeletePost(ctx context.Context, userID, postID uuid.UUID) error {
	post, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		return err
	}
	if post.UserID != userID {
		return errors.New("недостаточно прав для удаления")
	}
	return s.postRepo.DeletePost(ctx, postID)
}

func (s *postService) ListRootPosts(ctx context.Context, microtaskID uuid.UUID, limit int, afterCreatedAt *time.Time, afterID *uuid.UUID) ([]models.Post, error) {
	return s.postRepo.ListRootPostsByMicrotaskID(ctx, microtaskID, limit, afterCreatedAt, afterID)
}

func (s *postService) ListReplies(ctx context.Context, postID uuid.UUID, limit int, afterCreatedAt *time.Time, afterID *uuid.UUID) ([]models.Post, error) {
	return s.postRepo.ListRepliesByPostID(ctx, postID, limit, afterCreatedAt, afterID)
}

func (s *postService) GetPostThread(ctx context.Context, postID uuid.UUID) ([]models.Post, error) {
	// Фича на будущее: получить всю цепочку реплаев (рекурсивно или итеративно)
	return nil, errors.New("ещё не реализовано")
}
