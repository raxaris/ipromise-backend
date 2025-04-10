package handlers

//import (
//	"github.com/raxaris/ipromise-backend/config"
//	"github.com/raxaris/ipromise-backend/internal/models"
//	"net/http"
//
//	"github.com/gin-gonic/gin"
//	"github.com/google/uuid"
//	"github.com/raxaris/ipromise-backend/internal/dto"
//	"github.com/raxaris/ipromise-backend/internal/services"
//)
//
//func GetCurrentUserHandler(c *gin.Context) {
//	userID, _ := uuid.Parse(c.GetString("user_id"))
//
//	user, err := services.GetUserByID(userID)
//	if err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
//		return
//	}
//
//	c.JSON(http.StatusOK, user)
//}
//
//func GetPublicUserHandler(c *gin.Context) {
//	username := c.Param("username")
//
//	var user models.User
//	if err := config.DB.Select("id, username, created_at").
//		Where("username = ?", username).
//		First(&user).Error; err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
//		return
//	}
//
//	c.JSON(http.StatusOK, user)
//}
//
//func GetAllUsersHandler(c *gin.Context) {
//	users, err := services.GetAllUsers()
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения пользователей"})
//		return
//	}
//
//	c.JSON(http.StatusOK, users)
//}
//
//func GetUserByIDHandler(c *gin.Context) {
//	idStr := c.Param("id")
//
//	userID, err := uuid.Parse(idStr)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
//		return
//	}
//
//	user, err := services.GetUserByID(userID)
//	if err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
//		return
//	}
//
//	c.JSON(http.StatusOK, user)
//}
//
//func GetUserByUsernameHandler(c *gin.Context) {
//	username := c.Param("username")
//
//	user, err := services.GetUserByUsername(username)
//	if err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
//		return
//	}
//
//	c.JSON(http.StatusOK, user)
//}
//
//func UpdateUserHandler(c *gin.Context) {
//	var req dto.UpdateUserRequest
//	if err := c.ShouldBindJSON(&req); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	userID, _ := uuid.Parse(c.GetString("user_id"))
//
//	err := services.UpdateUser(userID, userID, &req, false)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "Данные пользователя обновлены"})
//}
//
//func DeleteUserHandler(c *gin.Context) {
//	userID, _ := uuid.Parse(c.GetString("user_id"))
//
//	// Проверяем, существует ли пользователь
//	_, err := services.GetUserByID(userID)
//	if err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
//		return
//	}
//
//	// Удаляем пользователя
//	err = services.DeleteUser(userID)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления пользователя"})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "Аккаунт удалён"})
//}
