package token

import "github.com/raxaris/ipromise-backend/internal/models"

type TokenRepository interface {
	Save(token *models.RefreshToken) error
	FindValid(token string) (*models.RefreshToken, error)
	Delete(token *models.RefreshToken) error
}
