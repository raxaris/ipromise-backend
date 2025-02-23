package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/raxaris/ipromise-backend/config"
	"github.com/raxaris/ipromise-backend/internal/handlers"
	"github.com/raxaris/ipromise-backend/internal/middleware"
)

func main() {
	// Инициализация БД и переменных окружения
	config.LoadEnv()
	config.ConnectDB()

	r := gin.Default()

	// 🔹 Аутентификация
	auth := r.Group("/auth")
	{
		auth.POST("/signup", handlers.SignupHandler)
		auth.POST("/login", handlers.LoginHandler)
		auth.POST("/refresh", handlers.RefreshTokenHandler)
	}

	// 🔹 Маршруты для авторизованных пользователей
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// 🔹 Профиль пользователя
		protected.GET("/user", handlers.GetCurrentUserHandler)
		protected.PUT("/user", handlers.UpdateUserHandler)
		protected.DELETE("/user", handlers.DeleteUserHandler)

		// 🔹 Обещания (только свои)
		protected.POST("/promises", handlers.CreatePromiseHandler)
		protected.GET("/promises", handlers.GetPromisesByUserIDHandler)
		protected.PUT("/promises/:id", handlers.UpdatePromiseHandler)
	}

	// 🔹 Маршруты для модераторов и админов
	moderator := r.Group("/moderation")
	moderator.Use(middleware.ModeratorMiddleware())
	{
		moderator.DELETE("/promises/:id", handlers.DeletePromiseHandler)
	}

	// 🔹 Админские маршруты (полный доступ)
	admin := r.Group("/admin")
	admin.Use(middleware.AdminMiddleware())
	{
		// Управление пользователями
		admin.GET("/users", handlers.GetAllUsersHandler)
		admin.GET("/users/:id", handlers.GetUserByIDHandler)
		admin.PUT("/users/:id", handlers.UpdateUserHandler)
		admin.DELETE("/users/:id", handlers.DeleteUserHandler)

		// Управление обещаниями
		admin.GET("/promises", handlers.GetAllPromisesHandler)
		admin.GET("/promises/:id", handlers.GetPromiseByIDHandler)
		admin.DELETE("/promises/:id", handlers.DeletePromiseHandler)
	}

	port := "8080"
	fmt.Println("🚀 Сервер запущен на порту " + port)
	log.Fatal(r.Run(":" + port))
}
