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

// ✅ POST /promises/:promise_id/microtasks
func (h *MicrotaskHandler) CreateMicrotask(c *gin.Context) {
	var req dto.CreateMicrotaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Неверный формат данных")
		return
	}

	promiseID, err := uuid.Parse(c.Param("promise_id"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Некорректный ID промиса")
		return
	}

	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	err = h.service.CreateMicrotask(c.Request.Context(), userID, promiseID, req.Title, req.Order)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusCreated, gin.H{"message": "Микротаск создан"})
}

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

// ✅ DELETE /microtasks/:id
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

// ✅ GET /promises/:promise_id/microtasks
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

// ✅ PATCH /promises/:promise_id/microtasks/reorder
func (h *MicrotaskHandler) ReorderMicrotasks(c *gin.Context) {
	var req dto.ReorderMicrotasksMapRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Неверный формат запроса")
		return
	}

	promiseID, err := uuid.Parse(c.Param("promise_id"))
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
