package token

import (
	"context"
	"github.com/raxaris/ipromise-backend/internal/models"
)

type TokenRepository interface {
	Save(ctx context.Context, token *models.RefreshToken) error
	FindValid(ctx context.Context, token string) (*models.RefreshToken, error)
	Delete(ctx context.Context, token *models.RefreshToken) error
}
