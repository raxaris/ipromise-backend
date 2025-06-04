package notification

import (
	"context"
)

type NotificationCache interface {
	SaveOfflineNotification(ctx context.Context, userID string, payload string) error
	GetOfflineNotifications(ctx context.Context, userID string) ([]string, error)
	ClearOfflineNotifications(ctx context.Context, userID string) error
}
