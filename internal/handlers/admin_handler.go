package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/services"
	"github.com/raxaris/ipromise-backend/internal/utils"
	"net/http"
)

type AdminHandler struct {
	adminService services.AdminService
	badgeService services.BadgeService
	userService  services.UserService
}

func NewAdminHandler(adminService services.AdminService, badgeService services.BadgeService, userService services.UserService) *AdminHandler {
	return &AdminHandler{
		adminService: adminService,
		badgeService: badgeService,
		userService:  userService,
	}
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

func (h *AdminHandler) CreateBadge(c *gin.Context) {
	var req dto.CreateBadgeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Неверный формат запроса")
		return
	}

	badge := &models.Badge{
		Code:        req.Code,
		Title:       req.Title,
		Description: req.Description,
		IconURL:     req.IconURL,
	}

	err := h.badgeService.CreateBadge(c.Request.Context(), badge)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}
	utils.RespondWithSuccess(c, http.StatusCreated, gin.H{"message": "Бейдж создан"})
}

func (h *AdminHandler) AssignBadge(c *gin.Context) {
	var req dto.AssignBadgeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Неверный формат запроса")
		return
	}

	user, err := h.userService.GetUserByUsername(c.Request.Context(), req.Username)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	err = h.badgeService.AssignBadgeByCode(c.Request.Context(), user.ID, req.BadgeCode)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"message": "Бейдж назначен"})
}
