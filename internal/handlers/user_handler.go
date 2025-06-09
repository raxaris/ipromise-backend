package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/services"
	"github.com/raxaris/ipromise-backend/internal/utils"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// GetUserByUsername godoc
// @Summary Получить пользователя по username
// @Description Возвращает краткую информацию о пользователе (id, username, role)
// @Tags users
// @Security BearerAuth
// @Produce json
// @Param username path string true "Имя пользователя"
// @Success 200 {object} map[string]interface{} "Пользователь найден"
// @Failure 404 {object} map[string]string "error: Пользователь не найден"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /users/{username} [get]
func (h *UserHandler) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")

	user, err := h.userService.GetUserByUsername(c.Request.Context(), username)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
	})
}

// UpdateUser godoc
// @Summary Обновить пользователя
// @Description Обновляет профиль пользователя (username, role), доступно владельцу или админу
// @Tags users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "ID пользователя"
// @Param input body dto.UpdateUserRequest true "Новые данные пользователя"
// @Success 200 {object} map[string]string "message: Профиль обновлен"
// @Failure 400 {object} map[string]string "error: Неверный формат данных"
// @Failure 403 {object} map[string]string "error: Нет прав на редактирование"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /users/{id} [patch]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	requesterID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}
	isAdmin := utils.IsAdmin(c)

	userIDStr := c.Param("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Неверный формат ID")
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Неверные данные")
		return
	}

	if err := h.userService.UpdateUser(c.Request.Context(), requesterID, userID, req, isAdmin); err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Профиль обновлен")
}
