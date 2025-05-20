package dto

import (
	"mime/multipart"
	"time"
)

type CreatePostRequest struct {
	PromiseID   string                  `form:"promise_id"`                      // можно опционально
	MicrotaskID string                  `form:"microtask_id" binding:"required"` // обязательный
	Content     string                  `form:"content" binding:"required"`
	Attachments []*multipart.FileHeader `form:"attachments"` // не binding
}

type CreateReplyRequest struct {
	Content     string                  `form:"content"`
	Attachments []*multipart.FileHeader `form:"attachments"`
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
	IsLikedByMe    bool                          `json:"is_liked_by_me"`
	IsYourFriend   bool                          `json:"is_your_friend"`
	CommentsCount  int64                         `json:"comments"`
	Attachments    []AttachmentResponse          `json:"attachments,omitempty"`
	Replies        []PostWithRepliesTreeResponse `json:"replies"`
	IsPrivate      bool                          `json:"is_private"`
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
	IsLikedByMe    bool                 `json:"is_liked_by_me"`
	IsYourFriend   bool                 `json:"is_your_friend"`
	CommentsCount  int64                `json:"comments"`
	Attachments    []AttachmentResponse `json:"attachments,omitempty"`
	IsPrivate      bool                 `json:"is_private"`
	CreatedAt      time.Time            `json:"created_at"`
}

type AttachmentResponse struct {
	ID       string `json:"id"`
	FileURL  string `json:"file_url"`
	FileType string `json:"file_type"`
}
