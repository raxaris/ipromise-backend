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

// ✅ POST /microtasks/:microtask_id/posts
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

// ✅ PATCH /posts/:id
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

// ✅ DELETE /posts/:id
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

// ✅ GET /microtasks/:microtask_id/posts
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

// ✅ GET /posts/:id/replies
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
