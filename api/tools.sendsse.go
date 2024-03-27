package api

import (
	"net/http"

	"atri.dad/lib"
	"atri.dad/lib/pubsub"
	"github.com/labstack/echo/v4"
)

func SSEDemoSend(c echo.Context, pubSub pubsub.PubSub) error {
	channel := c.QueryParam("channel")
	if channel == "" {
		channel = "default"
	}

	// Get message from query parameters, form value, or request body
	message := c.QueryParam("message")
	if message == "" {
		message = c.FormValue("message")
		if message == "" {
			var body map[string]string
			if err := c.Bind(&body); err != nil {
				return err
			}
			message = body["message"]
		}
	}

	if message == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "message parameter is required"})
	}

	lib.SendSSE(c.Request().Context(), pubSub, "default", message)

	return c.JSON(http.StatusOK, map[string]string{"status": "message sent"})
}
