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

	_ "github.com/raxaris/ipromise-backend/docs"
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
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 📌 Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 🔹 Публичные маршруты
	r.GET("/users/:id", handlers.GetPublicUserHandler)       // Публичный профиль без email
	r.GET("/promises", handlers.GetAllPublicPromisesHandler) // Все обещания (без личных данных)
	r.GET("/promises/:id", handlers.GetPromiseByIDHandler)   // Одно обещание

	// 🔹 Авторизованные пользователи
	user := r.Group("/users")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/me", handlers.GetCurrentUserHandler) // Личный профиль
		user.PUT("/me", handlers.UpdateUserHandler)     // Обновление своего профиля
		user.GET("/:id", handlers.GetUserByIDHandler)   // Получение пользователя (для себя)
		user.GET("/byusername/:username", handlers.GetUserByUsernameHandler)

		// Обещания авторизованного пользователя
		user.GET("/me/promises", handlers.GetUserPromisesHandler)      // Получить свои обещания
		user.POST("/me/promises", handlers.CreatePromiseHandler)       // Создать обещание
		user.PUT("/me/promises/:id", handlers.UpdatePromiseHandler)    // Обновить обещание
		user.DELETE("/me/promises/:id", handlers.DeletePromiseHandler) // Удалить обещание
	}

	// 🔹 Админские маршруты (полный доступ)
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		// Полный доступ к пользователям
		admin.GET("/users", handlers.GetAllUsersHandler)
		admin.PUT("/users/:id", handlers.UpdateUserHandler)
		admin.DELETE("/users/:id", handlers.DeleteUserHandler)

		// Полный доступ к обещаниям
		admin.GET("/promises", handlers.GetAllPromisesHandler)
		admin.PUT("/promises/:id", handlers.UpdatePromiseHandler)
		admin.DELETE("/promises/:id", handlers.DeletePromiseHandler)
	}

	port := "8080"
	fmt.Println("🚀 Сервер запущен на порту " + port)
	log.Fatal(r.Run(":" + port))
}
