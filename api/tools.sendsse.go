package api

import (
	"net/http"

	"atri.dad/lib"
	"github.com/labstack/echo/v4"
)

// SSEDemoSend godoc
// @Summary Send SSE message
// @Description Sends a message to a specified SSE channel
// @Tags sse,tools
// @Accept json
// @Produce json
// @Param channel query string false "Channel name"
// @Param message query string false "Message to send"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /tools/sendsse [post]
func SSEDemoSend(c echo.Context) error {
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

	// Send message
	lib.SSEServer.SendSSE(channel, message)

	return c.JSON(http.StatusOK, map[string]string{"status": "message sent"})
}
