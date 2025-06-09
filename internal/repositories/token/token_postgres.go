package token

import (
	"context"
	"time"

	"github.com/raxaris/ipromise-backend/internal/models"
	"gorm.io/gorm"
)

type tokenRepo struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepo{db: db}
}

func (r *tokenRepo) Save(ctx context.Context, token *models.RefreshToken) error {
	return r.db.WithContext(ctx).Create(token).Error
}

func (r *tokenRepo) FindValid(ctx context.Context, tokenStr string) (*models.RefreshToken, error) {
	var token models.RefreshToken
	err := r.db.WithContext(ctx).Where("token = ?", tokenStr).First(&token).Error
	if err != nil {
		return nil, err
	}

	if time.Now().After(token.ExpiresAt) {
		if err := r.Delete(ctx, &token); err != nil {
			return nil, err
		}
		return nil, gorm.ErrRecordNotFound
	}

	return &token, nil
}

func (r *tokenRepo) Delete(ctx context.Context, token *models.RefreshToken) error {
	return r.db.WithContext(ctx).Delete(token).Error
}
