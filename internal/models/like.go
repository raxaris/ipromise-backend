package models

import (
	"time"

	"github.com/google/uuid"
)

type Like struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index:idx_user_post,unique"`       // уникальность по паре (user_id, post_id)
	PostID    uuid.UUID `gorm:"type:uuid;not null;index:idx_user_post,unique;index"` // участие в UNIQUE + отдельный индекс только по post_id
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
