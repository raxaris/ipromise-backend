package main

import (
	"fmt"
	"github.com/raxaris/ipromise-backend/internal/repositories"
	"github.com/raxaris/ipromise-backend/internal/services"
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

	db := config.ConnectDB()
	config.InitGlobalDB(db)

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

	// 🔹 Маршруты для аутентификации
	userRepo := repositories.NewUserRepository(db)
	tokenRepo := repositories.NewTokenRepository(db)
	promiseRepo := repositories.NewPromiseRepositoryV1(db)
	authService := services.NewAuthService(userRepo, tokenRepo)
	authHandler := handlers.NewAuthHandler(authService)
	promiseService := services.NewPromiseService(promiseRepo, userRepo)
	promiseHandler := handlers.NewPromiseHandler(promiseService)

	// 📌 Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/users/:id/promises", promiseHandler.GetPublicByUserID) // публичные
	r.GET("/promises", promiseHandler.GetPublic)                   // лента
	r.GET("/promises/:id", promiseHandler.GetByID)                 // по id
	r.GET("/promises/:id/children", promiseHandler.GetChildren)    // прогресс

	auth := r.Group("/auth")
	{
		auth.POST("/signup", authHandler.Signup)
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh", authHandler.Refresh)
		auth.POST("/logout", authHandler.Logout)
	}

	user := r.Group("/profile")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/promises", promiseHandler.GetMy)
		user.POST("/promises", promiseHandler.Create)
		user.PUT("/promises/:id", promiseHandler.Update)
		user.DELETE("/promises/:id", promiseHandler.Delete)
	}

	// 🔹 Админские маршруты (полный доступ)
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		// Админ видит все обещания, может редактировать и удалять
		admin.GET("/promises", promiseHandler.GetAllForAdmin)
		admin.GET("/promises/:id", promiseHandler.GetByID)              // конкретное обещание
		admin.GET("/promises/:id/children", promiseHandler.GetChildren) // прогресс
		admin.PUT("/promises/:id", promiseHandler.Update)
		admin.DELETE("/promises/:id", promiseHandler.Delete)
	}

	port := "8080"
	fmt.Println("🚀 Сервер запущен на порту " + port)
	log.Fatal(r.Run(":" + port))
}
