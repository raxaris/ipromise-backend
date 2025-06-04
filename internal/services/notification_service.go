package services

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	notificationcache "github.com/raxaris/ipromise-backend/internal/cache/notification"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/mappers"
	"github.com/raxaris/ipromise-backend/internal/models"
	notificationrepo "github.com/raxaris/ipromise-backend/internal/repositories/notification"
	"github.com/raxaris/ipromise-backend/internal/ws"
	"log"
	"time"
)

type NotificationService interface {
	SendNotification(ctx context.Context, userID uuid.UUID, notification *models.Notification) error
	ListUserNotifications(ctx context.Context, userID uuid.UUID) ([]dto.NotificationResponse, error)
	MarkManyAsRead(ctx context.Context, ids []uuid.UUID) error
	DeliverOfflineNotifications(userID uuid.UUID)
	Exists(ctx context.Context, userID uuid.UUID, notificationType string, relatedID uuid.UUID) (bool, error)
	DeleteByTypeAndRelatedID(ctx context.Context, userID uuid.UUID, notificationType string, relatedID uuid.UUID) error
}

type notificationService struct {
	repo   notificationrepo.NotificationRepository
	cache  notificationcache.NotificationCache
	hub    *ws.NotificationHub
	mapper *mappers.NotificationMapper
}

func NewNotificationService(
	repo notificationrepo.NotificationRepository,
	cache notificationcache.NotificationCache,
	hub *ws.NotificationHub,
	mapper *mappers.NotificationMapper,
) NotificationService {
	return &notificationService{repo: repo, cache: cache, hub: hub, mapper: mapper}
}

func (s *notificationService) SendNotification(ctx context.Context, userID uuid.UUID, notification *models.Notification) error {
	notification.ID = uuid.New()
	notification.UserID = userID
	notification.CreatedAt = time.Now()

	if err := s.repo.Create(ctx, notification); err != nil {
		return err
	}

	payload := s.mapper.ToDTO(notification)

	if s.hub.IsUserConnected(userID.String()) {
		s.hub.SendToUser(userID.String(), payload)
	} else {
		bytes, err := json.Marshal(payload)
		if err != nil {
			log.Printf("❌ Failed to marshal notification: %v", err)
			return err
		}
		if err := s.cache.SaveOfflineNotification(ctx, userID.String(), string(bytes)); err != nil {
			log.Printf("❌ Failed to save to Redis: %v", err)
		}
	}

	return nil
}

func (s *notificationService) ListUserNotifications(ctx context.Context, userID uuid.UUID) ([]dto.NotificationResponse, error) {
	all, err := s.repo.ListUserNotifications(ctx, userID)
	if err != nil {
		return nil, err
	}
	return s.mapper.ToListDTO(all), nil
}

func (s *notificationService) MarkManyAsRead(ctx context.Context, ids []uuid.UUID) error {
	return s.repo.MarkManyAsRead(ctx, ids)
}

func (s *notificationService) DeliverOfflineNotifications(userID uuid.UUID) {
	ctx := context.Background()

	messages, err := s.cache.GetOfflineNotifications(ctx, userID.String())
	if err != nil {
		log.Printf("❌ Failed to get offline notifications from Redis: %v", err)
		return
	}

	for _, msg := range messages {
		s.hub.SendToUser(userID.String(), json.RawMessage(msg))
	}

	if err := s.cache.ClearOfflineNotifications(ctx, userID.String()); err != nil {
		log.Printf("❌ Failed to clear offline notifications: %v", err)
	}
}

func (s *notificationService) Exists(ctx context.Context, userID uuid.UUID, notificationType string, relatedID uuid.UUID) (bool, error) {
	return s.repo.Exists(ctx, userID, notificationType, relatedID)
}

func (s *notificationService) DeleteByTypeAndRelatedID(ctx context.Context, userID uuid.UUID, notificationType string, relatedID uuid.UUID) error {
	return s.repo.DeleteByTypeAndRelatedID(ctx, userID, notificationType, relatedID)
}
