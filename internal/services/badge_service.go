package services

import (
	"context"
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
}

type badgeService struct {
	badgeRepo badge.BadgeRepository
}

func NewBadgeService(badgeRepo badge.BadgeRepository) BadgeService {
	return &badgeService{badgeRepo: badgeRepo}
}

// ✅ Админ создаёт новый бейдж
func (s *badgeService) CreateBadge(ctx context.Context, badge *models.Badge) error {
	badge.ID = uuid.New()
	badge.CreatedAt = time.Now()
	return s.badgeRepo.CreateBadge(ctx, badge)
}

// ✅ Получение всех бейджей в системе
func (s *badgeService) ListAllBadges(ctx context.Context) ([]dto.BadgeResponse, error) {
	badges, err := s.badgeRepo.GetAllBadges(ctx)
	if err != nil {
		return nil, err
	}
	return dto.MapBadgesToDTO(badges), nil
}

// ✅ Назначение бейджа пользователю по коду
func (s *badgeService) AssignBadgeByCode(ctx context.Context, userID uuid.UUID, badgeCode string) error {
	has, err := s.badgeRepo.HasUserBadgeByCode(ctx, userID, badgeCode)
	if err != nil {
		return err
	}
	if has {
		return nil // уже есть
	}

	badge, err := s.badgeRepo.GetBadgeByCode(ctx, badgeCode)
	if err != nil {
		return err
	}

	return s.badgeRepo.AssignBadgeToUser(ctx, userID, badge.ID)
}

// ✅ Получение бейджиков конкретного пользователя
func (s *badgeService) GetUserBadges(ctx context.Context, userID uuid.UUID) ([]dto.BadgeResponse, error) {
	badges, err := s.badgeRepo.GetUserBadges(ctx, userID)
	if err != nil {
		return nil, err
	}
	return dto.MapBadgesToDTO(badges), nil
}
