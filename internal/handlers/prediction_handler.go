package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/services"
	"github.com/raxaris/ipromise-backend/internal/utils"
)

type PredictionHandler struct {
	promiseService    services.PromiseService
	predictionService services.PredictionService
}

func NewPredictionHandler(predictionService services.PredictionService, promiseService services.PromiseService) *PredictionHandler {
	return &PredictionHandler{
		predictionService: predictionService,
		promiseService:    promiseService,
	}
}

// GetPrediction godoc
// @Summary Получить существующее предсказание по Promise
// @Description Возвращает сохранённое предсказание без повторной генерации
// @Tags prediction
// @Security BearerAuth
// @Produce json
// @Param id path string true "ID Promise"
// @Success 200 {object} dto.PredictionResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /promises/{id}/prediction [get]
func (h *PredictionHandler) GetPrediction(c *gin.Context) {
	promiseID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	prediction, err := h.predictionService.GetPrediction(c.Request.Context(), promiseID)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}
	if prediction == nil {
		utils.RespondWithError(c, http.StatusNotFound, "Prediction not found")
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, prediction)
}

// GeneratePrediction godoc
// @Summary Сгенерировать или обновить предсказание
// @Description Вызывает AI и сохраняет новое предсказание
// @Tags prediction
// @Security BearerAuth
// @Produce json
// @Param id path string true "ID Promise"
// @Success 200 {object} dto.PredictionResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /promises/{id}/prediction [post]
func (h *PredictionHandler) GeneratePrediction(c *gin.Context) {
	promiseID, err := uuid.Parse(c.Param("id"))
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

	prediction, err := h.predictionService.CreateOrReplacePrediction(c.Request.Context(), promise)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, prediction)
}
