package services

import (
	"context"
	"errors"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/mappers"
	"github.com/raxaris/ipromise-backend/internal/repositories/follower"
	"github.com/raxaris/ipromise-backend/internal/repositories/microtask"
	"github.com/raxaris/ipromise-backend/internal/repositories/promise"
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

	GetPostByID(ctx context.Context, postID, viewerID uuid.UUID) (*dto.PostWithRepliesTreeResponse, error)
	ListPublicPostsLite(ctx context.Context, limit int, after *time.Time, afterID *uuid.UUID, viewerID uuid.UUID) ([]dto.PostLiteResponse, error)
	ListFeedPostsLite(ctx context.Context, viewerID uuid.UUID, limit int, after *time.Time, afterID *uuid.UUID) ([]dto.PostLiteResponse, error)
	ListPublicPostsTree(ctx context.Context, limit int, after *time.Time, afterID *uuid.UUID, viewerID uuid.UUID) ([]dto.PostWithRepliesTreeResponse, error)
	ListFeedPostsTree(ctx context.Context, viewerID uuid.UUID, limit int, after *time.Time, afterID *uuid.UUID) ([]dto.PostWithRepliesTreeResponse, error)
	ListPostsByMicrotaskID(ctx context.Context, microtaskID, viewerID uuid.UUID, limit int, afterCreatedAt *time.Time, afterID *uuid.UUID) ([]models.Post, error)
	ListPostsByPromiseID(ctx context.Context, promiseID, viewerID uuid.UUID, limit int, afterCreatedAt *time.Time, afterID *uuid.UUID) ([]models.Post, error)
	ListReplies(ctx context.Context, postID, viewerID uuid.UUID, limit int, afterCreatedAt *time.Time, afterID *uuid.UUID) ([]models.Post, error)
	CountReplies(ctx context.Context, postID uuid.UUID) (int64, error)
}

type postService struct {
	postRepo      post.PostRepository
	promiseRepo   promise.PromiseRepository
	microtaskRepo microtask.MicrotaskRepository
	postMapper    *mappers.PostMapper
	followerRepo  follower.FollowerRepository
}

func NewPostService(
	postRepo post.PostRepository,
	microtaskRepo microtask.MicrotaskRepository,
	postMapper *mappers.PostMapper,
	followerRepo follower.FollowerRepository,
	promiseRepo promise.PromiseRepository,
) PostService {
	return &postService{
		postRepo:      postRepo,
		microtaskRepo: microtaskRepo,
		postMapper:    postMapper,
		followerRepo:  followerRepo,
		promiseRepo:   promiseRepo,
	}
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

func (s *postService) GetPostByID(ctx context.Context, postID, viewerID uuid.UUID) (*dto.PostWithRepliesTreeResponse, error) {
	// 1. Получаем пост
	existingPost, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		return nil, err
	}

	// 2. Получаем информацию о промисе
	existingPromise, err := s.promiseRepo.GetPromiseByID(ctx, existingPost.PromiseID)
	if err != nil {
		return nil, err
	}

	// 3. Проверяем, имеет ли пользователь доступ к этому посту
	canView, err := s.canUserViewPromise(ctx, viewerID, existingPromise.UserID, existingPromise.IsPrivate)
	if err != nil {
		return nil, err
	}
	if !canView {
		return nil, errors.New("access denied: you cannot view this post")
	}

	// 4. Получаем дерево поста
	root, replies, err := s.postRepo.GetPostWithRepliesTree(ctx, postID)
	if err != nil {
		return nil, err
	}

	// 5. Собираем DTO
	return s.postMapper.BuildPostTree(ctx, root, replies, viewerID)
}

func (s *postService) ListPostsByPromiseID(
	ctx context.Context,
	promiseID, viewerID uuid.UUID,
	limit int,
	afterCreatedAt *time.Time,
	afterID *uuid.UUID,
) ([]models.Post, error) {
	existingPromise, err := s.promiseRepo.GetPromiseByID(ctx, promiseID)
	if err != nil {
		return nil, err
	}

	canView, err := s.canUserViewPromise(ctx, viewerID, existingPromise.UserID, existingPromise.IsPrivate)
	if err != nil {
		return nil, err
	}
	if !canView {
		return nil, errors.New("access denied: this promise is private")
	}

	return s.postRepo.ListPostsByPromiseID(ctx, promiseID, limit, afterCreatedAt, afterID)
}

