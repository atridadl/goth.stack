package adapters

import (
	"context"
	"os"

	"atri.dad/lib"
	"atri.dad/lib/pubsub"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
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
	redis_url := os.Getenv("REDIS_URL")

	opts, _ := redis.ParseURL(redis_url)

	lib.LogInfo.Printf("\n[PUBSUB/REDIS]Connecting to Redis at %s\n", opts.Addr)
	RedisClient = redis.NewClient(opts)

	return RedisClient
}

func (m *RedisPubSubMessage) ReceiveMessage(ctx context.Context) (*pubsub.Message, error) {
	msg, err := m.pubsub.ReceiveMessage(ctx)
	if err != nil {
		return nil, err
	}
	lib.LogInfo.Printf("\n[PUBSUB/REDIS] Received message: %s\n", msg.Payload)
	return &pubsub.Message{Payload: msg.Payload}, nil
}

func (ps *RedisPubSub) SubscribeToChannel(channel string) (pubsub.PubSubMessage, error) {
	pubsub := ps.Client.Subscribe(context.Background(), channel)
	_, err := pubsub.Receive(context.Background())
	if err != nil {
		return nil, err
	}
	lib.LogInfo.Printf("\n[PUBSUB/REDIS] Subscribed to channel %s\n", channel)

	return &RedisPubSubMessage{pubsub: pubsub}, nil
}

func (r *RedisPubSub) PublishToChannel(channel string, message string) error {
	err := r.Client.Publish(context.Background(), channel, message).Err()
	if err != nil {
		return err
	}
	lib.LogInfo.Printf("\n[PUBSUB/REDIS] Publishing message to channel %s: %s\n", channel, message)
	return nil
}
