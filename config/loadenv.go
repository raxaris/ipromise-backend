package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var JWTSecret string
var OpenAIApiKey string

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Ошибка загрузки .env файла")
	}

	JWTSecret = os.Getenv("JWT_SECRET")
	if JWTSecret == "" {
		log.Fatal("❌ Переменная окружения JWT_SECRET не установлена! Приложение не может работать без нее.")
	}

	OpenAIApiKey = os.Getenv("OPENAI_API_KEY")
	if OpenAIApiKey == "" {
		log.Fatal("❌ Переменная окружения OPENAI_API_KEY не установлена! GPT-функционал не будет работать.")
	}
}
