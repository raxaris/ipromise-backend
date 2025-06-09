package models

import (
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID      uuid.UUID      `gorm:"type:uuid;not null;index;constraint:OnDelete:CASCADE"`
	PromiseID   uuid.UUID      `gorm:"type:uuid;not null;index;constraint:OnDelete:CASCADE"`
	MicrotaskID uuid.UUID      `gorm:"type:uuid;not null;index;constraint:OnDelete:CASCADE"`
	ParentID    *uuid.UUID     `gorm:"type:uuid;index;constraint:OnDelete:CASCADE"`
	Content     string         `gorm:"type:text"`
	CreatedAt   time.Time      `gorm:"autoCreateTime;index"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
