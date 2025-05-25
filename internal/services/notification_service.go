package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/repositories/notification"
	"time"
)

type NotificationService interface {
	SendNotification(ctx context.Context, userID uuid.UUID, notificationType, message string, relatedID *uuid.UUID) error
	MarkAsRead(ctx context.Context, notificationID uuid.UUID) error
	GetMyNotifications(ctx context.Context, userID uuid.UUID) ([]dto.NotificationResponse, error)
}

type notificationService struct {
	notificationRepo notification.NotificationRepository
}

func NewNotificationService(notificationRepository notification.NotificationRepository) NotificationService {
	return &notificationService{
		notificationRepo: notificationRepository}
}

func (s *notificationService) SendNotification(ctx context.Context, userID uuid.UUID, notificationType, message string, relatedID *uuid.UUID) error {
	notification := &models.Notification{
		ID:        uuid.New(),
		UserID:    userID,
		Type:      notificationType,
		Message:   message,
		RelatedID: relatedID,
		IsRead:    false,
		CreatedAt: time.Now(),
	}
	return s.notificationRepo.Create(ctx, notification)
}

func (s *notificationService) MarkAsRead(ctx context.Context, notificationID uuid.UUID) error {
	return s.notificationRepo.MarkAsRead(ctx, notificationID)
}

func (s *notificationService) GetMyNotifications(ctx context.Context, userID uuid.UUID) ([]dto.NotificationResponse, error) {
	notifications, err := s.notificationRepo.ListUserNotifications(ctx, userID)
	if err != nil {
		return nil, err
	}
	var res []dto.NotificationResponse
	for _, n := range notifications {
		res = append(res, dto.NotificationResponse{
			ID:        n.ID.String(),
			Type:      n.Type,
			Message:   n.Message,
			RelatedID: n.RelatedID.String(),
			IsRead:    n.IsRead,
			CreatedAt: n.CreatedAt,
		})
	}
	return res, nil
}
