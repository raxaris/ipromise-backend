package main

import (
	"fmt"
	"github.com/raxaris/ipromise-backend/internal/middleware"
	"github.com/raxaris/ipromise-backend/internal/repositories/badge"
	"github.com/raxaris/ipromise-backend/internal/repositories/follower"
	"github.com/raxaris/ipromise-backend/internal/repositories/microtask"
	"github.com/raxaris/ipromise-backend/internal/repositories/promise"
	"github.com/raxaris/ipromise-backend/internal/repositories/token"
	"github.com/raxaris/ipromise-backend/internal/repositories/user"
	"github.com/raxaris/ipromise-backend/internal/services"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/raxaris/ipromise-backend/config"
	"github.com/raxaris/ipromise-backend/internal/handlers"
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
	r.Use(gin.Recovery())
	// CORS Middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 🔹 Репозитории
	userRepo := user.NewUserRepository(db)
	tokenRepo := token.NewTokenRepository(db)
	promiseRepo := promise.NewPromiseRepository(db)
	badgeRepo := badge.NewBadgeRepository(db)
	followerRepo := follower.NewFollowerRepository(db)
	microtaskRepo := microtask.NewMicrotaskRepository(db)
	// TODO: добавить когда будете реализовывать
	// microtaskRepo := microtask.NewMicrotaskRepository(db)
	// postRepo := post.NewPostRepository(db)
	// likeRepo := like.NewLikeRepository(db)
	// commentRepo := comment.NewCommentRepository(db)

	// 🔹 Сервисы
	authService := services.NewAuthService(userRepo, tokenRepo)
	promiseService := services.NewPromiseService(promiseRepo, followerRepo)
	badgeService := services.NewBadgeService(badgeRepo)
	profileService := services.NewProfileService(userRepo, badgeRepo, promiseRepo, followerRepo)
	userService := services.NewUserService(userRepo)
	microtaskService := services.NewMicrotaskService(microtaskRepo, promiseRepo)
	// TODO: добавить позже
	// microtaskService := services.NewMicrotaskService(microtaskRepo)
	// postService := services.NewPostService(postRepo)
	// followerService := services.NewFollowerService(followerRepo)

	// 🔹 Хендлеры
	authHandler := handlers.NewAuthHandler(authService)
	promiseHandler := handlers.NewPromiseHandler(promiseService, userService)
	badgeHandler := handlers.NewBadgeHandler(badgeService)
	profileHandler := handlers.NewProfileHandler(profileService)
	microtaskHandler := handlers.NewMicrotaskHandler(microtaskService)
	// TODO: добавить позже
	// microtaskHandler := handlers.NewMicrotaskHandler(microtaskService)
	// postHandler := handlers.NewPostHandler(postService)
	// followerHandler := handlers.NewFollowerHandler(followerService)

	// 📌 Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := r.Group("/auth")
	{
		auth.POST("/signup", authHandler.Signup)
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh", authHandler.Refresh)
		auth.POST("/logout", authHandler.Logout)
	}

	profile := r.Group("/profile")
	profile.Use(middleware.AuthMiddleware())
	{
		profile.GET("/me", profileHandler.GetMyProfile)
		profile.PATCH("", profileHandler.UpdateProfile)
		profile.GET("/:username", profileHandler.GetPublicProfile)
	}

	promiseGroup := r.Group("/promises")
	{
		promiseGroup.POST("/", promiseHandler.CreatePromise)
		promiseGroup.GET("/profile/:username", promiseHandler.ListPublicPromises)
	}

	badgeGroup := r.Group("/badges")
	{
		badgeGroup.GET("/", badgeHandler.ListAllBadges)
		badgeGroup.GET("/me", badgeHandler.ListUserBadges)
	}

	port := "8080"
	fmt.Println("🚀 Сервер запущен на порту " + port)
	log.Fatal(r.Run(":" + port))
}
