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

// GET /badges
func (h *BadgeHandler) ListAllBadges(c *gin.Context) {
	badges, err := h.badgeService.ListAllBadges(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, badges)
}

// GET /me/badges
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
