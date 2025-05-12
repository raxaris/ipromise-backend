package services

import (
	"context"
	"errors"
	"github.com/raxaris/ipromise-backend/internal/dto"
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
	ListRootPosts(ctx context.Context, microtaskID uuid.UUID, limit int, afterCreatedAt *time.Time, afterID *uuid.UUID) ([]models.Post, error)
	ListReplies(ctx context.Context, postID uuid.UUID, limit int, afterCreatedAt *time.Time, afterID *uuid.UUID) ([]models.Post, error)
	ListPublicPosts(ctx context.Context, limit int, after *time.Time, afterID *uuid.UUID) ([]models.Post, error)
	ListFeedPosts(ctx context.Context, userID uuid.UUID, limit int, after *time.Time, afterID *uuid.UUID) ([]models.Post, error)
	ListPostsByPromiseID(ctx context.Context, promiseID uuid.UUID) ([]models.Post, error)
	CountReplies(ctx context.Context, postID uuid.UUID) (int64, error)
	GetPostWithRepliesTree(ctx context.Context, postID uuid.UUID) ([]models.Post, error)
	GetPostByID(ctx context.Context, postID, viewerID uuid.UUID) (*dto.PostWithRepliesTreeResponse, error)
	ListPublicPostsLite(ctx context.Context, limit int, after *time.Time, afterID *uuid.UUID, viewerID uuid.UUID) ([]dto.PostLiteResponse, error)
	ListFeedPostsLite(ctx context.Context, viewerID uuid.UUID, limit int, after *time.Time, afterID *uuid.UUID) ([]dto.PostLiteResponse, error)
	ListPublicPostsTree(ctx context.Context, limit int, after *time.Time, afterID *uuid.UUID, viewerID uuid.UUID) ([]dto.PostWithRepliesTreeResponse, error)
	ListFeedPostsTree(ctx context.Context, viewerID uuid.UUID, limit int, after *time.Time, afterID *uuid.UUID) ([]dto.PostWithRepliesTreeResponse, error)
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
		return errors.New("content cannot be empty")
	}

	if parentID != nil {
		parent, err := s.postRepo.GetPostByID(ctx, *parentID)
		if err != nil {
			return errors.New("specified parent post not found")
		}

		if parent.MicrotaskID != microtaskID {
			return errors.New("parent post refers to another microtask")
		}
	}

	newPost := &models.Post{
		ID:          uuid.New(),
		UserID:      userID,
		PromiseID:   promiseID,
		MicrotaskID: microtaskID,
		ParentID:    parentID,
		Content:     content,
	}

	return s.postRepo.CreatePost(ctx, newPost)
}

func (s *postService) UpdatePost(ctx context.Context, userID, postID uuid.UUID, content string) error {
	if len(content) == 0 {
		return errors.New("content cannot be empty")
	}

	existingPost, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		return err
	}
	if existingPost.UserID != userID {
		return errors.New("not enough permissions to edit")
	}
	existingPost.Content = strings.TrimSpace(content)
	return s.postRepo.UpdatePost(ctx, existingPost)
}

func (s *postService) DeletePost(ctx context.Context, userID, postID uuid.UUID) error {
	existingPost, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		return err
	}
	if existingPost.UserID != userID {
		return errors.New("not enough permissions to delete")
	}
	return s.postRepo.DeletePost(ctx, postID)
}

func (s *postService) ListRootPosts(ctx context.Context, microtaskID uuid.UUID, limit int, afterCreatedAt *time.Time, afterID *uuid.UUID) ([]models.Post, error) {
	return s.postRepo.ListRootPostsByMicrotaskID(ctx, microtaskID, limit, afterCreatedAt, afterID)
}

func (s *postService) ListReplies(ctx context.Context, postID uuid.UUID, limit int, afterCreatedAt *time.Time, afterID *uuid.UUID) ([]models.Post, error) {
	return s.postRepo.ListRepliesByPostID(ctx, postID, limit, afterCreatedAt, afterID)
}

func (s *postService) GetPostWithRepliesTree(ctx context.Context, postID uuid.UUID) ([]models.Post, error) {
	root, replies, err := s.postRepo.GetPostWithRepliesTree(ctx, postID)
	if err != nil {
		return nil, err
	}

	result := []models.Post{*root}
	for _, r := range replies {
		result = append(result, *r)
	}

	return result, nil
}

func (s *postService) GetPostByID(ctx context.Context, postID, viewerID uuid.UUID) (*dto.PostWithRepliesTreeResponse, error) {
	// TODO: implement real logic
	return nil, nil
}

func (s *postService) ListPublicPosts(ctx context.Context, limit int, after *time.Time, afterID *uuid.UUID) ([]models.Post, error) {
	return s.postRepo.ListPublicPosts(ctx, limit, after, afterID)
}

func (s *postService) ListFeedPosts(ctx context.Context, userID uuid.UUID, limit int, after *time.Time, afterID *uuid.UUID) ([]models.Post, error) {
	return s.postRepo.ListFeedPosts(ctx, userID, limit, after, afterID)
}

func (s *postService) ListPostsByPromiseID(ctx context.Context, promiseID uuid.UUID) ([]models.Post, error) {
	return s.postRepo.ListPostsByPromiseID(ctx, promiseID)
}

func (s *postService) CountReplies(ctx context.Context, postID uuid.UUID) (int64, error) {
	return s.postRepo.CountRepliesByPostID(ctx, postID)
}

func (s *postService) ListPublicPostsLite(ctx context.Context, limit int, after *time.Time, afterID *uuid.UUID, viewerID uuid.UUID) ([]dto.PostLiteResponse, error) {
	// TODO: implement real logic
	return []dto.PostLiteResponse{}, nil
}

func (s *postService) ListFeedPostsLite(ctx context.Context, viewerID uuid.UUID, limit int, after *time.Time, afterID *uuid.UUID) ([]dto.PostLiteResponse, error) {
	// TODO: implement real logic
	return []dto.PostLiteResponse{}, nil
}

func (s *postService) ListPublicPostsTree(ctx context.Context, limit int, after *time.Time, afterID *uuid.UUID, viewerID uuid.UUID) ([]dto.PostWithRepliesTreeResponse, error) {
	// TODO: implement real logic
	return []dto.PostWithRepliesTreeResponse{}, nil
}

func (s *postService) ListFeedPostsTree(ctx context.Context, viewerID uuid.UUID, limit int, after *time.Time, afterID *uuid.UUID) ([]dto.PostWithRepliesTreeResponse, error) {
	// TODO: implement real logic
	return []dto.PostWithRepliesTreeResponse{}, nil
}
