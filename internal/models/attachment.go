package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Attachment struct {
	ID             uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	PostID         *uuid.UUID     `gorm:"type:uuid;index"`
	UserID         *uuid.UUID     `gorm:"type:uuid;index"`
	AttachmentType string         `gorm:"type:varchar(20)"`
	FileURL        string         `gorm:"type:text;not null"`
	FileType       string         `gorm:"type:varchar(50);not null"`
	CreatedAt      time.Time      `gorm:"autoCreateTime"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
