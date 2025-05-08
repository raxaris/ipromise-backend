package dto

import "github.com/google/uuid"

type CreatePostRequest struct {
	Content  string     `json:"content" binding:"required,min=1"`
	ParentID *uuid.UUID `json:"parent_id,omitempty"` // если это комментарий
}

type UpdatePostRequest struct {
	Content string `json:"content" binding:"required"`
}

type PostResponse struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
}
