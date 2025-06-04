package notification

import (
	"context"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"gorm.io/gorm"
)

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db: db}
}

func (r *notificationRepository) Create(ctx context.Context, notification *models.Notification) error {
	return r.db.WithContext(ctx).Create(notification).Error
}

func (r *notificationRepository) MarkAsRead(ctx context.Context, notificationID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Model(&models.Notification{}).
		Where("id = ?", notificationID).
		Update("is_read", true).Error
}

func (r *notificationRepository) MarkManyAsRead(ctx context.Context, ids []uuid.UUID) error {
	if len(ids) == 0 {
		return nil
	}

	return r.db.WithContext(ctx).
		Model(&models.Notification{}).
		Where("id IN ?", ids).
		Update("is_read", true).
		Error
}

func (r *notificationRepository) ListUserNotifications(ctx context.Context, userID uuid.UUID) ([]models.Notification, error) {
	var notifications []models.Notification
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&notifications).Error
	return notifications, err
}

func (r *notificationRepository) Exists(ctx context.Context, userID uuid.UUID, notificationType string, relatedID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Notification{}).
		Where("user_id = ? AND type = ? AND related_id = ?", userID, notificationType, relatedID).
		Count(&count).Error
	return count > 0, err
}

func (r *notificationRepository) DeleteByTypeAndRelatedID(ctx context.Context, userID uuid.UUID, notificationType string, relatedID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("user_id = ? AND type = ? AND related_id = ?", userID, notificationType, relatedID).
		Delete(&models.Notification{}).Error
}
