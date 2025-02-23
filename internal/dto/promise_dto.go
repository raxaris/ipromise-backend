package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreatePromiseRequest struct {
	ParentID    *uuid.UUID `json:"parent_id,omitempty"` // Если null, это основное обещание
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description"`
	Deadline    *time.Time `json:"deadline,omitempty"` // Можно передавать только для основного промиса
	Status      string     `json:"status" binding:"required"`
	IsPrivate   bool       `json:"is_private"`
}

type UpdatePromiseRequest struct {
	Title       *string    `json:"title,omitempty"`
	Description *string    `json:"description,omitempty"`
	Status      *string    `json:"status,omitempty"`
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
