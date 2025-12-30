package redis_handler

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	Ctx    = context.Background()
	Client *redis.Client
)

// Инициализация Redis
func InitRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	if err := Client.Ping(Ctx).Err(); err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
}
