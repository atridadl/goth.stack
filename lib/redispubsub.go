package lib

import (
	"context"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"goth.stack/lib/pubsub"
)

var RedisClient *redis.Client

type RedisPubSubMessage struct {
	pubsub *redis.PubSub
}

// RedisPubSub is a Redis implementation of the PubSub interface.
type RedisPubSub struct {
	Client *redis.Client
}

func NewRedisClient() *redis.Client {
	if RedisClient != nil {
		return RedisClient
	}

	godotenv.Load(".env")
	redis_host := os.Getenv("REDIS_HOST")
	redis_password := os.Getenv("REDIS_PASSWORD")

	LogInfo.Printf("Connecting to Redis at %s", redis_host)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redis_host,
		Password: redis_password,
		DB:       0,
	})

	return RedisClient
}

func (m *RedisPubSubMessage) ReceiveMessage(ctx context.Context) (*pubsub.Message, error) {
	msg, err := m.pubsub.ReceiveMessage(ctx)
	if err != nil {
		return nil, err
	}
	LogInfo.Printf("[PUBSUB/REDIS] Received message: %s", msg.Payload)
	return &pubsub.Message{Payload: msg.Payload}, nil
}

func (ps *RedisPubSub) SubscribeToChannel(channel string) (pubsub.PubSubMessage, error) {
	pubsub := ps.Client.Subscribe(context.Background(), channel)
	_, err := pubsub.Receive(context.Background())
	if err != nil {
		return nil, err
	}
	LogInfo.Printf("[PUBSUB/REDIS] Subscribed to channel %s", channel)

	return &RedisPubSubMessage{pubsub: pubsub}, nil
}

func (r *RedisPubSub) PublishToChannel(channel string, message string) error {
	err := r.Client.Publish(context.Background(), channel, message).Err()
	if err != nil {
		return err
	}
	color.Cyan("[PUBSUB/REDIS] Publishing message to channel %s: %s", channel, message)
	return nil
}
