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
	promiseService services.PromiseService
	userService    services.UserService
}

func NewPromiseHandler(promiseService services.PromiseService, userService services.UserService) *PromiseHandler {
	return &PromiseHandler{
		promiseService: promiseService,
		userService:    userService,
	}
}

// ✅ POST /promises
func (h *PromiseHandler) CreatePromise(c *gin.Context) {
	var req dto.CreatePromiseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Неверный формат данных")
		return
	}

	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	err = h.promiseService.CreatePromise(c.Request.Context(), userID, req.Title, req.Description, req.Deadline, req.IsPrivate)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusCreated, gin.H{"message": "Обещание создано"})
}

// ✅ GET /promises/:id
func (h *PromiseHandler) GetPromiseByID(c *gin.Context) {
	promiseIDStr := c.Param("id")
	promiseID, err := uuid.Parse(promiseIDStr)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Некорректный ID")
		return
	}

	viewerID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	promise, err := h.promiseService.GetPromiseByID(c.Request.Context(), viewerID, promiseID)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	// Преобразуем в DTO
	resp := dto.PromiseResponse{
		ID:          promise.ID.String(),
		Title:       promise.Title,
		Description: promise.Description,
		Deadline:    promise.Deadline,
		IsPrivate:   promise.IsPrivate,
		Status:      promise.Status,
		CreatedAt:   promise.CreatedAt,
	}

	utils.RespondWithSuccess(c, http.StatusOK, resp)
}

// ✅ Обновить обещание
func (h *PromiseHandler) UpdatePromise(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	promiseIDStr := c.Param("id")
	promiseID, err := uuid.Parse(promiseIDStr)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Неверный формат ID")
		return
	}

	var req dto.UpdatePromiseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Неверный формат тела запроса")
		return
	}

	err = h.promiseService.UpdatePromise(c, userID, promiseID, req.Title, req.Description, req.Deadline)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Обещание обновлено")
}

// ✅ Удалить обещание
func (h *PromiseHandler) DeletePromise(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	promiseIDStr := c.Param("id")
	promiseID, err := uuid.Parse(promiseIDStr)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Неверный формат ID")
		return
	}

	err = h.promiseService.DeletePromise(c, userID, promiseID)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Обещание удалено")
}

// ✅ Получить список обещаний профиля
func (h *PromiseHandler) ListProfilePromises(c *gin.Context) {
	viewerID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	username := c.Param("username")
	profileUser, err := h.userService.GetUserByUsername(c, username)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	limit, after := utils.ParsePaginationParams(c)

	promises, err := h.promiseService.ListProfilePromises(c, viewerID, profileUser.ID, limit, after)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, promises)
}

// ✅ Лента (обещания фолловеров)
func (h *PromiseHandler) ListFeedPromises(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	limit, after := utils.ParsePaginationParams(c)

	promises, err := h.promiseService.ListFeedPromises(c, userID, limit, after)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, promises)
}

// ✅ Публичные обещания (все)
func (h *PromiseHandler) ListPublicPromises(c *gin.Context) {
	limit, after := utils.ParsePaginationParams(c)

	promises, err := h.promiseService.ListPublicPromises(c, limit, after)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, promises)
}
