package models

import (
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey"`
	UserID      uuid.UUID  `gorm:"type:uuid;not null;index"` // индекс для поиска постов пользователя
	PromiseID   uuid.UUID  `gorm:"type:uuid;not null;index"` // индекс, если будем фильтровать по Promise
	MicrotaskID uuid.UUID  `gorm:"type:uuid;not null;index"` // основной индекс для фильтрации
	ParentID    *uuid.UUID `gorm:"type:uuid;index"`          // для дерева комментариев
	Content     string     `gorm:"type:text"`
	CreatedAt   time.Time  `gorm:"index"` // пагинация и сортировка
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
