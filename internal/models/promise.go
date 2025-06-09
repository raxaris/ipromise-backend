package models

import (
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
)

type Promise struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID      uuid.UUID      `gorm:"type:uuid;not null;index;constraint:OnDelete:CASCADE"`
	Title       string         `gorm:"type:varchar(255);not null"`
	Description string         `gorm:"type:text"`
	Deadline    time.Time      `gorm:"not null"`
	IsPrivate   bool           `gorm:"not null;default:false"`
	Status      string         `gorm:"type:varchar(20);not null"`
	Category    string         `gorm:"type:varchar(30)"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
