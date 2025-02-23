package handlers

import (
	"github.com/raxaris/ipromise-backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/services"
)

// CreatePromiseHandler создаёт новое обещание
// @Summary Создание нового обещания
// @Description Позволяет пользователю создать новое обещание
// @Tags promises
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body dto.CreatePromiseRequest true "Данные обещания"
// @Success 201 {object} map[string]string "message: Обещание успешно создано"
// @Failure 400 {object} map[string]string "error: Ошибка валидации"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /promises [post]
func CreatePromiseHandler(c *gin.Context) {
	var req dto.CreatePromiseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Получаем user_id из контекста
	userID, _ := uuid.Parse(c.GetString("user_id"))

	// Создаём обещание
	err := services.CreatePromise(userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Обещание успешно создано"})
}

// GetAllPromisesHandler получает все обещания
// @Summary Получение всех обещаний
// @Description Возвращает список всех обещаний
// @Tags promises
// @Security BearerAuth
// @Success 200 {array} models.Promise
// @Router /promises [get]
func GetAllPromisesHandler(c *gin.Context) {
	isAdmin := c.GetString("role") == "admin"

	var promises []models.Promise
	var err error

	if isAdmin {
		promises, err = services.GetAllPromises()
	} else {
		promises, err = services.GetAllPublicPromises() // 🔹 Только публичные обещания
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения обещаний"})
		return
	}

	c.JSON(http.StatusOK, promises)
}

func GetPromiseByIDHandler(c *gin.Context) {
	userID, _ := uuid.Parse(c.GetString("user_id"))
	isAdmin := c.GetString("role") == "admin"
	promiseID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID обещания"})
		return
	}

	promise, err := services.GetPromiseByID(promiseID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Обещание не найдено"})
		return
	}

	// ✅ Проверяем доступ: владелец или админ могут видеть обещание
	if promise.IsPrivate && promise.UserID != userID && !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Это приватное обещание"})
		return
	}

	c.JSON(http.StatusOK, promise)
}

// GetUserPromisesHandler получает список обещаний пользователя
// @Summary Получение обещаний пользователя
// @Description Возвращает список обещаний пользователя по его ID
// @Tags promises
// @Security BearerAuth
// @Param id path string true "ID пользователя"
// @Success 200 {array} models.Promise
// @Failure 400 {object} map[string]string "error: Неверный формат ID пользователя"
// @Failure 404 {object} map[string]string "error: Пользователь не найден"
// @Failure 500 {object} map[string]string "error: Ошибка получения обещаний"
// @Router /users/{id}/promises [get]
func GetUserPromisesHandler(c *gin.Context) {
	requestedUserID, err := uuid.Parse(c.Param("id")) // ID пользователя, чьи обещания запрашиваются
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID пользователя"})
		return
	}

	currentUserID, _ := uuid.Parse(c.GetString("user_id")) // ID текущего пользователя
	isAdmin := c.GetString("role") == "admin"

	// Получаем обещания пользователя
	promises, err := services.GetPromiseByUserID(requestedUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения обещаний"})
		return
	}

	// Если запрашивает не владелец и не админ – скрываем приватные обещания
	if requestedUserID != currentUserID && !isAdmin {
		var filteredPromises []models.Promise
		for _, promise := range promises {
			if !promise.IsPrivate {
				filteredPromises = append(filteredPromises, promise)
			}
		}
		promises = filteredPromises
	}

	c.JSON(http.StatusOK, promises)
}

// GetAllPublicPromisesHandler – получение всех публичных обещаний
func GetAllPublicPromisesHandler(c *gin.Context) {
	promises, err := services.GetAllPublicPromises()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения обещаний"})
		return
	}

	c.JSON(http.StatusOK, promises)
}

// UpdatePromiseHandler обновляет обещание (автор или админ)
// @Summary Обновление обещания
// @Description Позволяет обновить информацию о обещании
// @Tags promises
// @Security BearerAuth
// @Param id path string true "ID обещания"
// @Param input body dto.UpdatePromiseRequest true "Данные для обновления"
// @Success 200 {object} map[string]string "message: Обещание обновлено"
// @Failure 400 {object} map[string]string "error: Ошибка валидации"
// @Failure 403 {object} map[string]string "error: Нет прав на редактирование"
// @Router /promises/{id} [put]
func UpdatePromiseHandler(c *gin.Context) {
	var req dto.UpdatePromiseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Получаем ID пользователя
	userID, _ := uuid.Parse(c.GetString("user_id"))
	promiseID := c.Param("id")
	isAdmin := c.GetString("role") == "admin"

	// Обновляем обещание через сервис
	err := services.UpdatePromise(userID, promiseID, req, isAdmin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Обещание обновлено"})
}

// DeletePromiseHandler удаляет обещание (только для админов)
// @Summary Удаление обещания
// @Description Позволяет администратору удалить обещание по ID
// @Tags admin
// @Security BearerAuth
// @Param id path string true "ID обещания"
// @Success 200 {object} map[string]string "message: Обещание удалено"
// @Failure 403 {object} map[string]string "error: У вас нет прав на удаление обещания"
// @Failure 400 {object} map[string]string "error: Ошибка при удалении обещания"
// @Router /admin/promises/{id} [delete]
func DeletePromiseHandler(c *gin.Context) {
	// Проверяем, является ли пользователь админом
	isAdmin := c.GetString("role") == "admin"
	if !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "У вас нет прав на удаление обещания"})
		return
	}

	// ID обещания для удаления
	promiseID := c.Param("id")

	// Вызываем сервис удаления
	err := services.DeletePromise(promiseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Обещание удалено"})
}
