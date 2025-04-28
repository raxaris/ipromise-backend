package models

import (
	"time"

	"github.com/google/uuid"
)

type Follower struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FollowerID  uuid.UUID `gorm:"type:uuid;not null;index:idx_follower_following,unique"`
	FollowingID uuid.UUID `gorm:"type:uuid;not null;index:idx_follower_following,unique"`
	Status      string    `gorm:"type:varchar(20);not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
