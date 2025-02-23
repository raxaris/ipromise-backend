package services

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/repositories"
)

// Ошибки
var (
	ErrNotAllowedToUpdate = errors.New("вы не можете редактировать это обещание")
	ErrInvalidStatus      = errors.New("нельзя изменить статус на этот")
	ErrPromiseNotFound    = errors.New("обещание не найдено")
	ErrInvalidTitle       = errors.New("заголовок обещания не может быть пустым или короче 3 символов")
)

// CreatePromise – создание нового обещания (юзер/админ)
func CreatePromise(userID uuid.UUID, req dto.CreatePromiseRequest) error {
	// Убираем пробелы в заголовке и описании
	req.Title = strings.TrimSpace(req.Title)
	req.Description = strings.TrimSpace(req.Description)

	// Проверяем корректность заголовка
	if len(req.Title) < 5 {
		return ErrInvalidTitle
	}

	// Создаём новый объект обещания
	promise := models.Promise{
		ID:          uuid.New(),
		UserID:      userID,
		ParentID:    req.ParentID,
		Title:       req.Title,
		Description: req.Description,
		IsPrivate:   req.IsPrivate,
	}

	// Если это основное обещание (нет ParentID)
	if req.ParentID == nil {
		promise.Status = "pending" // Основное обещание всегда создаётся со статусом "pending"

		// Проверяем, указан ли дедлайн
		if req.Deadline != nil {
			promise.Deadline = *req.Deadline
		} else {
			return errors.New("основное обещание должно иметь дедлайн")
		}
	} else {
		// Если это прогресс (обновление обещания)
		parentPromise, err := repositories.GetPromiseByID(*req.ParentID)
		if err != nil {
			return errors.New("родительское обещание не найдено")
		}

		// Наследуем дедлайн у родителя
		promise.Deadline = parentPromise.Deadline

		// Прогресс может быть либо "in_progress", либо "completed"
		if req.Status != "in_progress" && req.Status != "completed" {
			return errors.New("прогресс должен быть 'in_progress' или 'completed'")
		}

		promise.Status = req.Status
	}

	// Создаём обещание в БД
	return repositories.CreatePromise(&promise)
}

// GetAllPublicPromises – получает только публичные обещания
func GetAllPublicPromises() ([]models.Promise, error) {
	return repositories.GetPublicPromises()
}

// GetPromiseByID – получает обещание по ID (проверка приватности)
func GetPromiseByID(promiseID uuid.UUID) (*models.Promise, error) {
	promise, err := repositories.GetPromiseByID(promiseID)
	if err != nil {
		return nil, errors.New("обещание не найдено")
	}

	// Если обещание приватное – вернуть ошибку
	if promise.IsPrivate {
		return nil, errors.New("обещание приватное")
	}

	return promise, nil
}

// GetAllPromises – получение всех обещаний
func GetAllPromises() ([]models.Promise, error) {
	return repositories.GetAllPromises()
}

// GetPromiseByUserID – получение обещаний пользователя
func GetPromiseByUserID(userID uuid.UUID) ([]models.Promise, error) {
	return repositories.GetPromisesByUserID(userID)
}

// UpdatePromise – обновление обещания (с учетом ролей)
func UpdatePromise(userID uuid.UUID, promiseID string, updateData dto.UpdatePromiseRequest, isAdmin bool) error {
	// Преобразуем promiseID в UUID
	promiseUUID, err := uuid.Parse(promiseID)
	if err != nil {
		return errors.New("Неверный формат ID обещания")
	}

	// Получаем текущее обещание
	existingPromise, err := repositories.GetPromiseByID(promiseUUID)
	if err != nil {
		return ErrPromiseNotFound
	}

	// 1️⃣ Проверяем, имеет ли право пользователь редактировать обещание
	if existingPromise.UserID != userID && !isAdmin {
		return ErrNotAllowedToUpdate
	}

	// 3️⃣ Нельзя менять `Deadline`, если это прогресс
	if existingPromise.ParentID != nil && updateData.Deadline != nil {
		return errors.New("Нельзя менять дедлайн у прогресса")
	}

	// 4️⃣ Проверяем корректность изменения статуса
	validTransitions := map[string]map[string]bool{
		"pending":     {"in_progress": true, "completed": true},
		"in_progress": {"completed": true},
		"completed":   {},
	}

	allowedNextStatuses, ok := validTransitions[existingPromise.Status]
	if !ok || (updateData.Status != nil && !allowedNextStatuses[*updateData.Status]) {
		return ErrInvalidStatus
	}

	// ✅ Всё в порядке – обновляем данные
	if updateData.Title != nil {
		existingPromise.Title = *updateData.Title
	}
	if updateData.Description != nil {
		existingPromise.Description = *updateData.Description
	}
	if updateData.Status != nil {
		existingPromise.Status = *updateData.Status
	}
	if updateData.IsPrivate != nil {
		if existingPromise.ParentID != nil {
			return errors.New("нельзя менять приватность у обновления прогресса")
		}
		existingPromise.IsPrivate = *updateData.IsPrivate
	}
	// Сохраняем обновления
	return repositories.UpdatePromise(existingPromise)
}

// DeletePromise – удаление обещания (только для админа/модератора)
func DeletePromise(promiseID string) error {
	// Преобразуем в UUID
	promiseUUID, err := uuid.Parse(promiseID)
	if err != nil {
		return errors.New("неверный формат ID обещания")
	}

	return repositories.DeletePromise(promiseUUID)
}
