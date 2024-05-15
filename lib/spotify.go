package lib

import (
	"context"
	"os"
	"strings"
	"sync"

	"atri.dad/lib/pubsub"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

var (
	spotifyOAuth2Endpoint = oauth2.Endpoint{
		TokenURL: "https://accounts.spotify.com/api/token",
		AuthURL:  "https://accounts.spotify.com/authorize",
	}
	config *oauth2.Config
	once   sync.Once
)

func NowPlayingTextFilter(s string) string {
	s = strings.Replace(s, "'", "&#39;", -1)
	s = strings.Replace(s, "\"", "&quot;", -1)
	return s
}

func GetOAuth2Config(clientID string, clientSecret string) *oauth2.Config {
	once.Do(func() {
		config = &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Scopes:       []string{spotify.ScopeUserReadCurrentlyPlaying},
			Endpoint:     spotifyOAuth2Endpoint,
		}
	})
	return config
}

func GetCurrentlyPlayingTrack(clientID string, clientSecret string, refreshToken string) (*spotify.CurrentlyPlaying, error) {
	// OAuth2 config
	config := GetOAuth2Config(clientID, clientSecret)

	// Token source
	tokenSource := config.TokenSource(context.Background(), &oauth2.Token{RefreshToken: refreshToken})

	// Get new token
	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, err
	}

	// Create new client
	client := spotify.Authenticator{}.NewClient(newToken)

	// Get currently playing track
	playing, err := client.PlayerCurrentlyPlaying()
	if err != nil {
		return nil, err
	}

	return playing, nil
}

func CurrentlyPlayingTrackSSE(ctx context.Context, pubSub pubsub.PubSub) error {
	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	refreshToken := os.Getenv("SPOTIFY_REFRESH_TOKEN")

	playing, err := GetCurrentlyPlayingTrack(clientID, clientSecret, refreshToken)
	if err != nil {
		return err
	}

	if playing.Item != nil && playing.Playing {
		songName := NowPlayingTextFilter(playing.Item.Name)
		artistName := NowPlayingTextFilter(playing.Item.Artists[0].Name)

		return SendSSE(ctx, pubSub, "spotify", `<div class="indicator-item badge badge-success"><a _='on mouseover put "🔥 Listening to `+songName+" by "+artistName+` 🔥" into my.textContent on mouseout put "🔥" into my.textContent' href="`+playing.Item.ExternalURLs["spotify"]+`" rel="noreferrer" target="_blank">🔥</a></div>`)
	} else {
		SendSSE(ctx, pubSub, "spotify", "<span></span>")
	}

	return nil
}
