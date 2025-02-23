package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/services"
)

// GetAllUsersHandler – получение всех пользователей (только для админов)
func GetAllUsersHandler(c *gin.Context) {
	role, _ := c.Get("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "У вас нет прав для просмотра пользователей"})
		return
	}

	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения пользователей"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// UpdateUserHandler – обновление информации о пользователе
func UpdateUserHandler(c *gin.Context) {
	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Получаем ID текущего пользователя
	requesterID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка обработки user_id"})
		return
	}

	// Получаем ID пользователя, которого обновляют
	userIDParam := c.Param("id")
	userID, err := uuid.Parse(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный user_id"})
		return
	}

	isAdmin := c.GetString("role") == "admin"

	// Вызываем сервис обновления пользователя
	err = services.UpdateUser(requesterID, userID, &req, isAdmin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Данные пользователя обновлены"})
}

// DeleteUserHandler – удаление пользователя (только для админов)
func DeleteUserHandler(c *gin.Context) {
	role, _ := c.Get("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "У вас нет прав для удаления пользователей"})
		return
	}

	userID := c.Param("id")
	err := services.DeleteUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления пользователя"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пользователь удален"})
}
