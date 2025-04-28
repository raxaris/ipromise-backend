package models

import (
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	MicrotaskID uuid.UUID      `gorm:"type:uuid;not null"` // FK → microtasks
	UserID      uuid.UUID      `gorm:"type:uuid;not null"` // FK → users
	ParentID    *uuid.UUID     `gorm:"type:uuid"`          // nullable для replies
	Content     string         `gorm:"type:text;not null"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
