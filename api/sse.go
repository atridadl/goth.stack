package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/uptrace/bunrouter"
	"goth.stack/lib"
)

func SSE(w http.ResponseWriter, req bunrouter.Request) error {
	queryParams := req.URL.Query()
	channel := queryParams.Get("channel")
	if channel == "" {
		channel = "default"
	}

	// Use the request context, which is cancelled when the client disconnects
	ctx := req.Context()

	pubsub, _ := lib.Subscribe(lib.RedisClient, channel)

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Cache-Control", "no-cache")

	// Create a ticker that fires every 15 seconds
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			// If the client has disconnected, stop the loop
			return nil
		case <-ticker.C:
			// Every 30 seconds, send a comment to keep the connection alive
			if _, err := w.Write([]byte(": keep-alive\n\n")); err != nil {
				return err
			}
			w.(http.Flusher).Flush()
		default:
			// Handle incoming messages as before
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				log.Printf("Failed to receive message: %v", err)
				continue
			}

			data := fmt.Sprintf("data: %s\n\n", msg.Payload)
			if _, err := w.Write([]byte(data)); err != nil {
				return err
			}

			w.(http.Flusher).Flush()
		}
	}
}
