package models

import (
	"github.com/google/uuid"
	"time"
)

type Notification struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null;index;constraint:OnDelete:CASCADE"`
	Type      string     `gorm:"type:text;not null"`
	Message   string     `gorm:"type:text;not null"`
	IsRead    bool       `gorm:"default:false"`
	RelatedID *uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
}
