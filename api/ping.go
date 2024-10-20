package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Ping godoc
// @Summary Ping the server
// @Description Get a pong response
// @Tags ping
// @Accept json
// @Produce json
// @Success 200 {string} string "Pong!"
// @Router /ping [get]
func Ping(c echo.Context) error {
	return c.String(http.StatusOK, "Pong!")
}
