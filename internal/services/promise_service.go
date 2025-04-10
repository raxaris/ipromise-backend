package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/repositories"
	"strings"
)

type PromiseService interface {
	Create(userID uuid.UUID, req dto.CreatePromiseRequest) error
	GetByUserID(userID uuid.UUID) ([]models.Promise, error)
	GetPublic() ([]models.Promise, error)
	GetPublicByUserID(userID uuid.UUID) ([]models.Promise, error)
	GetChildren(parentID uuid.UUID) ([]models.Promise, error)
	GetByID(id uuid.UUID) (*models.Promise, error)
	GetAllPromisesForAdmin() ([]models.Promise, error)
	Update(userID uuid.UUID, id string, req dto.UpdatePromiseRequest, isAdmin bool) error
	Delete(id string) error
}

type promiseService struct {
	repo repositories.PromiseRepository
}

func NewPromiseService(repo repositories.PromiseRepository) PromiseService {
	return &promiseService{repo: repo}
}

func (s *promiseService) Create(userID uuid.UUID, req dto.CreatePromiseRequest) error {
	req.Title = strings.TrimSpace(req.Title)
	req.Description = strings.TrimSpace(req.Description)

	if len(req.Title) < 5 {
		return errors.New("заголовок должен содержать минимум 5 символов")
	}

	promise := models.Promise{
		ID:          uuid.New(),
		UserID:      userID,
		ParentID:    req.ParentID,
		Title:       req.Title,
		Description: req.Description,
		IsPrivate:   req.IsPrivate,
	}

	if req.ParentID == nil {
		promise.Status = "pending"
		if req.Deadline == nil {
			return errors.New("основное обещание должно иметь дедлайн")
		}
		promise.Deadline = *req.Deadline
	} else {
		parent, err := s.repo.GetByID(*req.ParentID)
		if err != nil {
			return errors.New("родительское обещание не найдено")
		}
		promise.Deadline = parent.Deadline
		if req.Status != "in_progress" && req.Status != "completed" {
			return errors.New("прогресс должен быть 'in_progress' или 'completed'")
		}
		promise.Status = req.Status
	}

	return s.repo.Create(&promise)
}

func (s *promiseService) GetByUserID(userID uuid.UUID) ([]models.Promise, error) {
	return s.repo.GetByUserID(userID)
}

func (s *promiseService) GetPublic() ([]models.Promise, error) {
	return s.repo.GetAllPublic()
}

func (s *promiseService) GetByID(id uuid.UUID) (*models.Promise, error) {
	return s.repo.GetByID(id)
}

func (s *promiseService) GetPublicByUserID(userID uuid.UUID) ([]models.Promise, error) {
	return s.repo.GetPublicByUserID(userID)
}

func (s *promiseService) GetAllPromisesForAdmin() ([]models.Promise, error) {
	return s.repo.GetAll()
}

func (s *promiseService) GetChildren(parentID uuid.UUID) ([]models.Promise, error) {
	return s.repo.GetChildren(parentID)
}

func (s *promiseService) Update(userID uuid.UUID, id string, req dto.UpdatePromiseRequest, isAdmin bool) error {
	promiseID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("неверный ID")
	}

	existing, err := s.repo.GetByID(promiseID)
	if err != nil {
		return errors.New("обещание не найдено")
	}

	if existing.UserID != userID && !isAdmin {
		return errors.New("нет прав на редактирование")
	}

	if req.Title != nil {
		existing.Title = strings.TrimSpace(*req.Title)
	}
	if req.Description != nil {
		existing.Description = strings.TrimSpace(*req.Description)
	}
	if req.Status != nil {
		existing.Status = *req.Status
	}
	if req.IsPrivate != nil {
		existing.IsPrivate = *req.IsPrivate
	}
	if req.Deadline != nil {
		existing.Deadline = *req.Deadline
	}

	return s.repo.Update(existing)
}

func (s *promiseService) Delete(id string) error {
	promiseID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("неверный ID")
	}
	return s.repo.Delete(promiseID)
}
