package models

import (
	"time"

	"github.com/google/uuid"
)

type Like struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index:idx_user_post,unique;constraint:OnDelete:CASCADE"`
	PostID    uuid.UUID `gorm:"type:uuid;not null;index:idx_user_post,unique;index;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
