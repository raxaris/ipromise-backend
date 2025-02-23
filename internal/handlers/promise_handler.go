package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/config"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/models"
)

// CreatePromiseHandler – создаёт новое обещание
func CreatePromiseHandler(c *gin.Context) {
	var req dto.CreatePromiseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Получаем ID пользователя из middleware
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	// Преобразуем userID в uuid.UUID
	userUUID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка обработки user_id"})
		return
	}

	// Создаём новое обещание
	promise := models.Promise{
		ID:          uuid.New(),
		UserID:      userUUID,
		ParentID:    req.ParentID,
		Title:       req.Title,
		Description: req.Description,
	}

	// Если это основное обещание (ParentID == nil)
	if req.ParentID == nil {
		promise.Status = "pending" // Основное обещание всегда создаётся со статусом pending
		if req.Deadline != nil {
			promise.Deadline = *req.Deadline
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Основное обещание требует дедлайна"})
			return
		}
	} else {
		// Если это прогресс (обновление обещания)
		var parentPromise models.Promise
		if err := config.DB.First(&parentPromise, "id = ?", req.ParentID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Родительское обещание не найдено"})
			return
		}

		// Наследуем дедлайн у родителя
		promise.Deadline = parentPromise.Deadline

		// Устанавливаем статус в зависимости от запроса
		if req.Status == "completed" {
			promise.Status = "completed"
		} else {
			promise.Status = "in_progress"
		}
	}

	// Сохраняем в БД
	if err := config.DB.Create(&promise).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания обещания"})
		return
	}

	// Возвращаем JSON-ответ
	c.JSON(http.StatusCreated, gin.H{"message": "Обещание создано", "promise": promise})
}

// GetPromiseHandler – получает обещание по ID
func GetPromiseHandler(c *gin.Context) {
	promiseID := c.Param("id")

	var promise models.Promise
	if err := config.DB.First(&promise, "id = ?", promiseID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Обещание не найдено"})
		return
	}

	c.JSON(http.StatusOK, promise)
}

// DeletePromiseHandler – удаляет (soft-delete) обещание
func DeletePromiseHandler(c *gin.Context) {
	promiseID := c.Param("id")

	var promise models.Promise
	if err := config.DB.First(&promise, "id = ?", promiseID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Обещание не найдено"})
		return
	}

	if err := config.DB.Delete(&promise).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления обещания"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Обещание удалено"})
}
