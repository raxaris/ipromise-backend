package handlers

import (
	"github.com/raxaris/ipromise-backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/services"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Signup godoc
// @Summary Регистрация нового пользователя
// @Description Создаёт нового пользователя по email, имени и паролю
// @Tags auth
// @Accept json
// @Produce json
// @Param input body dto.SignupRequest true "Данные для регистрации пользователя"
// @Success 201 {object} map[string]string "message: Пользователь зарегистрирован"
// @Failure 400 {object} map[string]string "error: Неверные данные"
// @Failure 409 {object} map[string]string "error: Email или имя пользователя уже занято"
// @Router /auth/signup [post]
func (h *AuthHandler) Signup(c *gin.Context) {
	var req dto.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.authService.Signup(c.Request.Context(), req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusCreated, "User signed up successfully")
}

// Login godoc
// @Summary Авторизация пользователя
// @Description Логин по email и паролю, выдаёт JWT токены
// @Tags auth
// @Accept json
// @Produce json
// @Param input body dto.LoginRequest true "Данные для входа"
// @Success 200 {object} map[string]string "access_token: токен, refresh_token: токен"
// @Failure 400 {object} map[string]string "error: Ошибка валидации"
// @Failure 401 {object} map[string]string "error: Неверный email или пароль"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	access, refresh, err := h.authService.Login(c.Request.Context(), req)
	if err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.SetCookie("access_token", access, 24*3600, "/", "", false, true)

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{
		"access_token":  access,
		"refresh_token": refresh,
	})
}

// Refresh godoc
// @Summary Обновление Access Token
// @Description Использует Refresh Token для выдачи нового Access Token
// @Tags auth
// @Accept json
// @Produce json
// @Param input body dto.RefreshTokenRequest true "Refresh Token"
// @Success 200 {object} map[string]string "access_token: новый access-токен"
// @Failure 400 {object} map[string]string "error: Ошибка валидации"
// @Failure 401 {object} map[string]string "error: Недействительный Refresh-токен"
// @Router /auth/refresh [post]
func (h *AuthHandler) Refresh(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	newAccess, err := h.authService.Refresh(c.Request.Context(), req.RefreshToken)
	if err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"access_token": newAccess})
}

// Logout godoc
// @Summary Выход из системы
// @Description Удаляет Refresh Token из хранилища
// @Tags auth
// @Accept json
// @Produce json
// @Param input body dto.RefreshTokenRequest true "Refresh Token"
// @Success 200 {object} map[string]string "message: Вы успешно вышли из системы"
// @Failure 400 {object} map[string]string "error: Ошибка валидации"
// @Failure 401 {object} map[string]string "error: Недействительный Refresh-токен"
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.authService.Logout(c.Request.Context(), req.RefreshToken); err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "You successfully logged out")
}
