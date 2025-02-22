package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var JWTSecret string

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Ошибка загрузки .env файла")
	}

	JWTSecret = os.Getenv("JWT_SECRET")
	if JWTSecret == "" {
		log.Fatal("❌ Переменная окружения JWT_SECRET не установлена! Приложение не может работать без нее.")
	}
}
