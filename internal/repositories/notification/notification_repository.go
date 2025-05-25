package notification

import (
	"context"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
)

type NotificationRepository interface {
	Create(ctx context.Context, n *models.Notification) error
	MarkAsRead(ctx context.Context, id uuid.UUID) error
	ListUserNotifications(ctx context.Context, userID uuid.UUID) ([]models.Notification, error)
}
