package services

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/repositories/like"
	"github.com/raxaris/ipromise-backend/internal/repositories/post"
)

type LikeService interface {
	LikePost(ctx context.Context, userID, postID uuid.UUID) error
	UnlikePost(ctx context.Context, userID, postID uuid.UUID) error
	IsPostLikedByUser(ctx context.Context, userID, postID uuid.UUID) (bool, error)
	CountLikesByPostID(ctx context.Context, postID uuid.UUID) (int64, error)
	CountLikesAndIsLiked(ctx context.Context, postID, userID uuid.UUID) (count int64, liked bool, err error)
}

type likeService struct {
	repo     like.LikeRepository
	postRepo post.PostRepository
}

func NewLikeService(repo like.LikeRepository, postRepo post.PostRepository) LikeService {
	return &likeService{
		repo:     repo,
		postRepo: postRepo}
}

func (s *likeService) LikePost(ctx context.Context, userID, postID uuid.UUID) error {
	if err := s.validatePostExists(ctx, postID); err != nil {
		return err
	}

	liked, err := s.repo.IsPostLikedByUser(ctx, userID, postID)
	if err != nil {
		return err
	}
	if liked {
		return errors.New("post already liked")
	}

	return s.repo.LikePost(ctx, userID, postID)
}

func (s *likeService) UnlikePost(ctx context.Context, userID, postID uuid.UUID) error {
	if err := s.validatePostExists(ctx, postID); err != nil {
		return err
	}

	liked, err := s.repo.IsPostLikedByUser(ctx, userID, postID)
	if err != nil {
		return err
	}
	if !liked {
		return errors.New("like not found")
	}

	return s.repo.UnlikePost(ctx, userID, postID)
}

func (s *likeService) IsPostLikedByUser(ctx context.Context, userID, postID uuid.UUID) (bool, error) {
	return s.repo.IsPostLikedByUser(ctx, userID, postID)
}

func (s *likeService) CountLikesByPostID(ctx context.Context, postID uuid.UUID) (int64, error) {
	return s.repo.CountLikesByPostID(ctx, postID)
}

func (s *likeService) CountLikesAndIsLiked(ctx context.Context, postID, userID uuid.UUID) (int64, bool, error) {
	return s.repo.CountLikesAndIsLiked(ctx, postID, userID)
}

func (s *likeService) validatePostExists(ctx context.Context, postID uuid.UUID) error {
	_, err := s.postRepo.GetPostByID(ctx, postID)
	return err
}
