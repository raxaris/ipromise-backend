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

// CreatePromise godoc
// @Summary Создать обещание
// @Description Пользователь создаёт новое обещание
// @Tags promises
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body dto.CreatePromiseRequest true "Данные обещания"
// @Success 201 {object} map[string]string "message: Обещание создано"
// @Failure 400 {object} map[string]string "error: Неверный формат данных"
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 500 {object} map[string]string "error: Внутренняя ошибка сервера"
// @Router /promises [post]
func (h *PromiseHandler) CreatePromise(c *gin.Context) {
	var req dto.CreatePromiseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request format")
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

	utils.RespondWithSuccess(c, http.StatusCreated, "Promise created")
}

// GetPromiseByID godoc
// @Summary Получить обещание по ID
// @Description Возвращает конкретное обещание (если публичное или пользователь — владелец)
// @Tags promises
// @Security BearerAuth
// @Produce json
// @Param id path string true "ID обещания"
// @Success 200 {object} dto.PromiseResponse
// @Failure 400 {object} map[string]string "error: Некорректный ID"
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 403 {object} map[string]string "error: Нет доступа"
// @Failure 404 {object} map[string]string "error: Обещание не найдено"
// @Router /promises/{id} [get]
func (h *PromiseHandler) GetPromiseByID(c *gin.Context) {
	promiseIDStr := c.Param("id")
	promiseID, err := uuid.Parse(promiseIDStr)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid ID format")
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
	res := dto.PromiseResponse{
		ID:          promise.ID.String(),
		Title:       promise.Title,
		Description: promise.Description,
		Deadline:    promise.Deadline,
		IsPrivate:   promise.IsPrivate,
		Status:      promise.Status,
		CreatedAt:   promise.CreatedAt,
	}

	utils.RespondWithSuccess(c, http.StatusOK, res)
}

// UpdatePromise godoc
// @Summary Обновить обещание
// @Description Обновляет название, описание или дедлайн обещания
// @Tags promises
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "ID обещания"
// @Param input body dto.UpdatePromiseRequest true "Обновлённые данные"
// @Success 200 {object} map[string]string "message: Обещание обновлено"
// @Failure 400 {object} map[string]string "error: Неверные данные"
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 403 {object} map[string]string "error: Нет доступа"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /promises/{id} [patch]
func (h *PromiseHandler) UpdatePromise(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	promiseIDStr := c.Param("id")
	promiseID, err := uuid.Parse(promiseIDStr)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	var req dto.UpdatePromiseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Malformed request body")
		return
	}

	err = h.promiseService.UpdatePromise(c.Request.Context(), userID, promiseID, req.Title, req.Description, req.Deadline, req.IsPrivate)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Promise updated")
}

// DeletePromise godoc
// @Summary Удалить обещание
// @Description Удаляет обещание, если пользователь является владельцем
// @Tags promises
// @Security BearerAuth
// @Produce json
// @Param id path string true "ID обещания"
// @Success 200 {object} map[string]string "message: Обещание удалено"
// @Failure 400 {object} map[string]string "error: Неверный формат ID"
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 403 {object} map[string]string "error: Нет прав"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /promises/{id} [delete]
func (h *PromiseHandler) DeletePromise(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	promiseIDStr := c.Param("id")
	promiseID, err := uuid.Parse(promiseIDStr)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	err = h.promiseService.DeletePromise(c.Request.Context(), userID, promiseID)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Promise deleted")
}

// ListProfilePromises godoc
// @Summary Обещания пользователя по username
// @Description Возвращает список обещаний в профиле пользователя (с учётом приватности)
// @Tags promises
// @Security BearerAuth
// @Produce json
// @Param username path string true "Username пользователя"
// @Param limit query int false "Лимит"
// @Param after query string false "Дата (RFC3339) для пагинации"
// @Success 200 {array} dto.PromiseResponse
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 404 {object} map[string]string "error: Пользователь не найден"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /promises/user/{username} [get]
func (h *PromiseHandler) ListProfilePromises(c *gin.Context) {
	viewerID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	username := c.Param("username")
	profileUser, err := h.userService.GetUserByUsername(c.Request.Context(), username)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	limit, after := utils.ParsePaginationParams(c)

	promises, err := h.promiseService.ListProfilePromises(c.Request.Context(), viewerID, profileUser.ID, limit, after)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, promises)
}

// ListFeedPromises godoc
// @Summary Лента обещаний
// @Description Возвращает обещания от фолловеров текущего пользователя
// @Tags promises
// @Security BearerAuth
// @Produce json
// @Param limit query int false "Лимит"
// @Param after query string false "Дата и время, начиная с которого загружать обещания (формат RFC3339)"
func (h *PromiseHandler) ListFeedPromises(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	limit, after := utils.ParsePaginationParams(c)

	promises, err := h.promiseService.ListFeedPromises(c.Request.Context(), userID, limit, after)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, promises)
}

// ✅ Публичные обещания (все)
func (h *PromiseHandler) ListPublicPromises(c *gin.Context) {
	limit, after := utils.ParsePaginationParams(c)

	promises, err := h.promiseService.ListPublicPromises(c.Request.Context(), limit, after)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, promises)
}
