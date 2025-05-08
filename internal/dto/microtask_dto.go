package dto

import "github.com/google/uuid"

type CreateMicrotaskRequest struct {
	Title string `json:"title" binding:"required,min=1,max=255"`
	Order int    `json:"order"`
}

type ReorderMicrotasksMapRequest struct {
	Orders map[uuid.UUID]int `json:"orders" binding:"required"`
}

type UpdateMicrotaskRequest struct {
	Title  *string `json:"title,omitempty" binding:"omitempty,min=1,max=255"`
	Status *string `json:"status,omitempty" binding:"omitempty,oneof=pending in_progress completed"`
}
