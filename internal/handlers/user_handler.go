package handlers

import (
	"github.com/raxaris/ipromise-backend/config"
	"github.com/raxaris/ipromise-backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/services"
)

// GetCurrentUserHandler получает информацию о текущем пользователе
// @Summary Получение информации о себе
// @Description Возвращает данные текущего пользователя
// @Tags users
// @Security BearerAuth
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string "error: Пользователь не найден"
// @Router /user/me [get]
func GetCurrentUserHandler(c *gin.Context) {
	userID, _ := uuid.Parse(c.GetString("user_id"))

	user, err := services.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetPublicUserHandler(c *gin.Context) {
	username := c.Param("username")

	var user models.User
	if err := config.DB.Select("id, username, created_at").
		Where("username = ?", username).
		First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetAllUsersHandler получает список всех пользователей (только для админов)
// @Summary Получение всех пользователей
// @Description Возвращает список всех зарегистрированных пользователей
// @Tags admin
// @Security BearerAuth
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /admin/users/ [get]
func GetAllUsersHandler(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения пользователей"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByIDHandler получает пользователя по ID
// @Summary Получение пользователя по ID
// @Description Возвращает данные пользователя по ID
// @Tags users
// @Param id path string true "ID пользователя"
// @Security BearerAuth
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string "error: Неверный формат ID"
// @Failure 404 {object} map[string]string "error: Пользователь не найден"
// @Router /users/{id} [get]
func GetUserByIDHandler(c *gin.Context) {
	idStr := c.Param("id")

	userID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	user, err := services.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUserByUsernameHandler получает пользователя по username
// @Summary Получение пользователя по username
// @Description Возвращает данные пользователя по username
// @Tags users
// @Param username path string true "Имя пользователя"
// @Security BearerAuth
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string "error: Пользователь не найден"
// @Router /users/username/{username} [get]
func GetUserByUsernameHandler(c *gin.Context) {
	username := c.Param("username")

	user, err := services.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserHandler обновляет профиль пользователя (только username)
// @Summary Обновление профиля пользователя
// @Description Позволяет изменить username (доступно только самому пользователю)
// @Tags users
// @Security BearerAuth
// @Param input body dto.UpdateUserRequest true "Данные для обновления"
// @Success 200 {object} map[string]string "message: Данные пользователя обновлены"
// @Failure 400 {object} map[string]string "error: Ошибка валидации"
// @Failure 403 {object} map[string]string "error: Нет прав на редактирование"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /user/me [put]
func UpdateUserHandler(c *gin.Context) {
	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := uuid.Parse(c.GetString("user_id"))

	err := services.UpdateUser(userID, userID, &req, false)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Данные пользователя обновлены"})
}

// DeleteUserHandler удаляет аккаунт пользователя
// @Summary Удаление аккаунта
// @Description Удаляет аккаунт текущего пользователя
// @Tags users
// @Security BearerAuth
// @Success 200 {object} map[string]string "message: Аккаунт удалён"
// @Failure 404 {object} map[string]string "error: Пользователь не найден"
// @Failure 500 {object} map[string]string "error: Ошибка удаления"
// @Router /user [delete]
func DeleteUserHandler(c *gin.Context) {
	userID, _ := uuid.Parse(c.GetString("user_id"))

	// Проверяем, существует ли пользователь
	_, err := services.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	// Удаляем пользователя
	err = services.DeleteUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления пользователя"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Аккаунт удалён"})
}
