package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raxaris/ipromise-backend/internal/services"
	"github.com/raxaris/ipromise-backend/internal/utils"
)

type FollowHandler struct {
	followerService services.FollowerService
	userService     services.UserService
}

func NewFollowHandler(followerService services.FollowerService, userService services.UserService) *FollowHandler {
	return &FollowHandler{
		followerService: followerService,
		userService:     userService,
	}
}

// RequestFollow godoc
// @Summary Отправить запрос на подписку
// @Description Авторизованный пользователь отправляет запрос на подписку другому пользователю
// @Tags follow
// @Security BearerAuth
// @Param username path string true "Username, кому отправить запрос"
// @Produce json
// @Success 200 {object} utils.SuccessResponse{data=map[string]string} "message: Follow request sent"
// @Failure 400 {object} utils.ErrorResponse "error: Уже подписан или нельзя подписаться на себя"
// @Failure 404 {object} utils.ErrorResponse "error: Пользователь не найден"
// @Failure 401 {object} utils.ErrorResponse "error: Неавторизован"
// @Failure 500 {object} utils.ErrorResponse "error: Внутренняя ошибка сервера"
// @Router /follow/{username} [post]
func (h *FollowHandler) RequestFollow(c *gin.Context) {
	targetUsername := c.Param("username")
	targetUser, err := h.userService.GetUserByUsername(c, targetUsername)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "user not found")
		return
	}

	currentUserID, err := utils.GetUserIDFromContext(c)
	if err := h.followerService.RequestFollow(c, currentUserID, targetUser.ID); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "follow request sent")
}

// AcceptFollowRequest godoc
// @Summary Принять запрос на подписку
// @Description Подтверждает входящий запрос на подписку от указанного пользователя
// @Tags follow
// @Security BearerAuth
// @Param username path string true "Username отправителя запроса"
// @Produce json
// @Success 200 {object} utils.SuccessResponse{data=map[string]string} "message: Follow request accepted"
// @Failure 400 {object} utils.ErrorResponse "error: Нет ожидающего запроса от пользователя"
// @Failure 404 {object} utils.ErrorResponse "error: Пользователь не найден"
// @Failure 401 {object} utils.ErrorResponse "error: Неавторизован"
// @Failure 500 {object} utils.ErrorResponse "error: Внутренняя ошибка сервера"
// @Router /follow/{username}/accept [post]
func (h *FollowHandler) AcceptFollowRequest(c *gin.Context) {
	fromUsername := c.Param("username")
	fromUser, err := h.userService.GetUserByUsername(c, fromUsername)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "user not found")
		return
	}

	currentUserID, err := utils.GetUserIDFromContext(c)
	if err := h.followerService.AcceptFollowRequest(c, currentUserID, fromUser.ID); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "follow request accepted")
}

// DeclineFollowRequest godoc
// @Summary Отклонить запрос на подписку
// @Description Отклоняет входящий запрос на подписку от указанного пользователя
// @Tags follow
// @Security BearerAuth
// @Param username path string true "Username отправителя запроса"
// @Produce json
// @Success 200 {object} utils.SuccessResponse{data=map[string]string} "message: Follow request declined"
// @Failure 400 {object} utils.ErrorResponse "error: Нет ожидающего запроса от пользователя"
// @Failure 404 {object} utils.ErrorResponse "error: Пользователь не найден"
// @Failure 401 {object} utils.ErrorResponse "error: Неавторизован"
// @Failure 500 {object} utils.ErrorResponse "error: Внутренняя ошибка сервера"
// @Router /follow/{username}/decline [post]
func (h *FollowHandler) DeclineFollowRequest(c *gin.Context) {
	fromUsername := c.Param("username")
	fromUser, err := h.userService.GetUserByUsername(c, fromUsername)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "user not found")
		return
	}

	currentUserID, err := utils.GetUserIDFromContext(c)
	if err := h.followerService.DeclineFollowRequest(c, currentUserID, fromUser.ID); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "follow request declined")
}

