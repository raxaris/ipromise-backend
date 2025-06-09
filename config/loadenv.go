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
		log.Fatal("❌ Failed to load .env file")
	}

	JWTSecret = os.Getenv("JWT_SECRET")
	if JWTSecret == "" {
		log.Fatal("❌ Environment variable JWT_SECRET is not set! The application cannot run without it.")
	}

	OpenAIApiKey = os.Getenv("OPENAI_API_KEY")
	if OpenAIApiKey == "" {
		log.Fatal("❌ Environment variable OPENAI_API_KEY is not set! GPT functionality will not work.")
	}
}
