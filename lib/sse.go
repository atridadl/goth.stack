package lib

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"atri.dad/lib/pubsub"
	"github.com/labstack/echo/v4"
)

type SSEServerType struct {
	clients map[string]map[chan string]bool
	mu      sync.Mutex
}

var SSEServer *SSEServerType

func init() {
	SSEServer = &SSEServerType{
		clients: make(map[string]map[chan string]bool),
	}
}

func NewSSEServer() *SSEServerType {
	return &SSEServerType{
		clients: make(map[string]map[chan string]bool),
	}
}

func (s *SSEServerType) AddClient(channel string, client chan string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.clients[channel]; !ok {
		s.clients[channel] = make(map[chan string]bool)
	}
	s.clients[channel][client] = true
}

func (s *SSEServerType) RemoveClient(channel string, client chan string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.clients[channel], client)
	if len(s.clients[channel]) == 0 {
		delete(s.clients, channel)
	}
}

func (s *SSEServerType) ClientCount(channel string) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.clients[channel])
}

func SendSSE(ctx context.Context, messageBroker pubsub.PubSub, channel string, message string) error {
	LogInfo.Printf("Sending SSE message to channel %s", channel)

	errCh := make(chan error, 1)

	go func() {
		select {
		case <-ctx.Done():
			errCh <- ctx.Err()
		default:
			err := messageBroker.PublishToChannel(channel, message)
			errCh <- err
		}
	}()

	err := <-errCh
	if err != nil {
		LogError.Printf("Error sending SSE message: %v", err)

		return err
	}

	LogSuccess.Printf("SSE message sent successfully")

	return nil
}

func SetSSEHeaders(c echo.Context) {
	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().Header().Set(echo.HeaderConnection, "keep-alive")
	c.Response().Header().Set(echo.HeaderCacheControl, "no-cache")
	c.Response().Header().Set("X-Accel-Buffering", "no")
}

func HandleIncomingMessages(c echo.Context, pubsub pubsub.PubSubMessage, client chan string) {
	if c.Response().Writer == nil {
		LogError.Printf("Cannot handle incoming messages: ResponseWriter is nil")
		return
	}

	var mutex sync.Mutex

	for {
		msg, err := pubsub.ReceiveMessage(c.Request().Context())
		if err != nil {
			if err == context.Canceled {
				// The client has disconnected. Stop trying to send messages.
				return
			}
			LogError.Printf("Failed to receive message: %v", err)
			return
		}

		data := fmt.Sprintf("data: %s\n\n", msg.Payload)

		mutex.Lock()
		if c.Response().Writer != nil {
			_, err := c.Response().Write([]byte(data))
			if err != nil {
				LogError.Printf("Failed to write message: %v", err)
				mutex.Unlock()
				return
			}

			flusher, ok := c.Response().Writer.(http.Flusher)
			if ok {
				flusher.Flush()
			} else {
				LogError.Println("Failed to flush: ResponseWriter does not implement http.Flusher")
			}
		} else {
			LogError.Println("Failed to write: ResponseWriter is nil")
		}
		mutex.Unlock()
	}
}
