package pubsub

import "context"

type Message struct {
	Payload string
}

type PubSubMessage interface {
	ReceiveMessage(ctx context.Context) (*Message, error)
}

type PubSub interface {
	SubscribeToChannel(channel string) (PubSubMessage, error)
	PublishToChannel(channel string, message string) error
}
