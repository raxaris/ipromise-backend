package models

import (
	"time"

	"github.com/google/uuid"
)

type Badge struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Code        string    `gorm:"size:50;unique;not null"`
	Title       string    `gorm:"size:100;not null"`
	Description string    `gorm:"size:255"`
	IconURL     string    `gorm:"size:255"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

type UserBadge struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index:idx_user_badge,unique"`
	BadgeID   uuid.UUID `gorm:"type:uuid;not null;index:idx_user_badge,unique"`
	AwardedAt time.Time `gorm:"autoCreateTime"`
}
