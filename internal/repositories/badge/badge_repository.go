package badge

import (
	"context"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
)

type BadgeRepository interface {
	CreateBadge(ctx context.Context, badge *models.Badge) error
	GetAllBadges(ctx context.Context) ([]models.Badge, error)
	AssignBadgeToUser(ctx context.Context, userID uuid.UUID, badgeID uuid.UUID) error
	HasUserBadge(ctx context.Context, userID uuid.UUID, badgeID uuid.UUID) (bool, error)
	GetUserBadges(ctx context.Context, userID uuid.UUID) ([]models.Badge, error)
	GetBadgeByCode(ctx context.Context, code string) (*models.Badge, error)
	HasUserBadgeByCode(ctx context.Context, userID uuid.UUID, badgeCode string) (bool, error)
}
