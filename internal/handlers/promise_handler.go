package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/services"
	"github.com/raxaris/ipromise-backend/internal/utils"
	"net/http"
)

type PromiseHandler struct {
	service services.PromiseService
}

func NewPromiseHandler(service services.PromiseService) *PromiseHandler {
	return &PromiseHandler{service: service}
}

// Создание обещания
func (h *PromiseHandler) Create(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	var req dto.CreatePromiseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(userID, req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Обещание создано"})
}

// Обещания текущего пользователя (приватные + публичные)
func (h *PromiseHandler) GetMy(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	promises, err := h.service.GetByUserID(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Ошибка получения обещаний"})
		return
	}

	c.JSON(200, promises)
}

// Публичные обещания другого пользователя
func (h *PromiseHandler) GetPublicByUserID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Неверный формат ID"})
		return
	}

	promises, err := h.service.GetPublicByUserID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Ошибка получения обещаний"})
		return
	}

	c.JSON(200, promises)
}

// Получить обещание по ID
func (h *PromiseHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Неверный формат ID"})
		return
	}

	promise, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Обещание не найдено"})
		return
	}

	c.JSON(200, promise)
}

// Получить вложенные обещания (прогресс)
func (h *PromiseHandler) GetChildren(c *gin.Context) {
	parentID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Неверный формат ID"})
		return
	}

	children, err := h.service.GetChildren(parentID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Ошибка получения вложенных обещаний"})
		return
	}

	c.JSON(200, children)
}

// Лента всех публичных обещаний
func (h *PromiseHandler) GetPublic(c *gin.Context) {
	promises, err := h.service.GetPublic()
	if err != nil {
		c.JSON(500, gin.H{"error": "Ошибка получения обещаний"})
		return
	}

	c.JSON(200, promises)
}

func (h *PromiseHandler) GetAllForAdmin(c *gin.Context) {
	promises, err := h.service.GetAllPromisesForAdmin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения обещаний"})
		return
	}
	c.JSON(http.StatusOK, promises)
}

// Обновить обещание
func (h *PromiseHandler) Update(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	promiseID := c.Param("id")
	isAdmin := c.GetString("role") == "admin"

	var req dto.UpdatePromiseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Update(userID, promiseID, req, isAdmin); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Обещание обновлено"})
}

// Удалить обещание
func (h *PromiseHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	isAdmin := c.GetString("role") == "admin"

	if !isAdmin {
		c.JSON(403, gin.H{"error": "Только администратор может удалять обещания"})
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Обещание удалено"})
}
