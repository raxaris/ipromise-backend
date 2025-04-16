package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreatePromiseRequest struct {
	ParentID    *uuid.UUID `json:"parent_id,omitempty"`
	Title       string     `json:"title" binding:"required,min=5"`
	Description string     `json:"description"`
	Deadline    *time.Time `json:"deadline,omitempty"`
	Status      string     `json:"status" binding:"omitempty,oneof=in_progress completed"` // проверка внутри только если есть
	IsPrivate   bool       `json:"is_private"`
}

type UpdatePromiseRequest struct {
	Title       *string    `json:"title" binding:"omitempty,min=5"`
	Description *string    `json:"description,omitempty"`
	Status      *string    `json:"status" binding:"omitempty,oneof=pending in_progress completed"`
	Deadline    *time.Time `json:"deadline,omitempty"`   // Только для основного обещания
	IsPrivate   *bool      `json:"is_private,omitempty"` // 🔹 Добавлено
}

type PromiseResponse struct {
	ID          uuid.UUID  `json:"id"`
	UserID      uuid.UUID  `json:"user_id"`
	ParentID    *uuid.UUID `json:"parent_id,omitempty"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Deadline    time.Time  `json:"deadline"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
}
