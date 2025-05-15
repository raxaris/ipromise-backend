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
	postService services.PostService
}

func NewPostHandler(postService services.PostService) *PostHandler {
	return &PostHandler{postService: postService}
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
	err = h.postService.CreatePost(c.Request.Context(), userID, uuid.Nil, microtaskID, req.Content, req.ParentID)
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

	if err := h.postService.UpdatePost(c.Request.Context(), userID, postID, req.Content); err != nil {
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

	if err := h.postService.DeletePost(c.Request.Context(), userID, postID); err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"message": "Пост удалён"})
}

// ListPostsByMicrotaskID godoc
// @Summary Получить посты по микротаску
// @Description Возвращает все корневые посты, связанные с микротаском
// @Tags posts
// @Security BearerAuth
// @Param microtask_id path string true "ID микротаска"
// @Param limit query int false "Максимальное количество"
// @Param after query string false "Дата (RFC3339) для пагинации"
// @Param after_id query string false "ID поста для пагинации"
// @Produce json
// @Success 200 {array}  dto.PostWithRepliesTreeResponse
// @Failure 400 {object} map[string]string "error: Неверный ID"
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 403 {object} map[string]string "error: Нет доступа"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /microtasks/{microtask_id}/posts [get]
func (h *PostHandler) ListPostsByMicrotaskID(c *gin.Context) {
	microtaskID, err := uuid.Parse(c.Param("microtask_id"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Некорректный ID микротаска")
		return
	}

	viewerID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	limit, afterCreatedAt, afterID := utils.ParseCursorPaginationParams(c)

	posts, err := h.postService.ListPostsByMicrotaskIDTree(c.Request.Context(), microtaskID, viewerID, limit, afterCreatedAt, afterID)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"posts": posts})
}

// ListPostsByPromiseID godoc
// @Summary Получить посты по промису
// @Description Возвращает все посты, относящиеся к данному промису
// @Tags posts
// @Security BearerAuth
// @Param promise_id path string true "ID промиса"
// @Param limit query int false "Максимум постов"
// @Param after query string false "Дата (RFC3339) пагинации"
// @Param after_id query string false "ID последнего поста"
// @Produce json
// @Success 200 {array}  dto.PostWithRepliesTreeResponse
// @Failure 400 {object} map[string]string "error: Неверный ID"
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 403 {object} map[string]string "error: Нет доступа"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /promises/{promise_id}/posts [get]
func (h *PostHandler) ListPostsByPromiseID(c *gin.Context) {
	promiseID, err := uuid.Parse(c.Param("promise_id"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Некорректный ID промиса")
		return
	}

	viewerID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	limit, afterCreatedAt, afterID := utils.ParseCursorPaginationParams(c)

	posts, err := h.postService.ListPostsByPromiseIDTree(c.Request.Context(), promiseID, viewerID, limit, afterCreatedAt, afterID)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"posts": posts})
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
// @Success 200 {array} dto.PostWithRepliesTreeResponse
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

	viewerID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	limit, afterCreatedAt, afterID := utils.ParseCursorPaginationParams(c)

	replies, err := h.postService.ListReplies(c.Request.Context(), parentID, viewerID, limit, afterCreatedAt, afterID)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"replies": replies})
}

// GetFullPost godoc
// @Summary Получить пост с деревом комментариев
// @Description Возвращает пост, все его реплаи, автора, вложения и лайки
// @Tags posts
// @Security BearerAuth
// @Param id path string true "ID поста"
// @Produce json
// @Success 200 {object} dto.PostWithRepliesTreeResponse
// @Failure 400 {object} map[string]string "error: Неверный ID"
// @Failure 404 {object} map[string]string "error: Пост не найден"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /posts/{id}/full [get]
func (h *PostHandler) GetFullPost(c *gin.Context) {
	postID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid post ID format")
		return
	}
	viewerID, err := utils.GetUserIDFromContext(c)
	tree, err := h.postService.GetPostByID(c.Request.Context(), postID, viewerID)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, tree)
}

