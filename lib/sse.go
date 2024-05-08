package lib

import (
	"context"
	"fmt"
	"log"
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
	// Create a channel to receive an error from the goroutine
	errCh := make(chan error, 1)

	// Use a goroutine to send the message asynchronously
	go func() {
		select {
		case <-ctx.Done():
			// The client has disconnected, so return an error
			errCh <- ctx.Err()
		default:
			err := messageBroker.PublishToChannel(channel, message)
			errCh <- err // Send the error to the channel
		}
	}()

	// Wait for the goroutine to finish and check for errors
	err := <-errCh
	if err != nil {
		return err
	}

	return nil
}

func SetSSEHeaders(c echo.Context) {
	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().Header().Set(echo.HeaderConnection, "keep-alive")
	c.Response().Header().Set(echo.HeaderCacheControl, "no-cache")
}

func HandleIncomingMessages(c echo.Context, pubsub pubsub.PubSubMessage, client chan string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-c.Request().Context().Done()
		cancel()
	}()

	var mutex sync.Mutex

	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				log.Printf("Failed to receive message: %v", err)
				return
			}

			data := fmt.Sprintf("data: %s\n\n", msg.Payload)

			mutex.Lock()
			defer mutex.Unlock()

			if c.Response().Writer != nil {
				_, err = c.Response().Write([]byte(data))
				if err != nil {
					log.Printf("Failed to write message: %v", err)
					return
				}

				flusher, ok := c.Response().Writer.(http.Flusher)
				if ok {
					flusher.Flush()
				} else {
					log.Println("Failed to flush: ResponseWriter does not implement http.Flusher")
				}
			} else {
				log.Println("Failed to write: ResponseWriter is nil")
				return
			}
		}
	}
}
