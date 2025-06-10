package services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/repositories/badge"
)

type BadgeService interface {
	CreateBadge(ctx context.Context, badge *models.Badge) error
	ListAllBadges(ctx context.Context) ([]dto.BadgeResponse, error)
	AssignBadgeByCode(ctx context.Context, userID uuid.UUID, badgeCode string) error
	GetUserBadges(ctx context.Context, userID uuid.UUID) ([]dto.BadgeResponse, error)
	HasUserBadgeByCode(ctx context.Context, userID uuid.UUID, badgeCode string) (bool, error)
}

type badgeService struct {
	badgeRepo           badge.BadgeRepository
	notificationService NotificationService
}

func NewBadgeService(badgeRepo badge.BadgeRepository, notificationService NotificationService) BadgeService {
	return &badgeService{
		badgeRepo:           badgeRepo,
		notificationService: notificationService,
	}
}

func (s *badgeService) CreateBadge(ctx context.Context, badge *models.Badge) error {
	badge.ID = uuid.New()
	badge.CreatedAt = time.Now()
	return s.badgeRepo.CreateBadge(ctx, badge)
}

func (s *badgeService) ListAllBadges(ctx context.Context) ([]dto.BadgeResponse, error) {
	badges, err := s.badgeRepo.GetAllBadges(ctx)
	if err != nil {
		return nil, err
	}
	return dto.MapBadgesToDTO(badges), nil
}

func (s *badgeService) AssignBadgeByCode(ctx context.Context, userID uuid.UUID, badgeCode string) error {
	has, err := s.badgeRepo.HasUserBadgeByCode(ctx, userID, badgeCode)
	if err != nil {
		return err
	}
	if has {
		return nil
	}

	badge, err := s.badgeRepo.GetBadgeByCode(ctx, badgeCode)
	if err != nil {
		return err
	}

	if err := s.badgeRepo.AssignBadgeToUser(ctx, userID, badge.ID); err != nil {
		return err
	}

	notification := &models.Notification{
		Type:      "badge_assigned",
		Message:   fmt.Sprintf("🏅 You’ve earned the “%s” badge!", badge.Title),
		RelatedID: &badge.ID,
	}
	_ = s.notificationService.SendNotification(ctx, userID, notification)

	return nil
}

func (s *badgeService) GetUserBadges(ctx context.Context, userID uuid.UUID) ([]dto.BadgeResponse, error) {
	badges, err := s.badgeRepo.GetUserBadges(ctx, userID)
	if err != nil {
		return nil, err
	}
	return dto.MapBadgesToDTO(badges), nil
}

func (s *badgeService) HasUserBadgeByCode(ctx context.Context, userID uuid.UUID, badgeCode string) (bool, error) {
	return s.badgeRepo.HasUserBadgeByCode(ctx, userID, badgeCode)
}
