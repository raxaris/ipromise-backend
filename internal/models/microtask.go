package models

import (
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
)

type Microtask struct {
	ID             uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	PromiseID      uuid.UUID      `gorm:"type:uuid;not null"` // FK → promises
	Title          string         `gorm:"type:varchar(255);not null"`
	StepsPlanned   int            `gorm:"default:0" json:"steps_planned"`
	Status         string         `gorm:"type:varchar(20);not null"`
	MicrotaskOrder int            `gorm:"column:microtask_order;not null"`
	CreatedAt      time.Time      `gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
