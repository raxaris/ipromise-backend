package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/services"
	"github.com/raxaris/ipromise-backend/internal/utils"
	"net/http"
)

type PostHandler struct {
	service services.PostService
}

func NewPostHandler(service services.PostService) *PostHandler {
	return &PostHandler{service: service}
}

// CreatePost godoc
// @Summary Создать пост или комментарий
// @Description Создаёт корневой пост или комментарий (если указан parent_id)
// @Tags posts
// @Security BearerAuth
// @Param microtask_id path string true "ID микротаска"
// @Accept json
// @Produce json
// @Param input body dto.CreatePostRequest true "Данные поста"
// @Success 201 {object} map[string]string "message: Пост создан"
// @Failure 400 {object} map[string]string "error: Неверный формат"
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 403 {object} map[string]string "error: Нет прав доступа"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /microtasks/{microtask_id}/posts [post]
func (h *PostHandler) CreatePost(c *gin.Context) {
	var req dto.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Неверный формат запроса")
		return
	}

	microtaskID, err := uuid.Parse(c.Param("microtask_id"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Некорректный ID микротаска")
		return
	}

	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	// ⚠️ пока передаём пустой PromiseID — на будущее
	err = h.service.CreatePost(c.Request.Context(), userID, uuid.Nil, microtaskID, req.Content, req.ParentID)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusCreated, gin.H{"message": "Пост создан"})
}

// UpdatePost godoc
// @Summary Обновить пост
// @Description Обновляет содержимое поста
// @Tags posts
// @Security BearerAuth
// @Param id path string true "ID поста"
// @Accept json
// @Produce json
// @Param input body dto.UpdatePostRequest true "Новое содержимое"
// @Success 200 {object} map[string]string "message: Пост обновлён"
// @Failure 400 {object} map[string]string "error: Неверный формат"
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 403 {object} map[string]string "error: Нет прав"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /posts/{id} [patch]
func (h *PostHandler) UpdatePost(c *gin.Context) {
	var req dto.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Неверный формат данных")
		return
	}

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

	if err := h.service.UpdatePost(c.Request.Context(), userID, postID, req.Content); err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"message": "Пост обновлён"})
}

// DeletePost godoc
// @Summary Удалить пост
// @Description Удаляет пост по ID (soft delete)
// @Tags posts
// @Security BearerAuth
// @Param id path string true "ID поста"
// @Produce json
// @Success 200 {object} map[string]string "message: Пост удалён"
// @Failure 400 {object} map[string]string "error: Неверный ID"
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 403 {object} map[string]string "error: Нет прав"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /posts/{id} [delete]
func (h *PostHandler) DeletePost(c *gin.Context) {
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

	if err := h.service.DeletePost(c.Request.Context(), userID, postID); err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"message": "Пост удалён"})
}

// ListRootPosts godoc
// @Summary Получить посты по микротаску
// @Description Возвращает все корневые посты, связанные с микротаском
// @Tags posts
// @Security BearerAuth
// @Param microtask_id path string true "ID микротаска"
// @Param limit query int false "Максимальное количество"
// @Param after query string false "Дата (RFC3339) для пагинации"
// @Param after_id query string false "ID поста для пагинации"
// @Produce json
// @Success 200 {array} dto.PostResponse
// @Failure 400 {object} map[string]string "error: Неверный ID"
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /microtasks/{microtask_id}/posts [get]
func (h *PostHandler) ListRootPosts(c *gin.Context) {
	microtaskID, err := uuid.Parse(c.Param("microtask_id"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Некорректный ID микротаска")
		return
	}

	limit, afterCreatedAt, afterID := utils.ParseCursorPaginationParams(c)

	posts, err := h.service.ListRootPosts(c.Request.Context(), microtaskID, limit, afterCreatedAt, afterID)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, posts)
}

// ListReplies godoc
// @Summary Получить комментарии к посту
// @Description Возвращает список дочерних постов (реплаев)
// @Tags posts
// @Security BearerAuth
// @Param id path string true "ID родительского поста"
// @Param limit query int false "Максимальное количество"
// @Param after query string false "Дата (RFC3339) для пагинации"
// @Param after_id query string false "ID поста для пагинации"
// @Produce json
// @Success 200 {array} dto.PostResponse
// @Failure 400 {object} map[string]string "error: Неверный ID"
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /posts/{id}/replies [get]
func (h *PostHandler) ListReplies(c *gin.Context) {
	parentID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Некорректный ID поста")
		return
	}

	limit, afterCreatedAt, afterID := utils.ParseCursorPaginationParams(c)

	posts, err := h.service.ListReplies(c.Request.Context(), parentID, limit, afterCreatedAt, afterID)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, posts)
}
