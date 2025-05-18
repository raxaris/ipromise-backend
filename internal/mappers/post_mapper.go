package mappers

import (
	"context"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/repositories/attachment"
	"github.com/raxaris/ipromise-backend/internal/repositories/like"
	"github.com/raxaris/ipromise-backend/internal/repositories/microtask"
	"github.com/raxaris/ipromise-backend/internal/repositories/post"
	"github.com/raxaris/ipromise-backend/internal/repositories/promise"
	"github.com/raxaris/ipromise-backend/internal/repositories/user"
)

type PostMapper struct {
	userRepository       user.UserRepository
	likeRepository       like.LikeRepository
	attachmentRepository attachment.AttachmentRepository
	microtaskRepository  microtask.MicrotaskRepository
	promiseRepository    promise.PromiseRepository
	postRepository       post.PostRepository
}

func NewPostMapper(
	userRepository user.UserRepository,
	likeRepository like.LikeRepository,
	attachmentRepository attachment.AttachmentRepository,
	microtaskRepository microtask.MicrotaskRepository,
	promiseRepository promise.PromiseRepository,
	postRepository post.PostRepository,
) *PostMapper {
	return &PostMapper{
		userRepository:       userRepository,
		likeRepository:       likeRepository,
		attachmentRepository: attachmentRepository,
		microtaskRepository:  microtaskRepository,
		promiseRepository:    promiseRepository,
		postRepository:       postRepository,
	}
}

func (pm *PostMapper) BuildPostLite(
	ctx context.Context,
	post *models.Post,
	viewerID uuid.UUID,
) (*dto.PostLiteResponse, error) {
	userModel, err := pm.userRepository.GetUserByID(ctx, post.UserID)
	if err != nil {
		return nil, err
	}

	likesCount, err := pm.likeRepository.CountLikesByPostID(ctx, post.ID)
	if err != nil {
		return nil, err
	}

	commentsCount, err := pm.postRepository.CountRepliesByPostID(ctx, post.ID)
	if err != nil {
		return nil, err
	}

	isLiked, err := pm.likeRepository.IsPostLikedByUser(ctx, viewerID, post.ID)
	if err != nil {
		return nil, err
	}

	attachments, err := pm.attachmentRepository.ListAttachmentsByPostID(ctx, post.ID)
	if err != nil {
		return nil, err
	}

	attachmentDTOs := make([]dto.AttachmentResponse, 0, len(attachments))
	for _, a := range attachments {
		attachmentDTOs = append(attachmentDTOs, dto.AttachmentResponse{
			ID:       a.ID.String(),
			FileURL:  a.FileURL,
			FileType: a.FileType,
		})
	}

	microtaskModel, err := pm.microtaskRepository.GetMicrotaskByID(ctx, post.MicrotaskID)
	if err != nil {
		return nil, err
	}

	promiseModel, err := pm.promiseRepository.GetPromiseByID(ctx, post.PromiseID)
	if err != nil {
		return nil, err
	}

	return &dto.PostLiteResponse{
		ID:             post.ID.String(),
		Content:        post.Content,
		AuthorID:       userModel.ID.String(),
		Username:       userModel.Username,
		AvatarURL:      userModel.AvatarURL,
		PromiseID:      promiseModel.ID.String(),
		PromiseTitle:   promiseModel.Title,
		MicrotaskID:    microtaskModel.ID.String(),
		MicrotaskTitle: microtaskModel.Title,
		LikesCount:     likesCount,
		IsLikedByMe:    isLiked,
		CommentsCount:  commentsCount,
		Attachments:    attachmentDTOs,
		IsPrivate:      promiseModel.IsPrivate,
		CreatedAt:      post.CreatedAt,
	}, nil
}

func (pm *PostMapper) BuildPostTree(
	ctx context.Context,
	root *models.Post,
	replies []*models.Post,
	viewerID uuid.UUID,
) (*dto.PostWithRepliesTreeResponse, error) {
	// Кешируем все replies по parentID
	childrenMap := make(map[uuid.UUID][]*models.Post)
	for _, reply := range replies {
		if reply.ParentID != nil {
			childrenMap[*reply.ParentID] = append(childrenMap[*reply.ParentID], reply)
		}
	}

	// Вложенная рекурсивная функция
	var build func(*models.Post) (*dto.PostWithRepliesTreeResponse, error)
	build = func(post *models.Post) (*dto.PostWithRepliesTreeResponse, error) {
		dtoLite, err := pm.BuildPostLite(ctx, post, viewerID)
		if err != nil {
			return nil, err
		}

		result := &dto.PostWithRepliesTreeResponse{
			ID:             dtoLite.ID,
			Content:        dtoLite.Content,
			AuthorID:       dtoLite.AuthorID,
			Username:       dtoLite.Username,
			AvatarURL:      dtoLite.AvatarURL,
			PromiseID:      dtoLite.PromiseID,
			PromiseTitle:   dtoLite.PromiseTitle,
			MicrotaskID:    dtoLite.MicrotaskID,
			MicrotaskTitle: dtoLite.MicrotaskTitle,
			LikesCount:     dtoLite.LikesCount,
			IsLikedByMe:    dtoLite.IsLikedByMe,
			CommentsCount:  dtoLite.CommentsCount,
			Attachments:    dtoLite.Attachments,
			IsPrivate:      dtoLite.IsPrivate,
			CreatedAt:      dtoLite.CreatedAt,
		}

		for _, child := range childrenMap[post.ID] {
			childDTO, err := build(child)
			if err != nil {
				return nil, err
			}
			result.Replies = append(result.Replies, *childDTO)
		}

		return result, nil
	}

	return build(root)
}

func (pm *PostMapper) BuildPostTrees(
	ctx context.Context,
	rootPosts []*models.Post,
	allReplies []*models.Post,
	viewerID uuid.UUID,
) ([]dto.PostWithRepliesTreeResponse, error) {

	trees := make([]dto.PostWithRepliesTreeResponse, 0)
	for _, root := range rootPosts {
		tree, err := pm.BuildPostTree(ctx, root, allReplies, viewerID)
		if err != nil {
			return nil, err
		}
		trees = append(trees, *tree)
	}

	return trees, nil
}
