package api

import (
	"net/http"
	"os"
	"strings"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/labstack/echo/v4"
)

func Authed(c echo.Context) error {
	apiKey := os.Getenv("CLERK_SECRET_KEY")

	client, err := clerk.NewClient(apiKey)
	if err != nil {
		// handle error
		println(err.Error())
	}

	// get session token from Authorization header
	sessionToken := c.Request().Header.Get("Authorization")
	sessionToken = strings.TrimPrefix(sessionToken, "Bearer ")

	println(sessionToken)

	// verify the session
	sessClaims, err := client.VerifyToken(sessionToken)
	if err != nil {
		println(err.Error())
		return c.String(http.StatusUnauthorized, "Unauthorized!")
	}

	// get the user, and say welcome!
	user, err := client.Users().Read(sessClaims.Claims.Subject)
	if err != nil {
		panic(err)
	}

	return c.String(http.StatusOK, "Welcome "+*user.FirstName)

}
