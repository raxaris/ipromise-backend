package models

import (
	"github.com/google/uuid"
	"time"
)

type Notification struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID `gorm:"not null"`
	Type      string    `gorm:"not null"`
	Message   string    `gorm:"type:text;not null"`
	IsRead    bool      `gorm:"default:false"`
	RelatedID *uuid.UUID
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
