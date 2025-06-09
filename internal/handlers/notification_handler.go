package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/services"
	"github.com/raxaris/ipromise-backend/internal/utils"
	"net/http"
)

type NotificationHandler struct {
	service services.NotificationService
}

func NewNotificationHandler(service services.NotificationService) *NotificationHandler {
	return &NotificationHandler{service: service}
}

// ListMyNotifications godoc
// @Summary Получить все мои уведомления
// @Description Возвращает список всех уведомлений пользователя
// @Tags notifications
// @Security BearerAuth
// @Produce json
// @Success 200 {object} utils.SuccessResponse{data=[]dto.NotificationResponse}
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /notifications/me [get]
func (h *NotificationHandler) ListMyNotifications(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	notifications, err := h.service.ListUserNotifications(c.Request.Context(), userID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to load notifications")
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, notifications)
}

// MarkManyAsRead godoc
// @Summary Массовая отметка уведомлений как прочитанных
// @Description Отмечает переданные уведомления как прочитанные
// @Tags notifications
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param ids body []string true "Массив ID уведомлений"
// @Success 200 {object} utils.SuccessResponse{data=string}
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /notifications/mark-read [post]
func (h *NotificationHandler) MarkManyAsRead(c *gin.Context) {
	var idsRaw []string
	if err := c.ShouldBindJSON(&idsRaw); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid JSON")
		return
	}

	var ids []uuid.UUID
	for _, idStr := range idsRaw {
		id, err := uuid.Parse(idStr)
		if err != nil {
			utils.RespondWithError(c, http.StatusBadRequest, "Invalid ID format")
			return
		}
		ids = append(ids, id)
	}

	if err := h.service.MarkManyAsRead(c.Request.Context(), ids); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to mark as read")
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Notifications marked as read")
}

// TestNotification godoc
// @Summary Тестовое уведомление
// @Tags debug
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body dto.NotificationRequest true "Test data"
// @Success 200 {object} utils.SuccessResponse
// @Router /debug/test-notification [post]
func (h *NotificationHandler) TestNotification(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	notification := &models.Notification{
		Type:    "test",
		Message: "🚨 This is a test notification",
	}

	if err := h.service.SendNotification(c.Request.Context(), userID, notification); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Sent")
}
