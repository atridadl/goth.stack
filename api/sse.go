package api

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"goth.stack/lib"
)

// SSE godoc
// @Summary Server-Sent Events endpoint
// @Description Establishes a Server-Sent Events connection
// @Tags sse
// @Accept json
// @Produce text/event-stream
// @Param channel query string false "Channel name"
// @Success 200 {string} string "Event stream"
// @Router /sse [get]
func SSE(c echo.Context) error {
	channel := c.QueryParam("channel")
	if channel == "" {
		channel = "default"
	}

	// Use the request context, which is cancelled when the client disconnects
	ctx := c.Request().Context()

	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().Header().Set(echo.HeaderConnection, "keep-alive")
	c.Response().Header().Set(echo.HeaderCacheControl, "no-cache")

	// Create a channel to receive messages from the lib.SSEServer
	clientChan := make(chan string)

	// Add the client to the lib.SSEServer
	lib.SSEServer.AddClient(channel, clientChan)

	defer func() {
		// Remove the client from the lib.SSEServer when the connection is closed
		lib.SSEServer.RemoveClient(channel, clientChan)
	}()

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
		case msg := <-clientChan:
			// Handle incoming messages from the lib.SSEServer
			data := fmt.Sprintf("data: %s\n\n", msg)
			if _, err := c.Response().Write([]byte(data)); err != nil {
				return err
			}

			c.Response().Flush()
		}
	}
}
