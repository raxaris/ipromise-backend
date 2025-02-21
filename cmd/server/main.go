package main

import (
	"fmt"
	"github.com/raxaris/ipromise-backend/config"
	"github.com/raxaris/ipromise-backend/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB() // Инициализация базы данных
	r := gin.Default()
	r.POST("/auth/signup", handlers.SignupHandler)

	port := "8080"
	fmt.Println("🚀 Сервер запущен на порту " + port)
	log.Fatal(r.Run(":" + port))

}
