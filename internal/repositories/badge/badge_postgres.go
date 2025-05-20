package badge

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"gorm.io/gorm"
)

type badgeRepo struct {
	db *gorm.DB
}

func NewBadgeRepository(db *gorm.DB) BadgeRepository {
	return &badgeRepo{db: db}
}

// ✅ Создание нового бейджика
func (r *badgeRepo) CreateBadge(ctx context.Context, badge *models.Badge) error {
	return r.db.WithContext(ctx).Create(badge).Error
}

// ✅ Получение всех бейджиков
func (r *badgeRepo) GetAllBadges(ctx context.Context) ([]models.Badge, error) {
	var badges []models.Badge
	err := r.db.WithContext(ctx).Find(&badges).Error
	return badges, err
}

// ✅ Назначить бейджик пользователю (однократно)
func (r *badgeRepo) AssignBadgeToUser(ctx context.Context, userID uuid.UUID, badgeID uuid.UUID) error {
	hasBadge, err := r.HasUserBadge(ctx, userID, badgeID)
	if err != nil {
		return err
	}
	if hasBadge {
		return nil // уже есть — не дублируем
	}

	userBadge := &models.UserBadge{
		ID:        uuid.New(),
		UserID:    userID,
		BadgeID:   badgeID,
		AwardedAt: time.Now(),
	}
	return r.db.WithContext(ctx).Create(userBadge).Error
}

// ✅ Проверка: есть ли у юзера этот бейджик
func (r *badgeRepo) HasUserBadge(ctx context.Context, userID uuid.UUID, badgeID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.UserBadge{}).
		Where("user_id = ? AND badge_id = ?", userID, badgeID).
		Count(&count).Error
	return count > 0, err
}

// ✅ Получение всех бейджиков пользователя
func (r *badgeRepo) GetUserBadges(ctx context.Context, userID uuid.UUID) ([]models.Badge, error) {
	var badges []models.Badge
	err := r.db.WithContext(ctx).
		Model(&models.Badge{}).
		Joins("JOIN user_badges ON badges.id = user_badges.badge_id").
		Where("user_badges.user_id = ?", userID).
		Find(&badges).Error
	return badges, err
}

func (r *badgeRepo) GetBadgeByCode(ctx context.Context, code string) (*models.Badge, error) {
	var badge models.Badge
	err := r.db.WithContext(ctx).Where("code = ?", code).First(&badge).Error
	if err != nil {
		return nil, err
	}
	return &badge, nil
}

func (r *badgeRepo) HasUserBadgeByCode(ctx context.Context, userID uuid.UUID, badgeCode string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Table("user_badges").
		Joins("JOIN badges ON badges.id = user_badges.badge_id").
		Where("user_badges.user_id = ? AND badges.code = ?", userID, badgeCode).
		Count(&count).Error
	return count > 0, err
}
