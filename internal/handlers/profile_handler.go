package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/services"
	"github.com/raxaris/ipromise-backend/internal/utils"
)

type ProfileHandler struct {
	profileService services.ProfileService
}

func NewProfileHandler(profileService services.ProfileService) *ProfileHandler {
	return &ProfileHandler{profileService: profileService}
}

// GET /profile/me
func (h *ProfileHandler) GetMyProfile(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	profile, err := h.profileService.GetMyProfile(c.Request.Context(), userID)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, profile)
}

// GET /profile/:username
func (h *ProfileHandler) GetPublicProfile(c *gin.Context) {
	username := c.Param("username")
	viewerID, _ := utils.GetUserIDFromContext(c) // допускаем ноль UUID если неавторизован

	profile, err := h.profileService.GetPublicProfile(c.Request.Context(), viewerID, username)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, profile)
}

// PATCH /profile
func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Некорректные данные профиля")
		return
	}

	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	if err := h.profileService.UpdateProfile(c.Request.Context(), userID, req); err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"message": "Профиль обновлён"})
}
