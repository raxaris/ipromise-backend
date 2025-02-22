package main

import (
	"fmt"
	"github.com/raxaris/ipromise-backend/internal/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raxaris/ipromise-backend/config"
	"github.com/raxaris/ipromise-backend/internal/handlers"
)

func main() {
	config.ConnectDB() // Инициализация базы данных
	config.LoadEnv()   // Загружаем .env переменные

	r := gin.Default()

	// Аутентификация
	auth := r.Group("/auth")
	{
		auth.POST("/signup", handlers.SignupHandler)
		auth.POST("/login", handlers.LoginHandler)
		auth.POST("/refresh", handlers.RefreshTokenHandler)
	}

	protected := r.Group("/protected")
	protected.Use(middleware.AuthMiddleware()) // Подключаем middleware для проверки токена
	{
		protected.GET("/test", func(c *gin.Context) {
			userID, exists := c.Get("user_id") // Достаем user_id из контекста
			if !exists {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось получить user_id"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Доступ разрешен", "user_id": userID})
		})
	}
	
	port := "8080"
	fmt.Println("🚀 Сервер запущен на порту " + port)
	log.Fatal(r.Run(":" + port))
}
