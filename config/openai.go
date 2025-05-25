package config

import (
	"github.com/raxaris/ipromise-backend/internal/ai/openai"
	"log"
)

func InitOpenAIClient() *openai.OpenAIClient {
	client := openai.NewOpenAIClient(OpenAIApiKey)
	log.Println("✅ OpenAI client initialized")
	return client
}
