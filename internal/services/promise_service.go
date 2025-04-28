package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/constants"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/repositories"
	"strings"
)

type PromiseService interface {
	Create(userID uuid.UUID, req dto.CreatePromiseRequest) error
	GetByUserID(userID uuid.UUID) ([]models.PromiseV1, error)
	GetPublic() ([]models.PromiseV1, error)
	GetPublicByUserID(userID uuid.UUID) ([]models.PromiseV1, error)
	GetChildren(parentID uuid.UUID) ([]models.PromiseV1, error)
	GetByID(id uuid.UUID) (*models.PromiseV1, error)
	GetAllPromisesForAdmin() ([]models.PromiseV1, error)
	Update(userID uuid.UUID, id string, req dto.UpdatePromiseRequest, isAdmin bool) error
	Delete(userID uuid.UUID, promiseID uuid.UUID, isAdmin bool) error
}

type promiseService struct {
	repo     repositories.PromiseRepositoryV1
	userRepo repositories.UserRepository
}

func NewPromiseService(repo repositories.PromiseRepositoryV1, userRepo repositories.UserRepository) PromiseService {
	return &promiseService{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (s *promiseService) Create(userID uuid.UUID, req dto.CreatePromiseRequest) error {
	req.Title = strings.TrimSpace(req.Title)
	req.Description = strings.TrimSpace(req.Description)

	if req.ParentID == nil {
		return s.createMainPromise(userID, req)
	} else {
		return s.createProgress(userID, req)
	}
}

func (s *promiseService) createMainPromise(userID uuid.UUID, req dto.CreatePromiseRequest) error {
	if req.Deadline == nil {
		return errors.New("основное обещание должно иметь дедлайн")
	}

	promise := &models.PromiseV1{
		ID:          uuid.New(),
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		Deadline:    *req.Deadline,
		IsPrivate:   req.IsPrivate,
		Status:      constants.StatusPending,
	}

	return s.repo.Create(promise)
}

func (s *promiseService) createProgress(userID uuid.UUID, req dto.CreatePromiseRequest) error {
	parent, err := s.repo.GetByID(*req.ParentID)
	if err != nil {
		return errors.New("родительское обещание не найдено")
	}

	hasChild, err := s.repo.HasChild(parent.ID)
	if err != nil {
		return err
	}
	if hasChild {
		return errors.New("у этого обещания уже есть прогресс")
	}

	if req.Status != constants.StatusInProgress && req.Status != constants.StatusCompleted {
		return errors.New("прогресс должен быть in_progress или completed")
	}

	promise := &models.PromiseV1{
		ID:          uuid.New(),
		UserID:      userID,
		ParentID:    req.ParentID,
		Title:       req.Title,
		Description: req.Description,
		Deadline:    parent.Deadline,
		IsPrivate:   parent.IsPrivate || req.IsPrivate,
		Status:      req.Status,
	}

	if err := s.repo.Create(promise); err != nil {
		return err
	}

	// Автоматическое обновление родительской цепочки до completed
	if req.Status == constants.StatusCompleted {
		currentID := req.ParentID
		for currentID != nil {
			parent, err := s.repo.GetByID(*currentID)
			if err != nil {
				break
			}
			parent.Status = constants.StatusCompleted
			if err := s.repo.Update(parent); err != nil {
				break
			}
			currentID = parent.ParentID
		}
	}

	return nil
}

func (s *promiseService) GetByUserID(userID uuid.UUID) ([]models.PromiseV1, error) {
	return s.repo.GetByUserID(userID)
}

func (s *promiseService) GetPublic() ([]models.PromiseV1, error) {
	return s.repo.GetAllPublic()
}

func (s *promiseService) GetByID(id uuid.UUID) (*models.PromiseV1, error) {
	return s.repo.GetByID(id)
}

func (s *promiseService) GetPublicByUserID(userID uuid.UUID) ([]models.PromiseV1, error) {
	return s.repo.GetPublicByUserID(userID)
}

func (s *promiseService) GetAllPromisesForAdmin() ([]models.PromiseV1, error) {
	return s.repo.GetAll()
}

func (s *promiseService) GetChildren(parentID uuid.UUID) ([]models.PromiseV1, error) {
	return s.repo.GetChildren(parentID)
}

func (s *promiseService) Update(userID uuid.UUID, id string, req dto.UpdatePromiseRequest, isAdmin bool) error {

	promiseID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("неверный ID")
	}

	existingPromise, err := s.repo.GetByID(promiseID)
	if err != nil {
		return errors.New("обещание не найдено")
	}

	if !isAdmin && existingPromise.UserID != userID {
		return errors.New("нельзя обновить чужое обещание")
	}

	if existingPromise.UserID != userID && !isAdmin {
		return errors.New("нет прав на редактирование")
	}

	if req.Title != nil {
		existingPromise.Title = strings.TrimSpace(*req.Title)
	}
	if req.Description != nil {
		existingPromise.Description = strings.TrimSpace(*req.Description)
	}
	if req.Status != nil {
		existingPromise.Status = *req.Status
	}
	if req.IsPrivate != nil {
		existingPromise.IsPrivate = *req.IsPrivate
	}
	if req.Deadline != nil {
		existingPromise.Deadline = *req.Deadline
	}

	return s.repo.Update(existingPromise)
}

func (s *promiseService) Delete(userID uuid.UUID, promiseID uuid.UUID, isAdmin bool) error {
	promise, err := s.repo.GetByID(promiseID)
	if err != nil {
		return err
	}
	if !isAdmin && promise.UserID != userID {
		return errors.New("нет прав на удаление")
	}

	children, err := s.repo.GetAllDescendants(promiseID)
	if err != nil {
		return err
	}
	for _, child := range children {
		if err := s.repo.Delete(child.ID); err != nil {
			return err
		}
	}
	return s.repo.Delete(promiseID)
}
