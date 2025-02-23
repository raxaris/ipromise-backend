package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/services"
)

// CreatePromiseHandler – создать новое обещание (только юзер)
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

// UpdatePromiseHandler – обновить обещание (только автор или админ)
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

// DeletePromiseHandler – удалить обещание (только админ)
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

// GetAllPromisesHandler – получить все обещания
func GetAllPromisesHandler(c *gin.Context) {
	promises, err := services.GetAllPromises()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения обещаний"})
		return
	}

	c.JSON(http.StatusOK, promises)
}

func GetPromiseByIDHandler(c *gin.Context) {
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

	c.JSON(http.StatusOK, promise)
}

// GetPromisesByUserIDHandler – получить обещания конкретного пользователя
func GetPromisesByUserIDHandler(c *gin.Context) {
	// Получаем `user_id` из параметра запроса
	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID пользователя"})
		return
	}

	// Получаем обещания пользователя
	promises, err := services.GetPromiseByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения обещаний"})
		return
	}

	c.JSON(http.StatusOK, promises)
}
