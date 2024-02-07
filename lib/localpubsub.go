package lib

import (
	"context"
	"sync"
	"time"

	"goth.stack/lib/pubsub"
)

type LocalPubSub struct {
	subscribers map[string][]chan pubsub.Message
	lock        sync.RWMutex
}

type LocalPubSubMessage struct {
	messages <-chan pubsub.Message
}

func (ps *LocalPubSub) SubscribeToChannel(channel string) (pubsub.PubSubMessage, error) {
	ps.lock.Lock()
	defer ps.lock.Unlock()

	if ps.subscribers == nil {
		ps.subscribers = make(map[string][]chan pubsub.Message)
	}

	ch := make(chan pubsub.Message, 100)
	ps.subscribers[channel] = append(ps.subscribers[channel], ch)

	LogInfo.Printf("[PUBSUB/LOCAL] Subscribed to channel %s", channel)

	return &LocalPubSubMessage{messages: ch}, nil
}

func (ps *LocalPubSub) PublishToChannel(channel string, message string) error {
	ps.lock.Lock()         // Changed from RLock to Lock
	defer ps.lock.Unlock() // Changed from RUnlock to Unlock

	if subscribers, ok := ps.subscribers[channel]; ok {
		LogInfo.Printf("[PUBSUB/LOCAL] Publishing message to channel %s: %s", channel, message)
		for _, ch := range subscribers {
			ch <- pubsub.Message{Payload: message}
		}
	} else {
		LogWarning.Printf("[PUBSUB/LOCAL] No subscribers for channel %s", channel)
	}

	return nil
}

func (m *LocalPubSubMessage) ReceiveMessage(ctx context.Context) (*pubsub.Message, error) {
	for {
		select {
		case <-ctx.Done():
			// The client has disconnected. Stop trying to send messages.
			return nil, ctx.Err()
		case msg := <-m.messages:
			// A message has been received. Send it to the client.
			LogInfo.Printf("[PUBSUB/LOCAL] Received message: %s", msg.Payload)
			return &msg, nil
		case <-time.After(30 * time.Second):
			// No message has been received for 30 seconds. Send a keep-alive message.
			return &pubsub.Message{Payload: "keep-alive"}, nil
		}
	}
}

func (ps *LocalPubSub) UnsubscribeFromChannel(channel string, ch <-chan pubsub.Message) {
	ps.lock.Lock()
	defer ps.lock.Unlock()

	subscribers := ps.subscribers[channel]
	for i, subscriber := range subscribers {
		if subscriber == ch {
			// Remove the subscriber from the slice
			subscribers = append(subscribers[:i], subscribers[i+1:]...)
			break
		}
	}

	if len(subscribers) == 0 {
		delete(ps.subscribers, channel)
	} else {
		ps.subscribers[channel] = subscribers
	}
}
