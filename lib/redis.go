package lib

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var RedisClient *redis.Client

func NewClient() *redis.Client {
	if RedisClient != nil {
		return RedisClient
	}

	godotenv.Load(".env")
	redis_host := os.Getenv("REDIS_HOST")
	redis_password := os.Getenv("REDIS_PASSWORD")

	log.Printf("Connecting to Redis at %s", redis_host)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redis_host,
		Password: redis_password,
		DB:       0,
	})

	return RedisClient
}

func Publish(client *redis.Client, channel string, message string) error {
	if client == nil {
		client = NewClient()
	}

	return client.Publish(ctx, channel, message).Err()
}

func Subscribe(client *redis.Client, channel string) (*redis.PubSub, string) {
	if client == nil {
		client = NewClient()
	}

	pubsub := client.Subscribe(ctx, channel)
	_, err := pubsub.Receive(ctx)
	if err != nil {
		log.Fatalf("Error receiving subscription: %v", err)
	}
	return pubsub, channel
}
