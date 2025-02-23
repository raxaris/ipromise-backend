package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/services"
)

// GetCurrentUserHandler – получить информацию о себе
func GetCurrentUserHandler(c *gin.Context) {
	userID, _ := uuid.Parse(c.GetString("user_id"))

	user, err := services.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUserByIDHandler – получить пользователя по ID
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

// GetUserByUsernameHandler – получить пользователя по username
func GetUserByUsernameHandler(c *gin.Context) {
	username := c.Param("username")

	user, err := services.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserHandler – обновить профиль пользователя (только username)
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

// DeleteUserHandler – удалить аккаунт
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
