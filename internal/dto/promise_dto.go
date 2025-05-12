package dto

import (
	"time"
)

// ✅ DTO для создания Promise
type CreatePromiseRequest struct {
	Title       string    `json:"title" binding:"required,min=1,max=100"`
	Description string    `json:"description" binding:"max=2000"`
	Deadline    time.Time `json:"deadline" binding:"required"`
	IsPrivate   bool      `json:"is_private"`
}

// ✅ DTO для обновления Promise
type UpdatePromiseRequest struct {
	Title       *string    `json:"title" binding:"omitempty,min=1,max=100"`
	Description *string    `json:"description" binding:"omitempty,max=2000"`
	Deadline    *time.Time `json:"deadline" binding:"omitempty"`
	IsPrivate   *bool      `json:"is_private" binding:"omitempty"`
}

// ✅ Ответ клиенту
type PromiseResponse struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"` // можно добавить позже
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	IsPrivate   bool      `json:"is_private"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
