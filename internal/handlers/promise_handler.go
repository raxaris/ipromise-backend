package handlers

import (
	"github.com/raxaris/ipromise-backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/services"
)

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
