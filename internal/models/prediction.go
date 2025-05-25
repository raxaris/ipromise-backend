package models

import (
	"time"

	"github.com/google/uuid"
)

type Prediction struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	PromiseID   uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`
	SuccessRate float64   `gorm:"not null"`
	Advice      string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
