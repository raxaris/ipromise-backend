package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/repositories"
)

// Ошибки
var (
	ErrNotAllowedToUpdate = errors.New("Вы не можете редактировать это обещание")
	ErrInvalidStatus      = errors.New("Нельзя изменить статус на этот")
	ErrPromiseNotFound    = errors.New("Обещание не найдено")
)

// CreatePromise – создание нового обещания (юзер/админ)
func CreatePromise(userID uuid.UUID, promise *models.Promise) error {
	// Если у обещания есть родитель, проверяем его существование
	if promise.ParentID != nil {
		parentPromise, err := repositories.GetPromiseByID(promise.ParentID.String())
		if err != nil {
			return errors.New("Родительское обещание не найдено")
		}
		promise.Deadline = parentPromise.Deadline // Наследуем дедлайн
	}

	// Создаём обещание в БД
	return repositories.CreatePromise(promise)
}

// UpdatePromise – обновление обещания (с учетом ролей)
func UpdatePromise(userID uuid.UUID, promiseID string, updateData *models.Promise, isAdmin bool) error {
	// Получаем текущее обещание
	existingPromise, err := repositories.GetPromiseByID(promiseID)
	if err != nil {
		return ErrPromiseNotFound
	}

	// 1️⃣ Проверяем, имеет ли право пользователь редактировать обещание
	if existingPromise.UserID != userID && !isAdmin {
		return ErrNotAllowedToUpdate
	}

	// 2️⃣ Нельзя менять `ParentID`
	if existingPromise.ParentID != nil && updateData.ParentID != nil && *existingPromise.ParentID != *updateData.ParentID {
		return errors.New("Нельзя изменять родительское обещание")
	}

	// 3️⃣ Нельзя менять `Deadline`, если это прогресс
	if existingPromise.ParentID != nil && !updateData.Deadline.Equal(existingPromise.Deadline) {
		return errors.New("Нельзя менять дедлайн у прогресса")
	}

	// 4️⃣ Проверяем корректность изменения статуса
	validTransitions := map[string][]string{
		"pending":     {"in_progress", "completed"},
		"in_progress": {"completed"},
		"completed":   {},
	}

	allowedNextStatuses, ok := validTransitions[existingPromise.Status]
	if !ok {
		return ErrInvalidStatus
	}

	// Проверяем, разрешено ли изменение статуса
	statusValid := false
	for _, allowed := range allowedNextStatuses {
		if updateData.Status == allowed {
			statusValid = true
			break
		}
	}

	if !statusValid && existingPromise.Status != updateData.Status {
		return ErrInvalidStatus
	}

	// ✅ Всё в порядке – обновляем данные
	existingPromise.Title = updateData.Title
	existingPromise.Description = updateData.Description
	existingPromise.Status = updateData.Status

	// Сохраняем обновления
	return repositories.UpdatePromise(existingPromise)
}

// DeletePromise – удаление обещания (только для админа/модератора)
func DeletePromise(promiseID string) error {
	return repositories.DeletePromise(promiseID)
}