func (s *postService) ListPostsByMicrotaskID(
	ctx context.Context,
	microtaskID, viewerID uuid.UUID,
	limit int,
	afterCreatedAt *time.Time,
	afterID *uuid.UUID,
) ([]models.Post, error) {
	existingMicrotask, err := s.microtaskRepo.GetMicrotaskByID(ctx, microtaskID)
	if err != nil {
		return nil, err
	}

	existingPromise, err := s.promiseRepo.GetPromiseByID(ctx, existingMicrotask.PromiseID)
	if err != nil {
		return nil, err
	}

	canView, err := s.canUserViewPromise(ctx, viewerID, existingPromise.UserID, existingPromise.IsPrivate)
	if err != nil {
		return nil, err
	}
	if !canView {
		return nil, errors.New("access denied")
	}

	return s.postRepo.ListPostsByMicrotaskID(ctx, microtaskID, limit, afterCreatedAt, afterID)
}

func (s *postService) ListPublicPostsLite(
	ctx context.Context,
	limit int,
	after *time.Time,
	afterID *uuid.UUID,
	viewerID uuid.UUID,
) ([]dto.PostLiteResponse, error) {
	rootPosts, err := s.postRepo.ListPublicPosts(ctx, limit, after, afterID)
	if err != nil {
		return nil, err
	}

	var result []dto.PostLiteResponse
	for _, existingPost := range rootPosts {
		dtoLite, err := s.postMapper.BuildPostLite(ctx, &existingPost, viewerID)
		if err != nil {
			return nil, err
		}
		result = append(result, *dtoLite)
	}

	return result, nil
}

func (s *postService) ListFeedPostsLite(
	ctx context.Context,
	viewerID uuid.UUID,
	limit int,
	after *time.Time,
	afterID *uuid.UUID,
) ([]dto.PostLiteResponse, error) {
	rootPosts, err := s.postRepo.ListFeedPosts(ctx, viewerID, limit, after, afterID)
	if err != nil {
		return nil, err
	}

	var result []dto.PostLiteResponse
	for _, existingPost := range rootPosts {
		dtoLite, err := s.postMapper.BuildPostLite(ctx, &existingPost, viewerID)
		if err != nil {
			return nil, err
		}
		result = append(result, *dtoLite)
	}

	return result, nil
}

func (s *postService) ListPublicPostsTree(
	ctx context.Context,
	limit int,
	after *time.Time,
	afterID *uuid.UUID,
	viewerID uuid.UUID,
) ([]dto.PostWithRepliesTreeResponse, error) {
	rootPosts, allReplies, err := s.postRepo.ListPublicPostsWithReplies(ctx, limit, after, afterID)
	if err != nil {
		return nil, err
	}

	return s.postMapper.BuildPostTrees(ctx, rootPosts, allReplies, viewerID)
}

func (s *postService) ListFeedPostsTree(
	ctx context.Context,
	viewerID uuid.UUID,
	limit int,
	after *time.Time,
	afterID *uuid.UUID,
) ([]dto.PostWithRepliesTreeResponse, error) {
	rootPosts, allReplies, err := s.postRepo.ListFeedPostsWithReplies(ctx, viewerID, limit, after, afterID)
	if err != nil {
		return nil, err
	}

	return s.postMapper.BuildPostTrees(ctx, rootPosts, allReplies, viewerID)
}

func (s *postService) ListReplies(ctx context.Context, postID, viewerID uuid.UUID, limit int, afterCreatedAt *time.Time, afterID *uuid.UUID) ([]models.Post, error) {
	existingPost, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		return nil, err
	}
	existingPromise, err := s.promiseRepo.GetPromiseByID(ctx, existingPost.PromiseID)
	if err != nil {
		return nil, err
	}
	canView, err := s.canUserViewPromise(ctx, viewerID, existingPromise.UserID, existingPromise.IsPrivate)
	if err != nil {
		return nil, err
	}
	if !canView {
		return nil, errors.New("access denied")
	}
	return s.postRepo.ListRepliesByPostID(ctx, postID, limit, afterCreatedAt, afterID)
}

func (s *postService) CountReplies(ctx context.Context, postID uuid.UUID) (int64, error) {
	return s.postRepo.CountRepliesByPostID(ctx, postID)
}

func (s *postService) canUserViewPromise(ctx context.Context, viewerID, promiseOwnerID uuid.UUID, isPrivate bool) (bool, error) {
	if !isPrivate {
		return true, nil
	}
	if viewerID == promiseOwnerID {
		return true, nil
	}
	return s.followerRepo.IsMutualFollower(ctx, viewerID, promiseOwnerID)
}
