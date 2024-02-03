package lib

import "sync"

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

func SendSSE(channel string, message string) error {
	// Create a channel to receive an error from the goroutine
	errCh := make(chan error, 1)

	// Use a goroutine to send the message asynchronously
	go func() {
		err := Publish(RedisClient, channel, message)
		errCh <- err // Send the error to the channel
	}()

	// Wait for the goroutine to finish and check for errors
	err := <-errCh
	if err != nil {
		return err
	}

	return nil
}
