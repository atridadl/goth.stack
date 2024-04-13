package api

import (
	"errors"
	"fmt"

	"atri.dad/lib"
	"atri.dad/lib/pubsub"
	"github.com/labstack/echo/v4"
)

func SSE(c echo.Context, pubSub pubsub.PubSub) error {
	if pubSub == nil {
		return errors.New("pubSub is nil")
	}

	channel := c.QueryParam("channel")
	if channel == "" {
		channel = "default"
	}

	// Use the request context, which is cancelled when the client disconnects
	ctx := c.Request().Context()

	pubsub, err := pubSub.SubscribeToChannel(channel)
	if err != nil {
		return fmt.Errorf("failed to subscribe to channel: %w", err)
	}

	lib.SetSSEHeaders(c)

	// Create a client channel and add it to the SSE server
	client := make(chan string)
	lib.SSEServer.AddClient(channel, client)
	defer lib.SSEServer.RemoveClient(channel, client)

	go lib.HandleIncomingMessages(c, pubsub, client)

	for {
		select {
		case <-ctx.Done():
			// If the client has disconnected, stop the loop
			return nil
		}
	}
}
