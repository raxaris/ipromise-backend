package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Promise struct {
	gorm.Model
	ID          uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID      uuid.UUID  `gorm:"type:uuid;not null;index"`
	ParentID    *uuid.UUID `gorm:"type:uuid;index"` // NULL, если это основной Promise
	Title       string     `gorm:"type:varchar(255);not null"`
	Description string     `gorm:"type:text"`
	Deadline    time.Time  `gorm:"not null"`
	Status      string     `gorm:"type:varchar(20);not null"`
}

func (p *Promise) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ParentID != nil {
		var parent Promise
		if err := tx.First(&parent, "id = ?", p.ParentID).Error; err != nil {
			return err // Ошибка, если родительского обещания нет
		}
		p.Deadline = parent.Deadline // Наследуем дедлайн от родителя
	}
	return nil
}
