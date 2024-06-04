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

	LogInfo.Printf("\nClient connected to channel %s\n", channel)
}

func (s *SSEServerType) RemoveClient(channel string, client chan string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.clients[channel], client)
	if len(s.clients[channel]) == 0 {
		delete(s.clients, channel)
	}

	LogInfo.Printf("\nClient disconnected from channel %s\n", channel)
}

func (s *SSEServerType) ClientCount(channel string) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.clients[channel])
}

func SendSSE(channel string, message string) error {
	SSEServer.mu.Lock()
	defer SSEServer.mu.Unlock()

	for client := range SSEServer.clients[channel] {
		client <- message
	}

	LogDebug.Printf("\nMessage broadcast on channel %s: %s\n", channel, message)

	return nil
}
