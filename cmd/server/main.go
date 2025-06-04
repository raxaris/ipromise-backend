package main

import (
	"context"
	"fmt"
	notificationcache "github.com/raxaris/ipromise-backend/internal/cache/notification"
	"github.com/raxaris/ipromise-backend/internal/mappers"
	"github.com/raxaris/ipromise-backend/internal/middleware"
	"github.com/raxaris/ipromise-backend/internal/repositories/attachment"
	"github.com/raxaris/ipromise-backend/internal/repositories/badge"
	"github.com/raxaris/ipromise-backend/internal/repositories/follower"
	"github.com/raxaris/ipromise-backend/internal/repositories/like"
	"github.com/raxaris/ipromise-backend/internal/repositories/microtask"
	"github.com/raxaris/ipromise-backend/internal/repositories/notification"
	"github.com/raxaris/ipromise-backend/internal/repositories/post"
	"github.com/raxaris/ipromise-backend/internal/repositories/prediction"
	"github.com/raxaris/ipromise-backend/internal/repositories/promise"
	"github.com/raxaris/ipromise-backend/internal/repositories/token"
	"github.com/raxaris/ipromise-backend/internal/repositories/user"
	"github.com/raxaris/ipromise-backend/internal/services"
	"github.com/raxaris/ipromise-backend/internal/watchers"
	"github.com/raxaris/ipromise-backend/internal/ws"
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
	openaiClient := config.InitOpenAIClient()
	redisClient := config.NewRedisClient()
	notificationCache := notificationcache.NewRedisNotificationCache(redisClient)

	// WebSocket Hub
	notificationHub := ws.NewNotificationHub()
	go notificationHub.Run()

	r := gin.Default()
	r.Use(gin.Recovery())

	// CORS Middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5500"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
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
	predictionRepo := prediction.NewPredictionRepository(db)
	notificationRepo := notification.NewNotificationRepository(db)

	// Маппер
	postMapper := mappers.NewPostMapper(userRepo, likeRepo, attachmentRepo, microtaskRepo, promiseRepo, postRepo, followerRepo)
	notificationMapper := mappers.NewNotificationMapper()

	// 🧠 Сервисы
	authService := services.NewAuthService(userRepo, tokenRepo)
	userService := services.NewUserService(userRepo)
	profileService := services.NewProfileService(userRepo, badgeRepo, promiseRepo, followerRepo)
	predictionService := services.NewPredictionService(predictionRepo, openaiClient)
	promiseService := services.NewPromiseService(promiseRepo, microtaskRepo, followerRepo, userRepo, postRepo, predictionService)
	microtaskService := services.NewMicrotaskService(microtaskRepo, promiseRepo)
	postService := services.NewPostService(postRepo, microtaskRepo, postMapper, followerRepo, promiseRepo, attachmentRepo, storage)
	badgeService := services.NewBadgeService(badgeRepo)
	adminService := services.NewAdminService(userRepo, postRepo, microtaskRepo, promiseRepo)
	followerService := services.NewFollowerService(followerRepo, userRepo)
	attachmentService := services.NewAttachmentService(attachmentRepo, storage)
	notificationService := services.NewNotificationService(notificationRepo, notificationCache, notificationHub, notificationMapper)
	likeService := services.NewLikeService(likeRepo, postRepo, userRepo, notificationService)

	// 🤝 Хендлеры
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	profileHandler := handlers.NewProfileHandler(profileService, attachmentService)
	promiseHandler := handlers.NewPromiseHandler(promiseService, userService)
	microtaskHandler := handlers.NewMicrotaskHandler(microtaskService)
	postHandler := handlers.NewPostHandler(postService)
	badgeHandler := handlers.NewBadgeHandler(badgeService, userService)
	adminHandler := handlers.NewAdminHandler(adminService, badgeService, userService)
	followHandler := handlers.NewFollowHandler(followerService, userService)
	attachmentHandler := handlers.NewAttachmentHandler(attachmentService)
	likeHandler := handlers.NewLikeHandler(likeService)
	predictionHandler := handlers.NewPredictionHandler(predictionService, promiseService)
	notificationHandler := handlers.NewNotificationHandler(notificationService)

	// 📌 Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Watcher
	ctx := context.Background()
	go watchers.StartBadgeWatcher(ctx, userRepo, postRepo, promiseRepo, followerRepo, badgeService)
	go watchers.StartDeadlineWatcher(ctx, promiseRepo, notificationService)

	wsHandler := handlers.NewWebSocketHandler(notificationHub, notificationService)
	wsGroup := r.Group("/ws")
	{
		wsGroup.GET("/notifications", wsHandler.ServeWS)
	}

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
		badgeGroup.GET("/:username", badgeHandler.GetBadgesByUsername)
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
		profile.PATCH("/me", profileHandler.UpdateProfileWithAvatar)
	}

	r.GET("/promises/public", promiseHandler.ListPublicPromises)

	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.GET("/users", adminHandler.ListAllUsers)
		admin.GET("/posts", adminHandler.ListAllPosts)
		admin.GET("/microtasks", adminHandler.ListAllMicrotasks)
		admin.GET("/promises", adminHandler.ListAllPromises)
		admin.POST("/admin/badges/assign", adminHandler.AssignBadge)
		admin.POST("/admin/badges", adminHandler.CreateBadge)
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

	predictions := r.Group("/promises/:id/prediction")
	predictions.Use(middleware.AuthMiddleware())
	{
		predictions.GET("", predictionHandler.GetPrediction)
		predictions.POST("", predictionHandler.GeneratePrediction)
	}

	microtasks := r.Group("/microtasks")
	microtasks.Use(middleware.AuthMiddleware())
	{
		microtasks.PATCH("/:id", microtaskHandler.UpdateMicrotask)
		microtasks.GET("/:id/posts", postHandler.ListPostsByMicrotaskID)
		microtasks.POST("/:id/posts", postHandler.CreatePost)
		microtasks.DELETE("/:id", microtaskHandler.DeleteMicrotask)
	}

	posts := r.Group("/posts")
	posts.Use(middleware.AuthMiddleware())
	{
		posts.POST("", postHandler.CreatePost)
		posts.POST("/:id/comments", postHandler.CreateReply)
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
		follow.GET("/recommended", followHandler.GetRecommendedUsers)
		follow.GET("/requests", followHandler.ListPendingRequests)
		follow.GET("/requests/sent", followHandler.ListSentRequests)
		follow.DELETE("/:username", followHandler.Unfollow)
		follow.DELETE("/requests/:username/cancel", followHandler.CancelFollowRequest)
	}

	friends := r.Group("/friends")
	friends.Use(middleware.AuthMiddleware())
	{
		friends.GET("/:username", followHandler.ListFriends)
	}

	attachments := r.Group("/attachments")
	attachments.Use(middleware.AuthMiddleware())
	{
		attachments.POST("/posts/:id", attachmentHandler.UploadAttachmentsToPost)
		attachments.GET("/posts/:id", attachmentHandler.ListAttachmentsByPostID)
		attachments.DELETE("/:id", attachmentHandler.DeleteAttachmentByID)
	}

	likes := r.Group("/posts")
	likes.Use(middleware.AuthMiddleware())
	{
		likes.POST("/:id/like", likeHandler.LikePost)
		likes.POST("/:id/unlike", likeHandler.UnlikePost)
	}

	notifications := r.Group("/notifications")
	notifications.Use(middleware.AuthMiddleware())
	{
		notifications.GET("/me", notificationHandler.ListMyNotifications)
		notifications.POST("/:id/read", notificationHandler.MarkManyAsRead)
		notifications.POST("/test-notification", notificationHandler.TestNotification)
	}

	port := "8080"
	fmt.Println("🚀 Сервер запущен на порту " + port)
	log.Fatal(r.Run(":" + port))
}
