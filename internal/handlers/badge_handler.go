package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/services"
	"github.com/raxaris/ipromise-backend/internal/utils"
	"net/http"
)

type BadgeHandler struct {
	badgeService services.BadgeService
	userService  services.UserService
}

func NewBadgeHandler(badgeService services.BadgeService, userService services.UserService) *BadgeHandler {
	return &BadgeHandler{
		badgeService: badgeService,
		userService:  userService,
	}
}

// CreateBadge godoc
// @Summary Создать новый бейдж
// @Description Админ может создать новый бейдж вручную
// @Tags badges
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body models.Badge true "Данные бейджа"
// @Success 201 {object} map[string]string "message: Бейдж создан"
// @Failure 400 {object} map[string]string "error: Некорректные данные"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /admin/badges [post]
func (h *BadgeHandler) CreateBadge(c *gin.Context) {
	var badge models.Badge
	if err := c.ShouldBindJSON(&badge); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Неверные данные")
		return
	}
	err := h.badgeService.CreateBadge(c.Request.Context(), &badge)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}
	utils.RespondWithSuccess(c, http.StatusCreated, gin.H{"message": "Бейдж создан"})
}

// AssignBadge godoc
// @Summary Назначить бейдж пользователю
// @Description Админ вручную назначает бейдж по коду
// @Tags badges
// @Security BearerAuth
// @Produce json
// @Param username path string true "Имя пользователя"
// @Param code path string true "Код бейджа"
// @Success 200 {object} map[string]string "message: Бейдж назначен"
// @Failure 404 {object} map[string]string "error: Пользователь или бейдж не найден"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /admin/badges/{username}/{code} [post]
func (h *BadgeHandler) AssignBadge(c *gin.Context) {
	username := c.Param("username")
	code := c.Param("code")

	user, err := h.userService.GetUserByUsername(c.Request.Context(), username)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	err = h.badgeService.AssignBadgeByCode(c.Request.Context(), user.ID, code)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"message": "Бейдж назначен"})
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

// GetBadgesByUsername godoc
// @Summary Получить бейджи пользователя
// @Description Возвращает список бейджей по username
// @Tags badges
// @Produce json
// @Param username path string true "Имя пользователя"
// @Success 200 {array} dto.BadgeResponse
// @Failure 404 {object} map[string]string "error: Пользователь не найден"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /badges/{username} [get]
func (h *BadgeHandler) GetBadgesByUsername(c *gin.Context) {
	username := c.Param("username")

	user, err := h.userService.GetUserByUsername(c.Request.Context(), username)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	badges, err := h.badgeService.GetUserBadges(c.Request.Context(), user.ID)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, badges)
}
