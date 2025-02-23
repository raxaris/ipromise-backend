package main

import (
	"fmt"
	"log"

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
		protected.GET("/user", handlers.GetCurrentUserHandler)
		protected.PUT("/user", handlers.UpdateUserHandler)
		protected.DELETE("/user", handlers.DeleteUserHandler)

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
	//admin.Use(middleware.AdminMiddleware())
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
