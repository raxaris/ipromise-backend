package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/raxaris/ipromise-backend/internal/services"
	"github.com/raxaris/ipromise-backend/internal/utils"
	"net/http"
)

type BadgeHandler struct {
	badgeService services.BadgeService
}

func NewBadgeHandler(badgeService services.BadgeService) *BadgeHandler {
	return &BadgeHandler{badgeService: badgeService}
}

// ListAllBadges godoc
// @Summary Получить все бейджи
// @Description Возвращает список всех доступных бейджей системы
// @Tags badges
// @Produce json
// @Success 200 {array} dto.BadgeResponse
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /badges [get]
func (h *BadgeHandler) ListAllBadges(c *gin.Context) {
	badges, err := h.badgeService.ListAllBadges(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, badges)
}

// ListUserBadges godoc
// @Summary Получить бейджи текущего пользователя
// @Description Возвращает список бейджей, полученных текущим авторизованным пользователем
// @Tags badges
// @Security BearerAuth
// @Produce json
// @Success 200 {array} dto.BadgeResponse
// @Failure 401 {object} map[string]string "error: Требуется авторизация"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /badges/me [get]
func (h *BadgeHandler) ListUserBadges(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}
	badges, err := h.badgeService.GetUserBadges(c, userID)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, badges)
}