// Unfollow godoc
// @Summary Отписаться от пользователя
// @Description Удаляет подписку на указанного пользователя
// @Tags follow
// @Security BearerAuth
// @Param username path string true "Username, от кого отписаться"
// @Produce json
// @Success 200 {object} utils.SuccessResponse{data=map[string]string} "message: Unfollowed successfully"
// @Failure 400 {object} utils.ErrorResponse "error: Не подписан на этого пользователя или пытаешься отписаться от себя"
// @Failure 404 {object} utils.ErrorResponse "error: Пользователь не найден"
// @Failure 401 {object} utils.ErrorResponse "error: Неавторизован"
// @Failure 500 {object} utils.ErrorResponse "error: Внутренняя ошибка сервера"
// @Router /follow/{username} [delete]
func (h *FollowHandler) Unfollow(c *gin.Context) {
	username := c.Param("username")
	user, err := h.userService.GetUserByUsername(c, username)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "user not found")
		return
	}

	currentUserID, err := utils.GetUserIDFromContext(c)
	if err := h.followerService.Unfollow(c, currentUserID, user.ID); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "unfollowed successfully")
}

// CancelFollowRequest godoc
// @Summary Отменить запрос на подписку
// @Description Удаляет запрос на подписку, если он ещё не принят (status = "pending")
// @Tags follows
// @Security BearerAuth
// @Param username path string true "Username пользователя, которому был отправлен запрос"
// @Success 200 {object} map[string]string "message: Запрос отменён"
// @Failure 400 {object} map[string]string "error: Неверный формат"
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /follow/requests/{username}/cancel [delete]
func (h *FollowHandler) CancelFollowRequest(c *gin.Context) {
	username := c.Param("username")

	targetUser, err := h.userService.GetUserByUsername(c.Request.Context(), username)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Пользователь не найден")
		return
	}

	currentUserID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	err = h.followerService.CancelFollowRequest(c.Request.Context(), currentUserID, targetUser.ID)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"message": "Запрос на подписку отменён"})
}

// ListFriends godoc
// @Summary Получить список друзей (взаимных подписок)
// @Description Возвращает список пользователей, у которых с данным пользователем установлена взаимная подписка (дружба)
// @Tags follow
// @Security BearerAuth
// @Param username path string true "Username пользователя, чьих друзей получить"
// @Produce json
// @Success 200 {object} utils.SuccessResponse{data=[]dto.UserLiteResponse} "Список взаимных друзей"
// @Failure 404 {object} utils.ErrorResponse "error: Пользователь не найден"
// @Failure 401 {object} utils.ErrorResponse "error: Неавторизован"
// @Failure 500 {object} utils.ErrorResponse "error: Внутренняя ошибка сервера"
// @Router /friends/{username} [get]
func (h *FollowHandler) ListFriends(c *gin.Context) {
	username := c.Param("username")
	friends, err := h.followerService.ListFriendsByUsername(c.Request.Context(), username)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, friends)
}

// ListPendingRequests godoc
// @Summary Получить входящие заявки в друзья
// @Description Возвращает список пользователей, которые отправили запрос на добавление в друзья (ожидают подтверждения)
// @Tags follow
// @Security BearerAuth
// @Produce json
// @Success 200 {object} utils.SuccessResponse{data=[]dto.UserLiteResponse}
// @Failure 401 {object} utils.ErrorResponse "error: Неавторизован"
// @Failure 500 {object} utils.ErrorResponse "error: Внутренняя ошибка сервера"
// @Router /follow/requests [get]
func (h *FollowHandler) ListPendingRequests(c *gin.Context) {
	currentUserID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	requests, err := h.followerService.ListPendingRequests(c.Request.Context(), currentUserID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, requests)
}

// ListSentRequests godoc
// @Summary Отправленные заявки в друзья
// @Description Возвращает список пользователей, которым текущий пользователь отправил follow-запросы
// @Tags follow
// @Security BearerAuth
// @Produce json
// @Success 200 {object} utils.SuccessResponse{data=[]dto.UserLiteResponse}
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /follow/requests/sent [get]
func (h *FollowHandler) ListSentRequests(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	requests, err := h.followerService.ListSentRequests(c.Request.Context(), userID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, requests)
}

// GetRecommendedUsers godoc
// @Summary Рекомендуемые друзья
// @Description Возвращает список пользователей с общими друзьями, затем случайных
// @Tags follow
// @Security BearerAuth
// @Param limit query int false "Максимум пользователей"
// @Param after query string false "Дата (RFC3339) для пагинации"
// @Produce json
// @Success 200 {object} utils.SuccessResponse{data=[]dto.UserLiteResponse}
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /follow/recommended [get]
func (h *FollowHandler) GetRecommendedUsers(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	limit, after := utils.ParsePaginationParams(c)

	users, err := h.followerService.GetRecommendedUsers(c, userID, limit, after)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, users)
}
