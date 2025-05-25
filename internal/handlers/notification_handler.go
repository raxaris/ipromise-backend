package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
// @Description Возвращает список всех уведомлений, отправленных текущему пользователю
// @Tags notifications
// @Security BearerAuth
// @Produce json
// @Success 200 {object} utils.SuccessResponse{data=[]dto.NotificationResponse}
// @Failure 401 {object} utils.ErrorResponse "Неавторизован"
// @Failure 500 {object} utils.ErrorResponse "Ошибка сервера"
// @Router /notifications/me [get]
func (h *NotificationHandler) ListMyNotifications(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	notifications, err := h.service.GetMyNotifications(c.Request.Context(), userID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to load notifications")
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, notifications)
}

// MarkAsRead godoc
// @Summary Отметить уведомление как прочитанное
// @Description Обновляет статус уведомления по ID и помечает его как прочитанное
// @Tags notifications
// @Security BearerAuth
// @Param id path string true "ID уведомления"
// @Produce json
// @Success 200 {object} utils.SuccessResponse{data=string} "message: Marked as read"
// @Failure 400 {object} utils.ErrorResponse "Неверный формат ID"
// @Failure 401 {object} utils.ErrorResponse "Неавторизован"
// @Failure 500 {object} utils.ErrorResponse "Ошибка сервера"
// @Router /notifications/{id}/read [post]
func (h *NotificationHandler) MarkAsRead(c *gin.Context) {
	notificationID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.service.MarkAsRead(c.Request.Context(), notificationID); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to mark as read")
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Marked as read")
}
