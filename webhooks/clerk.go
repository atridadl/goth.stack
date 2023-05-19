package webhooks

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"atri.dad/lib"
	"github.com/labstack/echo/v4"
	svix "github.com/svix/svix-webhooks/go"
)

// Types
type ClerkEventEmail struct {
	EmailAddress string `json:"email_address"`
}

type ClerkEventData struct {
	EmailAddresses []ClerkEventEmail `json:"email_addresses,omitempty"`
	Id             string            `json:"id"`
}
type ClerkEvent struct {
	Data ClerkEventData
	Type string
}

// Event Handlers
func userCreatedHandler(event ClerkEvent) {
	welcomeEmail := `
		<h1>Thank you for making an atri.dad account!</h1>
		<h2>There are a number of apps this account give you access to!</h2>
		<br/>
		<ul>
		<li>Atash Demo: https://atash.atri.dad/</li>
		<li>Commodore: https://commodore.atri.dad/</li>
		</ul>
		`

	lib.SendEmail(event.Data.EmailAddresses[0].EmailAddress, "apps@atri.dad", "Atri's Apps", welcomeEmail, "Welcome to Atri's Apps!")
}

// Main Handler/Router
func ClerkWebhookHandler(c echo.Context) error {
	secret := os.Getenv("CLERK_WEBHOOK_SECRET")

	wh, err := svix.NewWebhook(secret)
	if err != nil {
		return c.String(http.StatusBadRequest, "Unknown Validation Error")
	}

	headers := c.Request().Header

	payload, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, "Failed to read request body!")
	}

	err = wh.Verify(payload, headers)
	if err != nil {
		return c.String(http.StatusBadRequest, "Cannot validate webhook authenticity!")
	}

	var parsed ClerkEvent

	err = json.Unmarshal(payload, &parsed)

	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid Json!")
	}

	switch parsed.Type {
	case "user.created":
		userCreatedHandler(parsed)
	}

	return c.String(http.StatusOK, "Success!")
}
