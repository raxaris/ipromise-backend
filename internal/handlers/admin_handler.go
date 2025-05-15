package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/raxaris/ipromise-backend/internal/services"
	"github.com/raxaris/ipromise-backend/internal/utils"
	"net/http"
)

type AdminHandler struct {
	adminService services.AdminService
}

func NewAdminHandler(adminService services.AdminService) *AdminHandler {
	return &AdminHandler{adminService: adminService}
}

func (h *AdminHandler) ListAllUsers(c *gin.Context) {
	users, err := h.adminService.ListAllUsers(c.Request.Context())
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to fetch users")
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, users)
}

func (h *AdminHandler) ListAllPosts(c *gin.Context) {
	posts, err := h.adminService.ListAllPosts(c.Request.Context())
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to fetch posts")
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, posts)
}

func (h *AdminHandler) ListAllPromises(c *gin.Context) {
	promises, err := h.adminService.ListAllPromises(c.Request.Context())
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to fetch promises")
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, promises)
}

func (h *AdminHandler) ListAllMicrotasks(c *gin.Context) {
	microtasks, err := h.adminService.ListAllMicrotasks(c.Request.Context())
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to fetch microtasks")
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, microtasks)
}
