package api

import (
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"goth.stack/lib"
)

func SSE(c echo.Context) error {
	channel := c.QueryParam("channel")
	if channel == "" {
		channel = "default"
	}

	// Use the request context, which is cancelled when the client disconnects
	ctx := c.Request().Context()

	pubsub, _ := lib.Subscribe(lib.RedisClient, channel)

	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().Header().Set(echo.HeaderConnection, "keep-alive")
	c.Response().Header().Set(echo.HeaderCacheControl, "no-cache")

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
			if _, err := c.Response().Write([]byte(": keep-alive\n\n")); err != nil {
				return err
			}
			c.Response().Flush()
		default:
			// Handle incoming messages as before
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				log.Printf("Failed to receive message: %v", err)
				continue
			}

			data := fmt.Sprintf("data: %s\n\n", msg.Payload)
			if _, err := c.Response().Write([]byte(data)); err != nil {
				return err
			}

			c.Response().Flush()
		}
	}
}
