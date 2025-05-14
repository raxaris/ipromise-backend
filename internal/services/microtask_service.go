package services

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/repositories/microtask"
	"github.com/raxaris/ipromise-backend/internal/repositories/promise"
)

type MicrotaskService interface {
	CreateMicrotask(ctx context.Context, userID, promiseID uuid.UUID, title string, order int) error
	UpdateMicrotask(ctx context.Context, userID, microtaskID uuid.UUID, title *string, status *string) error
	DeleteMicrotask(ctx context.Context, userID, microtaskID uuid.UUID) error
	ListMicrotasksByPromiseID(ctx context.Context, viewerID, promiseID uuid.UUID) ([]models.Microtask, error)
	ReorderMicrotasks(ctx context.Context, promiseID uuid.UUID, orders map[uuid.UUID]int) error
}

type microtaskService struct {
	microtaskRepo microtask.MicrotaskRepository
	promiseRepo   promise.PromiseRepository
}

func NewMicrotaskService(microtaskRepo microtask.MicrotaskRepository, promiseRepo promise.PromiseRepository) MicrotaskService {
	return &microtaskService{
		microtaskRepo: microtaskRepo,
		promiseRepo:   promiseRepo,
	}
}

func (s *microtaskService) CreateMicrotask(ctx context.Context, userID, promiseID uuid.UUID, title string, order int) error {
	p, err := s.promiseRepo.GetPromiseByID(ctx, promiseID)
	if err != nil {
		return err
	}
	if p.UserID != userID {
		return errors.New("нет доступа к этому промису")
	}

	title = strings.TrimSpace(title)
	if title == "" {
		return errors.New("название не может быть пустым")
	}

	mt := &models.Microtask{
		ID:             uuid.New(),
		PromiseID:      promiseID,
		Title:          title,
		Status:         "in_progress",
		MicrotaskOrder: order,
	}

	return s.microtaskRepo.CreateMicrotask(ctx, mt)
}

func (s *microtaskService) UpdateMicrotask(ctx context.Context, userID, microtaskID uuid.UUID, title *string, status *string) error {
	mt, err := s.microtaskRepo.GetMicrotaskByID(ctx, microtaskID)
	if err != nil {
		return err
	}

	p, err := s.promiseRepo.GetPromiseByID(ctx, mt.PromiseID)
	if err != nil {
		return err
	}
	if p.UserID != userID {
		return errors.New("нет доступа к микротаске")
	}

	if title != nil {
		trimmed := strings.TrimSpace(*title)
		if trimmed != "" {
			mt.Title = trimmed
		}
	}

	if status != nil {
		switch *status {
		case "in_progress", "completed":
			mt.Status = *status
		default:
			return errors.New("некорректный статус")
		}
	}

	return s.microtaskRepo.UpdateMicrotask(ctx, mt)
}

func (s *microtaskService) DeleteMicrotask(ctx context.Context, userID, microtaskID uuid.UUID) error {
	// 1. Получаем микротаск
	mt, err := s.microtaskRepo.GetMicrotaskByID(ctx, microtaskID)
	if err != nil {
		return err
	}

	// 2. Получаем промис и проверяем владельца
	promise, err := s.promiseRepo.GetPromiseByID(ctx, mt.PromiseID)
	if err != nil {
		return err
	}
	if promise.UserID != userID {
		return errors.New("access denied: not your microtask")
	}

	// 3. Удаляем микротаск
	if err := s.microtaskRepo.DeleteMicrotask(ctx, microtaskID); err != nil {
		return err
	}

	// 4. Получаем оставшиеся микротаски и переупорядочиваем
	microtasks, err := s.microtaskRepo.ListMicrotasksByPromiseID(ctx, mt.PromiseID)
	if err != nil {
		return err
	}

	orders := make(map[uuid.UUID]int)
	for index, m := range microtasks {
		orders[m.ID] = index
	}

	return s.microtaskRepo.ReorderMicrotasks(ctx, mt.PromiseID, orders)
}

func (s *microtaskService) ListMicrotasksByPromiseID(ctx context.Context, viewerID, promiseID uuid.UUID) ([]models.Microtask, error) {
	p, err := s.promiseRepo.GetPromiseByID(ctx, promiseID)
	if err != nil {
		return nil, err
	}

	if p.IsPrivate && p.UserID != viewerID {
		return nil, errors.New("обещание приватное")
	}

	return s.microtaskRepo.ListMicrotasksByPromiseID(ctx, promiseID)
}

func (s *microtaskService) ReorderMicrotasks(ctx context.Context, promiseID uuid.UUID, orders map[uuid.UUID]int) error {
	return s.microtaskRepo.ReorderMicrotasks(ctx, promiseID, orders)
}
