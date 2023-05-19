package api

import (
	"net/http"
	"os"

	"atri.dad/lib"
	"github.com/labstack/echo/v4"
)

func NowPlayingHandler(c echo.Context) error {
	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	refreshToken := os.Getenv("SPOTIFY_REFRESH_TOKEN")

	playing, err := lib.GetCurrentlyPlayingTrack(clientID, clientSecret, refreshToken)
	if err != nil {
		http.Error(c.Response().Writer, err.Error(), http.StatusInternalServerError)
		return err
	}

	if playing.Item != nil && playing.Playing {
		return c.String(http.StatusOK, `<div class="indicator-item badge badge-success"><a _="on mouseover put '🎧 Listening to `+playing.Item.Name+" by "+playing.Item.Artists[0].Name+` 🎧' into my.textContent on mouseout put '🎧' into my.textContent" href="`+playing.Item.ExternalURLs["spotify"]+`" rel="noreferrer" target="_blank">🎧</a></div>`)
	} else {
		return c.String(http.StatusOK, "")
	}
}
