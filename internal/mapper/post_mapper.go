package mapper

import (
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/models"
)

func MapPostToTreeResponse(
	post *models.Post,
	replies []*models.Post,
	attachments []dto.AttachmentResponse,
	likesCount int64,
	commentsCount int64,
	isLikedByMe bool,
	username string,
	avatarURL *string,
	promiseTitle string,
	microtaskTitle string,
) dto.PostWithRepliesTreeResponse {
	children := make([]dto.PostWithRepliesTreeResponse, 0)

	// Рекурсивно ищем дочерние элементы
	for _, reply := range replies {
		if reply.ParentID != nil && *reply.ParentID == post.ID {
			// Собираем "внуков"
			children = append(children, MapPostToTreeResponse(
				reply,
				replies,
				nil, // вложения для реплаев пока не передаём
				0,   // лайки для реплаев — можно оптимизировать потом
				0,
				false,
				"", nil, "", "",
			))
		}
	}

	return dto.PostWithRepliesTreeResponse{
		ID:             post.ID.String(),
		Content:        post.Content,
		AuthorID:       post.UserID.String(),
		Username:       username,
		AvatarURL:      avatarURL,
		PromiseID:      post.PromiseID.String(),
		PromiseTitle:   promiseTitle,
		MicrotaskID:    post.MicrotaskID.String(),
		MicrotaskTitle: microtaskTitle,
		LikesCount:     likesCount,
		CommentsCount:  commentsCount,
		Attachments:    attachments,
		Replies:        children,
		CreatedAt:      post.CreatedAt,
	}
}

func filterRepliesByParentID(replies []*models.Post, parentID uuid.UUID) []*models.Post {
	var filtered []*models.Post
	for _, r := range replies {
		if r.ParentID != nil && *r.ParentID == parentID {
			filtered = append(filtered, r)
		}
	}
	return filtered
}
