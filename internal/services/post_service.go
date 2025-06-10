package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/mappers"
	"github.com/raxaris/ipromise-backend/internal/repositories/follower"
	"github.com/raxaris/ipromise-backend/internal/repositories/microtask"
	"github.com/raxaris/ipromise-backend/internal/repositories/promise"
	"github.com/raxaris/ipromise-backend/internal/repositories/user"
	"io"
	"log"
	"mime/multipart"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/repositories/post"
)

type PostService interface {
	CreatePostWithAttachments(ctx context.Context, userID uuid.UUID, promiseID *uuid.UUID, microtaskID uuid.UUID, content string, attachments []*multipart.FileHeader) error
	CreateReplyWithAttachments(ctx context.Context, userID uuid.UUID, parentID uuid.UUID, content string, attachments []*multipart.FileHeader) error
	UpdatePost(ctx context.Context, userID, postID uuid.UUID, content string) error
	DeletePost(ctx context.Context, userID, postID uuid.UUID) error

	GetPostByID(ctx context.Context, postID, viewerID uuid.UUID) (*dto.PostWithRepliesTreeResponse, error)
	ListPublicPostsLite(ctx context.Context, limit int, after *time.Time, afterID *uuid.UUID, viewerID uuid.UUID) ([]dto.PostLiteResponse, error)
	ListFeedPostsLite(ctx context.Context, viewerID uuid.UUID, limit int, after *time.Time, afterID *uuid.UUID) ([]dto.PostLiteResponse, error)
	ListPublicPostsTree(ctx context.Context, limit int, after *time.Time, afterID *uuid.UUID, viewerID uuid.UUID) ([]dto.PostWithRepliesTreeResponse, error)
	ListFeedPostsTree(ctx context.Context, viewerID uuid.UUID, limit int, after *time.Time, afterID *uuid.UUID) ([]dto.PostWithRepliesTreeResponse, error)
	ListPostsByMicrotaskIDTree(ctx context.Context, microtaskID, viewerID uuid.UUID, limit int, afterCreatedAt *time.Time, afterID *uuid.UUID) ([]dto.PostWithRepliesTreeResponse, error)
	ListPostsByPromiseIDTree(ctx context.Context, promiseID, viewerID uuid.UUID, limit int, afterCreatedAt *time.Time, afterID *uuid.UUID) ([]dto.PostWithRepliesTreeResponse, error)
	ListUserPostsTree(ctx context.Context, username string, viewerID uuid.UUID, limit int, after *time.Time, afterID *uuid.UUID) ([]dto.PostWithRepliesTreeResponse, error)
	ListReplies(ctx context.Context, postID, viewerID uuid.UUID, limit int, afterCreatedAt *time.Time, afterID *uuid.UUID) ([]dto.PostWithRepliesTreeResponse, error)
	CountReplies(ctx context.Context, postID uuid.UUID) (int64, error)
}

type postService struct {
	postRepo            post.PostRepository
	promiseRepo         promise.PromiseRepository
	microtaskRepo       microtask.MicrotaskRepository
	postMapper          *mappers.PostMapper
	followerRepo        follower.FollowerRepository
	userRepo            user.UserRepository
	attachmentService   AttachmentService
	notificationService NotificationService
}

func NewPostService(
	postRepo post.PostRepository,
	microtaskRepo microtask.MicrotaskRepository,
	postMapper *mappers.PostMapper,
	followerRepo follower.FollowerRepository,
	promiseRepo promise.PromiseRepository,
	attachmentService AttachmentService,
	notificationService NotificationService,
	userRepo user.UserRepository,
) PostService {
	return &postService{
		postRepo:            postRepo,
		microtaskRepo:       microtaskRepo,
		postMapper:          postMapper,
		followerRepo:        followerRepo,
		promiseRepo:         promiseRepo,
		attachmentService:   attachmentService,
		notificationService: notificationService,
		userRepo:            userRepo,
	}
}

func (s *postService) CreatePostWithAttachments(
	ctx context.Context,
	userID uuid.UUID,
	promiseID *uuid.UUID,
	microtaskID uuid.UUID,
	content string,
	attachments []*multipart.FileHeader,
) error {
	if promiseID != nil {
		promise, err := s.promiseRepo.GetPromiseByID(ctx, *promiseID)
		if err != nil {
			return err
		}
		if promise.UserID != userID {
			return errors.New("you are not the owner of this promise")
		}
	}

	microtask, err := s.microtaskRepo.GetMicrotaskByID(ctx, microtaskID)
	if err != nil {
		return err
	}
	if promiseID != nil && microtask.PromiseID != *promiseID {
		return errors.New("microtask does not belong to the specified promise")
	}

	promise, err := s.promiseRepo.GetPromiseByID(ctx, microtask.PromiseID)
	if err != nil {
		return err
	}
	if promise.UserID != userID {
		return errors.New("you are not the owner of the microtask")
	}

	post := &models.Post{
		ID:          uuid.New(),
		UserID:      userID,
		PromiseID:   microtask.PromiseID,
		MicrotaskID: microtaskID,
		Content:     content,
	}

	if err := s.postRepo.CreatePost(ctx, post); err != nil {
		return err
	}

	s.uploadPostAttachments(ctx, post.ID, attachments)
	return nil
}

