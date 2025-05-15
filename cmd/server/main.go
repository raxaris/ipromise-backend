package main

import (
	"fmt"
	"github.com/raxaris/ipromise-backend/internal/mappers"
	"github.com/raxaris/ipromise-backend/internal/middleware"
	"github.com/raxaris/ipromise-backend/internal/repositories/attachment"
	"github.com/raxaris/ipromise-backend/internal/repositories/badge"
	"github.com/raxaris/ipromise-backend/internal/repositories/follower"
	"github.com/raxaris/ipromise-backend/internal/repositories/like"
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
	storage := config.CreateStorage()

	r := gin.Default()
	r.Use(gin.Recovery())
	// CORS Middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
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
	likeRepo := like.NewLikeRepository(db)
	attachmentRepo := attachment.NewAttachmentRepository(db)

	// Маппер
	postMapper := mappers.NewPostMapper(
		userRepo,
		likeRepo,
		attachmentRepo,
		microtaskRepo,
		promiseRepo,
		postRepo,
	)

	// 🧠 Сервисы
	authService := services.NewAuthService(userRepo, tokenRepo)
	userService := services.NewUserService(userRepo)
	profileService := services.NewProfileService(userRepo, badgeRepo, promiseRepo, followerRepo)
	promiseService := services.NewPromiseService(promiseRepo, microtaskRepo, followerRepo, userRepo, postRepo)
	microtaskService := services.NewMicrotaskService(microtaskRepo, promiseRepo)
	postService := services.NewPostService(postRepo, microtaskRepo, postMapper, followerRepo, promiseRepo)
	badgeService := services.NewBadgeService(badgeRepo)
	adminService := services.NewAdminService(userRepo, postRepo, microtaskRepo, promiseRepo)
	followerService := services.NewFollowerService(followerRepo, userRepo)
	attachmentService := services.NewAttachmentService(attachmentRepo, storage)

	// 🤝 Хендлеры
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	profileHandler := handlers.NewProfileHandler(profileService)
	promiseHandler := handlers.NewPromiseHandler(promiseService, userService)
	microtaskHandler := handlers.NewMicrotaskHandler(microtaskService)
	postHandler := handlers.NewPostHandler(postService)
	badgeHandler := handlers.NewBadgeHandler(badgeService)
	adminHandler := handlers.NewAdminHandler(adminService)
	followHandler := handlers.NewFollowHandler(followerService, userService)
	attachmentHandler := handlers.NewAttachmentHandler(attachmentService)
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
	badgeGroup.Use(middleware.AuthMiddleware())
	{
		badgeGroup.GET("/", badgeHandler.ListAllBadges)
		badgeGroup.GET("/me", badgeHandler.ListUserBadges)
	}

	users := r.Group("/users")
	users.Use(middleware.AuthMiddleware())
	{
		users.GET("/:username", userHandler.GetUserByUsername)
		users.PATCH("/:id", userHandler.UpdateUser)
	}

	profile := r.Group("/profile")
	profile.Use(middleware.AuthMiddleware())
	{
		profile.GET("/me", profileHandler.GetMyProfile)
		profile.GET("/:username", profileHandler.GetPublicProfile)
		profile.PATCH("/me", profileHandler.UpdateProfile)
	}

	r.GET("/promises/public", promiseHandler.ListPublicPromises)

	admin := r.Group("/admin")
	{
		admin.GET("/users", adminHandler.ListAllUsers)
		admin.GET("/posts", adminHandler.ListAllPosts)
		admin.GET("/microtasks", adminHandler.ListAllMicrotasks)
		admin.GET("/promises", adminHandler.ListAllPromises)
	}

	promises := r.Group("/promises")
	promises.Use(middleware.AuthMiddleware())
	{
		promises.POST("", promiseHandler.CreatePromise)
		promises.POST("/full", promiseHandler.CreatePromiseWithMicrotasks)
		promises.GET("/:id", promiseHandler.GetPromiseByID)
		promises.PATCH("/:id", promiseHandler.UpdatePromise)
		promises.DELETE("/:id", promiseHandler.DeletePromise)
		promises.GET("/feed", promiseHandler.ListFeedPromises)
		promises.POST("/:id/microtasks", microtaskHandler.CreateMicrotask)
		promises.PATCH("/:id/microtasks/reorder", microtaskHandler.ReorderMicrotasks)
		promises.GET("/user/:username", promiseHandler.ListProfilePromises)
		promises.GET("/user/:username/progress", promiseHandler.GetUserPromisesWithProgress)
	}

	microtasks := r.Group("/microtasks")
	microtasks.Use(middleware.AuthMiddleware())
	{
		microtasks.PATCH("/:id", microtaskHandler.UpdateMicrotask)
		microtasks.GET("/:id/posts", postHandler.ListPostsByMicrotaskID)
		microtasks.POST("/:id/posts", postHandler.CreatePost)
	}

	posts := r.Group("/posts")
	posts.Use(middleware.AuthMiddleware())
	{
		posts.PATCH("/:id", postHandler.UpdatePost)                // Обновить пост
		posts.DELETE("/:id", postHandler.DeletePost)               // Удалить пост
		posts.GET("/:id/replies", postHandler.ListReplies)         // Получить комментарии
		posts.GET("/:id/full", postHandler.GetFullPost)            // Получить дерево поста
		posts.GET("/public", postHandler.ListPublicPostsLite)      // Публичные посты (лайт)
		posts.GET("/feed", postHandler.ListFeedPostsLite)          // Лента подписок (лайт)
		posts.GET("/public/tree", postHandler.ListPublicPostsTree) // Публичные посты (дерево)
		posts.GET("/feed/tree", postHandler.ListFeedPostsTree)     // Лента подписок (дерево)
		posts.GET("/user/:username", postHandler.ListUserPostsTree)
	}

	follow := r.Group("/follow")
	follow.Use(middleware.AuthMiddleware())
	{
		follow.POST("/:username", followHandler.RequestFollow)
		follow.POST("/:username/accept", followHandler.AcceptFollowRequest)
		follow.POST("/:username/decline", followHandler.DeclineFollowRequest)
		follow.GET("/requests", followHandler.ListPendingRequests)
		follow.DELETE("/:username", followHandler.Unfollow)
	}

	friends := r.Group("/friends")
	friends.Use(middleware.AuthMiddleware())
	{
		friends.GET("/:username", followHandler.ListFriends)
	}

	attachments := r.Group("/attachments")
	{
		attachments.POST("/posts/:id", attachmentHandler.UploadAttachmentsToPost)
		attachments.GET("/posts/:id", attachmentHandler.ListAttachmentsByPostID)
		attachments.DELETE("/:id", attachmentHandler.DeleteAttachmentByID)
	}

	port := "8080"
	fmt.Println("🚀 Сервер запущен на порту " + port)
	log.Fatal(r.Run(":" + port))
}
