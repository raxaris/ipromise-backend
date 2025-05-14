package dto

import (
	"time"
)

type CreatePromiseRequest struct {
	Title       string    `json:"title" binding:"required,min=1,max=100"`
	Description string    `json:"description" binding:"max=2000"`
	Deadline    time.Time `json:"deadline" binding:"required"`
	IsPrivate   bool      `json:"is_private"`
}

type CreatePromiseWithMicrotasksRequest struct {
	Title       string                 `json:"title" binding:"required"`
	Description string                 `json:"description"`
	Deadline    time.Time              `json:"deadline" binding:"required"`
	IsPrivate   bool                   `json:"is_private"`
	Microtasks  []CreateMicrotaskInput `json:"microtasks"`
}

type CreateMicrotaskInput struct {
	Title string `json:"title" binding:"required"`
	Order int    `json:"order"`
}

type UpdatePromiseRequest struct {
	Title       *string    `json:"title" binding:"omitempty,min=1,max=100"`
	Description *string    `json:"description" binding:"omitempty,max=2000"`
	Deadline    *time.Time `json:"deadline" binding:"omitempty"`
	IsPrivate   *bool      `json:"is_private" binding:"omitempty"`
}

type PromiseWithMicrotasksProgressResponse struct {
	ID          string              `json:"id"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Deadline    time.Time           `json:"deadline"`
	IsPrivate   bool                `json:"is_private"`
	Status      string              `json:"status"`
	Microtasks  []MicrotaskProgress `json:"microtasks"`
}

type MicrotaskProgress struct {
	ID              string  `json:"id"`
	Title           string  `json:"title"`
	StepsPlanned    int     `json:"steps_planned"`
	PostsCount      int64   `json:"posts_count"`
	ProgressPercent float64 `json:"completion_ratio"`
	Status          string  `json:"status"`
	Order           int     `json:"order"`
}

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
