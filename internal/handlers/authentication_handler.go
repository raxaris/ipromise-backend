package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/config"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/services"
)

// SignupHandler регистрирует нового пользователя
// @Summary Регистрация нового пользователя
// @Description Создаёт нового пользователя по email и паролю
// @Tags auth
// @Accept json
// @Produce json
// @Param input body dto.SignupRequest true "Данные для регистрации"
// @Success 201 {object} map[string]string "message: Пользователь успешно зарегистрирован"
// @Failure 400 {object} map[string]string "error: Ошибка валидации"
// @Failure 409 {object} map[string]string "error: Email уже используется"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /auth/signup [post]

func SignupHandler(c *gin.Context) {
	var req dto.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверяем, совпадает ли пароль и подтверждение пароля
	if req.Password != req.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пароли не совпадают"})
		return
	}

	// Проверяем, существует ли уже email или username
	var existingUser models.User
	if err := config.DB.Where("email = ? OR username = ?", req.Email, req.Username).First(&existingUser).Error; err == nil {
		if existingUser.Email == req.Email {
			c.JSON(http.StatusConflict, gin.H{"error": "Email уже используется"})
		} else {
			c.JSON(http.StatusConflict, gin.H{"error": "Username уже используется"})
		}
		return
	}
	// Создаем пользователя **с паролем**
	user := models.User{
		ID:       uuid.New(),
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password, // Добавляем пароль
	}

	// Теперь хешируем его внутри структуры
	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка хеширования пароля"})
		return
	}

	// Сохраняем в БД
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания пользователя"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Пользователь успешно зарегистрирован"})
}

// LoginHandler аутентифицирует пользователя
// @Summary Авторизация пользователя
// @Description Логин по email и паролю, выдаёт JWT токены
// @Tags auth
// @Accept json
// @Produce json
// @Param input body dto.LoginRequest true "Данные для входа"
// @Success 200 {object} map[string]string "access_token: токен, refresh_token: токен"
// @Failure 400 {object} map[string]string "error: Ошибка валидации"
// @Failure 401 {object} map[string]string "error: Неверный email или пароль"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /auth/login [post]
func LoginHandler(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
		return
	}

	if !user.CheckPassword(req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
		return
	}

	accessToken, err := services.GenerateAccessToken(user.ID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка генерации Access-токена"})
		return
	}

	refreshToken, err := services.GenerateRefreshToken(user.ID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка генерации Refresh-токена"})
		return
	}

	// Сохраняем Refresh Token в БД
	refreshTokenEntry := models.RefreshToken{
		ID:        uuid.New(),
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	if err := config.DB.Create(&refreshTokenEntry).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сохранения Refresh-токена"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken, "refresh_token": refreshToken})
}

// RefreshTokenHandler обновляет Access Token
// @Summary Обновление Access Token
// @Description Использует Refresh Token для выдачи нового Access Token
// @Tags auth
// @Accept json
// @Produce json
// @Param input body dto.RefreshTokenRequest true "Refresh Token"
// @Success 200 {object} map[string]string "access_token: новый access-токен"
// @Failure 400 {object} map[string]string "error: Ошибка валидации"
// @Failure 401 {object} map[string]string "error: Недействительный Refresh-токен"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /auth/refresh [post]
func RefreshTokenHandler(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверяем `refresh_token` через сервис
	refreshToken, err := services.ValidateRefreshTokenFromDB(config.DB, req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Недействительный или истёкший Refresh-токен"})
		return
	}

	// Генерируем новый `access_token`
	newAccessToken, err := services.GenerateAccessToken(refreshToken.UserID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка генерации нового Access-токена"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})
}
