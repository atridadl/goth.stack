package adapters

import (
	"context"
	"sync"
	"time"

	"atri.dad/lib"
	"atri.dad/lib/pubsub"
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

	lib.LogInfo.Printf("[PUBSUB/LOCAL] Subscribed to channel %s\n", channel)

	return &LocalPubSubMessage{messages: ch}, nil
}

func (ps *LocalPubSub) PublishToChannel(channel string, message string) error {
	subscribers, ok := ps.subscribers[channel]
	if !ok {
		lib.LogWarning.Printf("\n[PUBSUB/LOCAL] No subscribers for channel %s\n", channel)
		return nil
	}

	ps.lock.Lock()
	defer ps.lock.Unlock()

	lib.LogInfo.Printf("\n[PUBSUB/LOCAL] Publishing message to channel %s: %s\n", channel, message)
	for _, ch := range subscribers {
		ch <- pubsub.Message{Payload: message}
	}

	return nil
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

func (m *LocalPubSubMessage) ReceiveMessage(ctx context.Context) (*pubsub.Message, error) {
	for {
		select {
		case <-ctx.Done():
			// The client has disconnected. Stop trying to send messages.
			return nil, ctx.Err()
		case msg := <-m.messages:
			// A message has been received. Send it to the client.
			lib.LogInfo.Printf("\n[PUBSUB/LOCAL] Received message: %s\n", msg.Payload)
			return &msg, nil
		case <-time.After(30 * time.Second):
			// No message has been received for 30 seconds. Send a keep-alive message.
			return &pubsub.Message{Payload: "keep-alive"}, nil
		}
	}
}
