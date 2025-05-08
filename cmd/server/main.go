package main

import (
	"fmt"
	"github.com/raxaris/ipromise-backend/internal/repositories/badge"
	"github.com/raxaris/ipromise-backend/internal/repositories/follower"
	"github.com/raxaris/ipromise-backend/internal/repositories/microtask"
	"github.com/raxaris/ipromise-backend/internal/repositories/post"
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

	// 📦 Репозитории
	userRepo := user.NewUserRepository(db)
	tokenRepo := token.NewTokenRepository(db)
	promiseRepo := promise.NewPromiseRepository(db)
	followerRepo := follower.NewFollowerRepository(db)
	microtaskRepo := microtask.NewMicrotaskRepository(db)
	postRepo := post.NewPostRepository(db)
	badgeRepo := badge.NewBadgeRepository(db)

	// 🧠 Сервисы
	authService := services.NewAuthService(userRepo, tokenRepo)
	userService := services.NewUserService(userRepo)
	profileService := services.NewProfileService(userRepo, badgeRepo, promiseRepo, followerRepo)
	promiseService := services.NewPromiseService(promiseRepo, followerRepo)
	microtaskService := services.NewMicrotaskService(microtaskRepo, promiseRepo)
	postService := services.NewPostService(postRepo)
	badgeService := services.NewBadgeService(badgeRepo)

	// 🤝 Хендлеры
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	profileHandler := handlers.NewProfileHandler(profileService)
	promiseHandler := handlers.NewPromiseHandler(promiseService, userService)
	microtaskHandler := handlers.NewMicrotaskHandler(microtaskService)
	postHandler := handlers.NewPostHandler(postService)
	badgeHandler := handlers.NewBadgeHandler(badgeService)

	// 📌 Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := r.Group("/auth")
	{
		auth.POST("/signup", authHandler.Signup)
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh", authHandler.Refresh)
		auth.POST("/logout", authHandler.Logout)
	}

	badgeGroup := r.Group("/badges")
	{
		badgeGroup.GET("/", badgeHandler.ListAllBadges)
		badgeGroup.GET("/me", badgeHandler.ListUserBadges)
	}

	users := r.Group("/users")
	{
		users.GET("/:username", userHandler.GetUserByUsername)
		users.PATCH("/:id", userHandler.UpdateUser)
	}

	profile := r.Group("/profile")
	{
		profile.GET("/me", profileHandler.GetMyProfile)
		profile.GET("/:username", profileHandler.GetPublicProfile)
		profile.PATCH("/me", profileHandler.UpdateProfile)
	}

	promises := r.Group("/promises")
	{
		promises.POST("", promiseHandler.CreatePromise)
		promises.GET("/:id", promiseHandler.GetPromiseByID)
		promises.PATCH("/:id", promiseHandler.UpdatePromise)
		promises.DELETE("/:id", promiseHandler.DeletePromise)
		promises.GET("/feed", promiseHandler.ListFeedPromises)
		promises.GET("/public", promiseHandler.ListPublicPromises)
		promises.GET("/user/:username", promiseHandler.ListProfilePromises)

		promises.POST("/:promise_id/microtasks", microtaskHandler.CreateMicrotask)
		promises.PATCH("/:promise_id/microtasks/reorder", microtaskHandler.ReorderMicrotasks)
	}

	microtasks := r.Group("/microtasks")
	{
		microtasks.PATCH("/:id", microtaskHandler.UpdateMicrotask)
		// TODO: добавить GET /:id/posts при необходимости
	}

	posts := r.Group("/posts")
	{
		posts.POST("/:microtask_id/posts", postHandler.CreatePost) // или перенести внутрь microtasks
		posts.PATCH("/:id", postHandler.UpdatePost)
		posts.DELETE("/:id", postHandler.DeletePost)
		posts.GET("/:id/replies", postHandler.ListReplies)
	}

	badges := r.Group("/badges")
	{
		badges.GET("", badgeHandler.ListAllBadges)
		badges.GET("/me", badgeHandler.ListUserBadges)
	}
	port := "8080"
	fmt.Println("🚀 Сервер запущен на порту " + port)
	log.Fatal(r.Run(":" + port))
}
