package dto

import (
	"github.com/google/uuid"
	"time"
)

type CreateMicrotaskRequest struct {
	Title        string `json:"title" binding:"required,min=1,max=255"`
	StepsPlanned int    `json:"steps_planned"`
	Order        int    `json:"order"`
}

type ReorderMicrotasksMapRequest struct {
	Orders map[uuid.UUID]int `json:"orders" binding:"required"`
}

type UpdateMicrotaskRequest struct {
	Title  *string `json:"title,omitempty" binding:"omitempty,min=1,max=255"`
	Status *string `json:"status,omitempty" binding:"omitempty,oneof=pending in_progress completed"`
}

type MicrotaskResponse struct {
	ID        uuid.UUID `json:"id"`
	PromiseID uuid.UUID `json:"promise_id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	Order     int       `json:"order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
