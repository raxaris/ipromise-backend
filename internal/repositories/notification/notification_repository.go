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
	MarkManyAsRead(ctx context.Context, ids []uuid.UUID) error
	Exists(ctx context.Context, userID uuid.UUID, notificationType string, relatedID uuid.UUID) (bool, error)
	DeleteByTypeAndRelatedID(ctx context.Context, userID uuid.UUID, notificationType string, relatedID uuid.UUID) error
}
