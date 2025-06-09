package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/services"
	"github.com/raxaris/ipromise-backend/internal/utils"
)

type MicrotaskHandler struct {
	service services.MicrotaskService
}

func NewMicrotaskHandler(service services.MicrotaskService) *MicrotaskHandler {
	return &MicrotaskHandler{service: service}
}

// CreateMicrotask godoc
// @Summary Создать микротаск
// @Description Добавляет микротаск к определённому промису
// @Tags microtasks
// @Security BearerAuth
// @Param promise_id path string true "ID промиса"
// @Accept json
// @Produce json
// @Param input body dto.CreateMicrotaskRequest true "Данные микротаска"
// @Success 201 {object} map[string]string "message: Микротаск создан"
// @Failure 400 {object} map[string]string "error: Ошибка запроса"
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 403 {object} map[string]string "error: Нет прав доступа"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /promises/{promise_id}/microtasks [post]
func (h *MicrotaskHandler) CreateMicrotask(c *gin.Context) {
	var req dto.CreateMicrotaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Неверный формат данных")
		return
	}

	promiseID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Некорректный ID промиса")
		return
	}

	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	err = h.service.CreateMicrotask(c.Request.Context(), userID, promiseID, req.Title, req.StepsPlanned)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusCreated, gin.H{"message": "Микротаск создан"})
}

// UpdateMicrotask godoc
// @Summary Обновить микротаск
// @Description Позволяет изменить заголовок и/или статус микротаска
// @Tags microtasks
// @Security BearerAuth
// @Param id path string true "ID микротаска"
// @Accept json
// @Produce json
// @Param input body dto.UpdateMicrotaskRequest true "Обновляемые данные"
// @Success 200 {object} map[string]string "message: Микротаск обновлён"
// @Failure 400 {object} map[string]string "error: Ошибка запроса"
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 403 {object} map[string]string "error: Нет доступа"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /microtasks/{id} [patch]
func (h *MicrotaskHandler) UpdateMicrotask(c *gin.Context) {
	var req dto.UpdateMicrotaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Неверный формат данных")
		return
	}

	if req.Title == nil && req.Status == nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Нужно указать хотя бы одно поле для обновления")
		return
	}

	microtaskID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Некорректный ID микротаски")
		return
	}

	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	err = h.service.UpdateMicrotask(c.Request.Context(), userID, microtaskID, req.Title, req.Status)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"message": "Микротаск обновлён"})
}

// DeleteMicrotask godoc
// @Summary Удалить микротаск
// @Description Удаляет микротаск (soft delete)
// @Tags microtasks
// @Security BearerAuth
// @Param id path string true "ID микротаска"
// @Produce json
// @Success 200 {object} map[string]string "message: Микротаск удалён"
// @Failure 400 {object} map[string]string "error: Ошибка ID"
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 403 {object} map[string]string "error: Нет прав"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /microtasks/{id} [delete]
func (h *MicrotaskHandler) DeleteMicrotask(c *gin.Context) {
	microtaskID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Некорректный ID микротаски")
		return
	}

	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	err = h.service.DeleteMicrotask(c.Request.Context(), userID, microtaskID)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"message": "Микротаск удалён"})
}

// ListMicrotasksByPromiseID godoc
// @Summary Получить список микротасков промиса
// @Description Возвращает все микротаски, прикреплённые к заданному промису
// @Tags microtasks
// @Security BearerAuth
// @Param promise_id path string true "ID промиса"
// @Produce json
// @Success 200 {array} dto.MicrotaskResponse
// @Failure 400 {object} map[string]string "error: Ошибка ID"
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 403 {object} map[string]string "error: Приватный промис"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /promises/{promise_id}/microtasks [get]
func (h *MicrotaskHandler) ListMicrotasksByPromiseID(c *gin.Context) {
	promiseID, err := uuid.Parse(c.Param("promise_id"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Некорректный ID промиса")
		return
	}

	viewerID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	microtasks, err := h.service.ListMicrotasksByPromiseID(c.Request.Context(), viewerID, promiseID)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, microtasks)
}

// ReorderMicrotasks godoc
// @Summary Обновить порядок микротасков
// @Description Массовое переупорядочивание микротасков внутри промиса
// @Tags microtasks
// @Security BearerAuth
// @Param promise_id path string true "ID промиса"
// @Accept json
// @Produce json
// @Param input body dto.ReorderMicrotasksMapRequest true "Порядок микротасков"
// @Success 200 {object} map[string]string "message: Микротаски упорядочены"
// @Failure 400 {object} map[string]string "error: Неверный ввод"
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 403 {object} map[string]string "error: Нет доступа"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /promises/{promise_id}/microtasks/reorder [patch]
func (h *MicrotaskHandler) ReorderMicrotasks(c *gin.Context) {
	var req dto.ReorderMicrotasksMapRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Неверный формат запроса")
		return
	}

	promiseID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Некорректный ID промиса")
		return
	}

	if len(req.Orders) == 0 {
		utils.RespondWithError(c, http.StatusBadRequest, "Список упорядочивания пуст")
		return
	}

	if err := h.service.ReorderMicrotasks(c.Request.Context(), promiseID, req.Orders); err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"message": "Микротаски упорядочены"})
}