func (s *postService) CreateReplyWithAttachments(
	ctx context.Context,
	userID uuid.UUID,
	parentID uuid.UUID,
	content string,
	attachments []*multipart.FileHeader,
) error {
	parent, err := s.postRepo.GetPostByID(ctx, parentID)
	if err != nil {
		return err
	}

	post := &models.Post{
		ID:          uuid.New(),
		UserID:      userID,
		PromiseID:   parent.PromiseID,
		MicrotaskID: parent.MicrotaskID,
		ParentID:    &parent.ID,
		Content:     content,
	}

	if err := s.postRepo.CreatePost(ctx, post); err != nil {
		return err
	}

	s.uploadPostAttachments(ctx, post.ID, attachments)

	if parent.UserID != userID {
		user, err := s.userRepo.GetUserByID(ctx, userID)
		if err == nil {
			var notificationType, message string

			if parent.ParentID == nil {
				notificationType = "post_reply"
				message = fmt.Sprintf("💬 %s replied to your post.", user.Username)
			} else {
				notificationType = "comment_reply"
				message = fmt.Sprintf("💬 %s replied to your comment.", user.Username)
			}

			notification := &models.Notification{
				Type:      notificationType,
				Message:   message,
				RelatedID: &post.ID,
			}

			_ = s.notificationService.SendNotification(ctx, parent.UserID, notification)
		}
	}

	return nil
}

func (s *postService) uploadPostAttachments(
	ctx context.Context,
	postID uuid.UUID,
	files []*multipart.FileHeader,
) {
	if len(files) == 0 {
		return
	}

	for _, file := range files {
		f, err := file.Open()
		if err != nil {
			continue
		}

		data, err := io.ReadAll(f)
		_ = f.Close()
		if err != nil {
			continue
		}

		if _, err := s.attachmentService.UploadAttachmentToPost(ctx,
			postID,
			data,
			file.Filename,
			file.Header.Get("Content-Type")); err != nil {
			log.Printf("failed to upload attachment: %v", err)
		}
	}
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

	_ = s.notificationService.DeleteByTypeAndRelatedID(ctx, existingPost.UserID, "post_reply", postID)
	_ = s.notificationService.DeleteByTypeAndRelatedID(ctx, existingPost.UserID, "comment_reply", postID)

	return s.postRepo.DeletePost(ctx, postID)
}

func (s *postService) GetPostByID(ctx context.Context, postID, viewerID uuid.UUID) (*dto.PostWithRepliesTreeResponse, error) {
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
		return nil, errors.New("access denied: you cannot view this post")
	}

	root, replies, err := s.postRepo.GetPostWithRepliesTree(ctx, postID)
	if err != nil {
		return nil, err
	}

	return s.postMapper.BuildPostTree(ctx, root, replies, viewerID)
}

func (s *postService) ListPostsByMicrotaskIDTree(
	ctx context.Context,
	microtaskID, viewerID uuid.UUID,
	limit int,
	after *time.Time,
	afterID *uuid.UUID,
) ([]dto.PostWithRepliesTreeResponse, error) {
	microtask, err := s.microtaskRepo.GetMicrotaskByID(ctx, microtaskID)
	if err != nil {
		return nil, err
	}

	promise, err := s.promiseRepo.GetPromiseByID(ctx, microtask.PromiseID)
	if err != nil {
		return nil, err
	}

	canView, err := s.canUserViewPromise(ctx, viewerID, promise.UserID, promise.IsPrivate)
	if err != nil {
		return nil, err
	}
	if !canView {
		return nil, errors.New("access denied")
	}

	roots, replies, err := s.postRepo.ListPostsWithRepliesByMicrotaskID(ctx, microtaskID, limit, after, afterID)
	if err != nil {
		return nil, err
	}

	return s.postMapper.BuildPostTrees(ctx, roots, replies, viewerID)
}

func (s *postService) ListPostsByPromiseIDTree(
	ctx context.Context,
	promiseID, viewerID uuid.UUID,
	limit int,
	after *time.Time,
	afterID *uuid.UUID,
) ([]dto.PostWithRepliesTreeResponse, error) {
	promise, err := s.promiseRepo.GetPromiseByID(ctx, promiseID)
	if err != nil {
		return nil, err
	}

	canView, err := s.canUserViewPromise(ctx, viewerID, promise.UserID, promise.IsPrivate)
	if err != nil {
		return nil, err
	}
	if !canView {
		return nil, errors.New("access denied")
	}

	roots, replies, err := s.postRepo.ListPostsWithRepliesByPromiseID(ctx, promiseID, limit, after, afterID)
	if err != nil {
		return nil, err
	}

	return s.postMapper.BuildPostTrees(ctx, roots, replies, viewerID)
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

func (s *postService) ListUserPostsTree(
	ctx context.Context,
	username string,
	viewerID uuid.UUID,
	limit int,
	after *time.Time,
	afterID *uuid.UUID,
) ([]dto.PostWithRepliesTreeResponse, error) {
	rootPosts, allReplies, err := s.postRepo.ListUserPostsWithReplies(ctx, username, viewerID, limit, after, afterID)
	if err != nil {
		return nil, err
	}

	return s.postMapper.BuildPostTrees(ctx, rootPosts, allReplies, viewerID)
}

func (s *postService) ListReplies(
	ctx context.Context,
	postID, viewerID uuid.UUID,
	limit int,
	afterCreatedAt *time.Time,
	afterID *uuid.UUID,
) ([]dto.PostWithRepliesTreeResponse, error) {
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

	replies, err := s.postRepo.ListRepliesByPostID(ctx, postID, limit, afterCreatedAt, afterID)
	if err != nil {
		return nil, err
	}

	return s.postMapper.BuildPostTrees(ctx, nil, replies, viewerID)
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
