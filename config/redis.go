package config

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("❌ Redis connection failed: %v", err)
	}

	log.Println("✅ Redis connected")
	return rdb
}
