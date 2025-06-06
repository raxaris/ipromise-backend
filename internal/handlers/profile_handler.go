package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/services"
	"github.com/raxaris/ipromise-backend/internal/utils"
)

type ProfileHandler struct {
	profileService    services.ProfileService
	attachmentService services.AttachmentService
}

func NewProfileHandler(profileService services.ProfileService, attachmentService services.AttachmentService) *ProfileHandler {
	return &ProfileHandler{
		profileService:    profileService,
		attachmentService: attachmentService,
	}
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

// UpdateProfileWithAvatar godoc
// @Summary Обновить профиль с аватаркой
// @Description Обновляет имя, био и аватар пользователя одним multipart-запросом
// @Tags profile
// @Security BearerAuth
// @Accept multipart/form-data
// @Produce json
// @Param username formData string false "Новое имя пользователя"
// @Param bio formData string false "Описание профиля"
// @Param avatar formData file false "Аватарка пользователя"
// @Success 200 {object} map[string]string "message: Профиль обновлён"
// @Failure 400 {object} map[string]string "error: Неверные данные"
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /profile [patch]
func (h *ProfileHandler) UpdateProfileWithAvatar(c *gin.Context) {
	var form dto.UpdateProfileFormRequest
	if err := c.ShouldBind(&form); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Некорректные данные формы")
		return
	}

	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	var avatarURL *string
	if form.Avatar != nil {
		file, err := form.Avatar.Open()
		if err != nil {
			utils.RespondWithError(c, http.StatusBadRequest, "Ошибка при открытии аватара")
			return
		}
		defer file.Close()

		fileBytes := make([]byte, form.Avatar.Size)
		_, _ = file.Read(fileBytes)

		uploaded, err := h.attachmentService.UploadAvatar(
			c.Request.Context(),
			userID,
			fileBytes,
			form.Avatar.Filename,
			form.Avatar.Header.Get("Content-Type"),
		)
		if err != nil {
			utils.RespondWithMappedError(c, err)
			return
		}
		avatarURL = &uploaded.FileURL
	}

	updateData := dto.UpdateProfileRequest{
		Username:  form.Username,
		Bio:       form.Bio,
		AvatarURL: avatarURL,
	}

	if err := h.profileService.UpdateProfile(c.Request.Context(), userID, updateData); err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"message": "Профиль обновлён"})
}
