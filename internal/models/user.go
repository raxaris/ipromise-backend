package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Username  string         `gorm:"type:varchar(50);unique;not null" json:"username"`
	Email     string         `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string         `gorm:"type:text;not null" json:"-"`
	Role      string         `gorm:"type:varchar(15);not null;default:'user'" json:"role"`
	AvatarURL string         `gorm:"type:text"`
	Bio       string         `gorm:"type:varchar(160)" json:"bio"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
