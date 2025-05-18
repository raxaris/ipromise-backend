package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/services"
	"github.com/raxaris/ipromise-backend/internal/utils"
	"net/http"
)

type LikeHandler struct {
	likeService services.LikeService
}

func NewLikeHandler(likeService services.LikeService) *LikeHandler {
	return &LikeHandler{likeService: likeService}
}

// LikePost godoc
// @Summary Лайкнуть пост
// @Description Ставит лайк посту. Если уже лайкнут — ошибка
// @Tags likes
// @Security BearerAuth
// @Param id path string true "ID поста"
// @Success 200 {object} map[string]string "message: Пост лайкнут"
// @Failure 400 {object} map[string]string "error: Некорректный ID"
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /posts/{id}/like [post]
func (h *LikeHandler) LikePost(c *gin.Context) {
	postID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Некорректный ID поста")
		return
	}

	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	if err := h.likeService.LikePost(c.Request.Context(), userID, postID); err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"message": "Пост лайкнут"})
}

// UnlikePost godoc
// @Summary Убрать лайк с поста
// @Description Удаляет лайк пользователя с поста. Если лайка не было — ошибка
// @Tags likes
// @Security BearerAuth
// @Param id path string true "ID поста"
// @Success 200 {object} map[string]string "message: Лайк удалён"
// @Failure 400 {object} map[string]string "error: Некорректный ID"
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /posts/{id}/unlike [post]
func (h *LikeHandler) UnlikePost(c *gin.Context) {
	postID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Некорректный ID поста")
		return
	}

	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	if err := h.likeService.UnlikePost(c.Request.Context(), userID, postID); err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"message": "Лайк удалён"})
}
