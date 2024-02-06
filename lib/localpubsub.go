package lib

import (
	"context"
	"log"
	"sync"
	"time"
)

type LocalPubSub struct {
	subscribers map[string][]chan Message
	lock        sync.RWMutex
}

type LocalPubSubMessage struct {
	messages <-chan Message
}

func (ps *LocalPubSub) SubscribeToChannel(channel string) (PubSubMessage, error) {
	ps.lock.Lock()
	defer ps.lock.Unlock()

	if ps.subscribers == nil {
		ps.subscribers = make(map[string][]chan Message)
	}

	ch := make(chan Message, 100)
	ps.subscribers[channel] = append(ps.subscribers[channel], ch)

	log.Printf("Subscribed to channel %s", channel)

	return &LocalPubSubMessage{messages: ch}, nil
}

func (ps *LocalPubSub) PublishToChannel(channel string, message string) error {
	ps.lock.RLock()
	defer ps.lock.RUnlock()

	if subscribers, ok := ps.subscribers[channel]; ok {
		log.Printf("Publishing message to channel %s: %s", channel, message)
		for _, ch := range subscribers {
			ch <- Message{Payload: message}
		}
	} else {
		log.Printf("No subscribers for channel %s", channel)
	}

	return nil
}

func (m *LocalPubSubMessage) ReceiveMessage(ctx context.Context) (*Message, error) {
	for {
		select {
		case <-ctx.Done():
			// The client has disconnected. Stop trying to send messages.
			return nil, ctx.Err()
		case msg := <-m.messages:
			// A message has been received. Send it to the client.
			log.Printf("Received message: %s", msg.Payload)
			return &msg, nil
		case <-time.After(30 * time.Second):
			// No message has been received for 30 seconds. Send a keep-alive message.
			return &Message{Payload: "keep-alive"}, nil
		}
	}
}

func (ps *LocalPubSub) UnsubscribeFromChannel(channel string, ch <-chan Message) {
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
