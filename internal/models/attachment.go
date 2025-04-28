package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Attachment struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	PostID    uuid.UUID      `gorm:"type:uuid;not null"`        // FK → posts
	FileURL   string         `gorm:"type:text;not null"`        // ссылка на файл (где он лежит)
	FileType  string         `gorm:"type:varchar(50);not null"` // "image/png", "application/pdf", "video/mp4"
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"` // если вдруг хочешь soft delete для вложений
}