// ListPublicPostsLite godoc
// @Summary Публичные посты (лайт)
// @Description Публичная лента постов без дерева комментариев
// @Tags posts
// @Security BearerAuth
// @Param limit query int false "Максимум постов"
// @Param after query string false "Дата (RFC3339) для пагинации"
// @Param after_id query string false "ID последнего поста"
// @Produce json
// @Success 200 {array} dto.PostLiteResponse
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /posts/public [get]
func (h *PostHandler) ListPublicPostsLite(c *gin.Context) {
	limit, after, afterID := utils.ParseCursorPaginationParams(c)
	viewerID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	posts, err := h.postService.ListPublicPostsLite(c.Request.Context(), limit, after, afterID, viewerID)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"posts": posts})
}

// ListFeedPostsLite godoc
// @Summary Лента подписок (лайт)
// @Description Посты от тех, на кого подписан пользователь
// @Tags posts
// @Security BearerAuth
// @Param limit query int false "Максимум постов"
// @Param after query string false "Дата (RFC3339)"
// @Param after_id query string false "ID последнего поста"
// @Produce json
// @Success 200 {array} dto.PostLiteResponse
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /posts/feed [get]
func (h *PostHandler) ListFeedPostsLite(c *gin.Context) {
	limit, after, afterID := utils.ParseCursorPaginationParams(c)
	viewerID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	posts, err := h.postService.ListFeedPostsLite(c.Request.Context(), viewerID, limit, after, afterID)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"posts": posts})
}

// ListPublicPostsTree godoc
// @Summary Публичные посты (с деревом)
// @Description Публичная лента постов с комментариями
// @Tags posts
// @Security BearerAuth
// @Param limit query int false "Максимум постов"
// @Param after query string false "Дата (RFC3339)"
// @Param after_id query string false "ID последнего поста"
// @Produce json
// @Success 200 {array} dto.PostWithRepliesTreeResponse
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /posts/public/tree [get]
func (h *PostHandler) ListPublicPostsTree(c *gin.Context) {
	limit, after, afterID := utils.ParseCursorPaginationParams(c)
	viewerID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	posts, err := h.postService.ListPublicPostsTree(c.Request.Context(), limit, after, afterID, viewerID)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"posts": posts})
}

// ListFeedPostsTree godoc
// @Summary Лента подписок (с деревом)
// @Description Посты с реплаями от пользователей, на которых подписан текущий
// @Tags posts
// @Security BearerAuth
// @Param limit query int false "Максимум постов"
// @Param after query string false "Дата (RFC3339)"
// @Param after_id query string false "ID последнего поста"
// @Produce json
// @Success 200 {array} dto.PostWithRepliesTreeResponse
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /posts/feed/tree [get]
func (h *PostHandler) ListFeedPostsTree(c *gin.Context) {
	limit, after, afterID := utils.ParseCursorPaginationParams(c)
	viewerID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	posts, err := h.postService.ListFeedPostsTree(c.Request.Context(), viewerID, limit, after, afterID)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"posts": posts})
}

// ListUserPostsTree godoc
// @Summary Посты пользователя (с деревом)
// @Description Получает все посты пользователя по username, включая комментарии (если приватные – только для владельца)
// @Tags posts
// @Security BearerAuth
// @Param username path string true "Имя пользователя (username)"
// @Param limit query int false "Максимум постов"
// @Param after query string false "Дата (RFC3339) для пагинации"
// @Param after_id query string false "ID последнего поста"
// @Produce json
// @Success 200 {array} dto.PostWithRepliesTreeResponse
// @Failure 400 {object} map[string]string "error: Некорректный username"
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /posts/user/{username} [get]
func (h *PostHandler) ListUserPostsTree(c *gin.Context) {
	username := c.Param("username")
	viewerID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	limit, after, afterID := utils.ParseCursorPaginationParams(c)

	posts, err := h.postService.ListUserPostsTree(c.Request.Context(), username, viewerID, limit, after, afterID)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"posts": posts})
}
