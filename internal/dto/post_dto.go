package dto

import (
	"github.com/google/uuid"
	"time"
)

type CreatePostRequest struct {
	Content  string     `json:"content" binding:"required,min=1"`
	ParentID *uuid.UUID `json:"parent_id,omitempty"` // если это комментарий
}

type UpdatePostRequest struct {
	Content string `json:"content" binding:"required"`
}

type PostWithRepliesTreeResponse struct {
	ID             string                        `json:"id"`
	Content        string                        `json:"content"`
	AuthorID       string                        `json:"author_id"`
	Username       string                        `json:"username"`
	AvatarURL      string                        `json:"avatar_url,omitempty"`
	PromiseID      string                        `json:"promise_id"`
	PromiseTitle   string                        `json:"promise_title"`
	MicrotaskID    string                        `json:"microtask_id"`
	MicrotaskTitle string                        `json:"microtask_title"`
	LikesCount     int64                         `json:"likes"`
	CommentsCount  int64                         `json:"comments"`
	Attachments    []AttachmentResponse          `json:"attachments,omitempty"`
	Replies        []PostWithRepliesTreeResponse `json:"replies"`
	CreatedAt      time.Time                     `json:"created_at"`
}

type PostLiteResponse struct {
	ID             string               `json:"id"`
	Content        string               `json:"content"`
	AuthorID       string               `json:"author_id"`
	Username       string               `json:"username"`
	AvatarURL      string               `json:"avatar_url,omitempty"`
	PromiseID      string               `json:"promise_id"`
	PromiseTitle   string               `json:"promise_title"`
	MicrotaskID    string               `json:"microtask_id"`
	MicrotaskTitle string               `json:"microtask_title"`
	LikesCount     int64                `json:"likes"`
	CommentsCount  int64                `json:"comments"`
	Attachments    []AttachmentResponse `json:"attachments,omitempty"`
	CreatedAt      time.Time            `json:"created_at"`
}

type AttachmentResponse struct {
	URL      string `json:"url"`
	FileType string `json:"file_type"` // image, pdf, etc.
}
