package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/raxaris/ipromise-backend/config"
	"github.com/raxaris/ipromise-backend/internal/handlers"
	"github.com/raxaris/ipromise-backend/internal/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/raxaris/ipromise-backend/docs" // 🚀 Правильный импорт (после генерации)
)

// @title iPromise API
// @version 1.0
// @description API для отслеживания обещаний пользователей.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	config.LoadEnv()
	config.ConnectDB()

	r := gin.Default()
	// CORS Middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Разрешаем все домены (можно ограничить)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// 📌 Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
		protected.GET("/user/me", handlers.GetCurrentUserHandler)
		protected.PUT("/user/me", handlers.UpdateUserHandler)

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
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.GET("/users", handlers.GetAllUsersHandler)
		admin.GET("/users/:id", handlers.GetUserByIDHandler)
		admin.GET("/users/byusername/:username", handlers.GetUserByUsernameHandler)
		admin.PUT("/users/:id", handlers.UpdateUserHandler)
		admin.DELETE("/users/:id", handlers.DeleteUserHandler)

		admin.GET("/promises", handlers.GetAllPromisesHandler)
		admin.GET("/promises/:id", handlers.GetPromisesByIDHandler)
		admin.DELETE("/promises/:id", handlers.DeletePromiseHandler)
	}

	port := "8080"
	fmt.Println("🚀 Сервер запущен на порту " + port)
	log.Fatal(r.Run(":" + port))
}
