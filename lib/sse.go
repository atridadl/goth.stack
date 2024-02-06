package lib

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

type SSEServerType struct {
	clients map[string]map[chan string]bool
	mu      sync.Mutex
}

var SSEServer *SSEServerType
var mutex = &sync.Mutex{}

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

func SendSSE(ctx context.Context, messageBroker PubSub, channel string, message string) error {
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

func CreateTickerAndKeepAlive(c echo.Context, duration time.Duration) *time.Ticker {
	ticker := time.NewTicker(duration)
	go func() {
		for range ticker.C {
			if _, err := c.Response().Write([]byte(": keep-alive\n\n")); err != nil {
				log.Printf("Failed to write keep-alive: %v", err)
			}
			c.Response().Flush()
		}
	}()
	return ticker
}

func HandleIncomingMessages(c echo.Context, pubsub PubSubMessage, client chan string) {
	for {
		select {
		case <-c.Request().Context().Done():
			// The client has disconnected. Stop trying to send messages.
			return
		default:
			// The client is still connected. Continue processing messages.
			msg, err := pubsub.ReceiveMessage(c.Request().Context())
			if err != nil {
				log.Printf("Failed to receive message: %v", err)
				continue
			}

			data := fmt.Sprintf("data: %s\n\n", msg.Payload)

			mutex.Lock()
			_, err = c.Response().Write([]byte(data))
			mutex.Unlock()

			if err != nil {
				log.Printf("Failed to write message: %v", err)
				return // Stop processing if an error occurs
			}

			// Check if the ResponseWriter is nil before trying to flush it
			if c.Response().Writer != nil {
				// Check if the ResponseWriter implements http.Flusher before calling Flush
				flusher, ok := c.Response().Writer.(http.Flusher)
				if ok {
					flusher.Flush()
				} else {
					log.Println("Failed to flush: ResponseWriter does not implement http.Flusher")
				}
			} else {
				log.Println("Failed to flush: ResponseWriter is nil")
			}
		}
	}
}
