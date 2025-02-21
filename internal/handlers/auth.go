package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/config"
	"github.com/raxaris/ipromise-backend/internal/models"
	"net/http"
)

type SignupRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func SignupHandler(c *gin.Context) {
	var req SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		ID:       uuid.New(),
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка хеширования пароля"})
		return
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания пользователя"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Пользователь успешно зарегистрирован"})
}
