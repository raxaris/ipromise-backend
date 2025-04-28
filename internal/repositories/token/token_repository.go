package token

import (
	"time"

	"github.com/raxaris/ipromise-backend/internal/models"
	"gorm.io/gorm"
)

type TokenRepository interface {
	Save(token *models.RefreshToken) error
	FindValid(token string) (*models.RefreshToken, error)
	Delete(token *models.RefreshToken) error
}

type tokenRepo struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepo{db: db}
}

func (r *tokenRepo) Save(token *models.RefreshToken) error {
	return r.db.Create(token).Error
}

func (r *tokenRepo) FindValid(tokenStr string) (*models.RefreshToken, error) {
	var token models.RefreshToken
	err := r.db.Where("token = ?", tokenStr).First(&token).Error
	if err != nil {
		return nil, err
	}

	if time.Now().After(token.ExpiresAt) {
		err := r.Delete(&token)
		if err != nil {
			return nil, err
		}
		return nil, gorm.ErrRecordNotFound
	}

	return &token, nil
}

func (r *tokenRepo) Delete(token *models.RefreshToken) error {
	return r.db.Delete(token).Error
}
