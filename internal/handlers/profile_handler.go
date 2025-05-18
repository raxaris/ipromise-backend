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

// GetMyProfile godoc
// @Summary Получить свой профиль
// @Description Возвращает расширенный профиль текущего пользователя
// @Tags profile
// @Security BearerAuth
// @Produce json
// @Success 200 {object} dto.MyProfileResponse
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /profile/me [get]
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

// GetPublicProfile godoc
// @Summary Получить публичный профиль пользователя
// @Description Возвращает публичный профиль по username (с расширенной статистикой, если взаимная подписка)
// @Tags profile
// @Produce json
// @Param username path string true "Имя пользователя"
// @Success 200 {object} dto.PublicProfileResponse
// @Failure 404 {object} map[string]string "error: Пользователь не найден"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /profile/{username} [get]
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

// UpdateProfile godoc
// @Summary Обновить профиль
// @Description Обновляет имя, аватар и био. Аватар предварительно загружается как attachment и передаётся в avatar_url
// @Tags profile
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body dto.UpdateProfileRequest true "Поля профиля"
// @Success 200 {object} map[string]string "message: Профиль обновлён"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /profile [patch]
func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	var req dto.UpdateProfileFormRequest
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
