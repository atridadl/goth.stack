package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"atri.dad/api"
	"atri.dad/lib"
	"atri.dad/lib/pubsub"
	"atri.dad/lib/pubsub/adapters"
	"atri.dad/pages"
	"atri.dad/webhooks"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed public/*
var PublicFS embed.FS

func main() {
	// Load environment variables
	godotenv.Load(".env")

	// Initialize Redis client
	adapters.RedisClient = adapters.NewRedisClient()

	// Test Redis connection
	_, err := adapters.RedisClient.Ping(context.Background()).Result()

	// Initialize pubsub
	var pubSub pubsub.PubSub
	if err != nil {
		lib.LogWarning.Printf("\n[PUBSUB/INIT] Failed to connect to Redis: %v\n", err)
		lib.LogWarning.Printf("\n[PUBSUB/INIT] Falling back to LocalPubSub\n")
		pubSub = &adapters.LocalPubSub{}
	} else {
		pubSub = &adapters.RedisPubSub{
			Client: adapters.RedisClient,
		}
	}

	// Initialize Echo router
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(50)))

	// Static server
	fs := http.FS(PublicFS)
	e.GET("/public/*", echo.WrapHandler(http.FileServer(fs)))

	// Page routes
	e.GET("/", pages.Home)
	e.GET("/projects", pages.Projects)
	e.GET("/talks", pages.Talks)
	e.GET("/blog", pages.Blog)
	e.GET("/post/:post", pages.Post)
	e.GET("/sse", pages.SSEDemo)
	e.GET("/rss", pages.RSSFeedHandler)

	// API Routes:
	apiGroup := e.Group("/api")
	apiGroup.GET("/ping", api.Ping)
	apiGroup.GET("/authed/ping", api.Authed)
	apiGroup.POST("/pay", api.Pay)
	apiGroup.GET("/post/copy", api.PostCopy)

	apiGroup.GET("/sse", func(c echo.Context) error {
		return api.SSE(c, pubSub)
	})

	apiGroup.POST("/sendsse", func(c echo.Context) error {
		return api.SSEDemoSend(c, pubSub)
	})

	apiGroup.GET("/nowplaying", api.NowPlayingHandler)

	// Webhook Routes:
	webhookGroup := e.Group("/webhook")
	webhookGroup.POST("/clerk", webhooks.ClerkWebhookHandler)

	// Spotify Polling
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			// Check if there are any clients connected to the "spotify" channel
			if lib.SSEServer.ClientCount("spotify") > 0 {
				// Get the currently playing track
				err := lib.CurrentlyPlayingTrackSSE(context.Background(), pubSub)
				if err != nil {
					// Handle error
					continue
				}
			}
		}
	}()

	// Parse command-line arguments for IP and port
	ip := flag.String("ip", "", "IP address to bind the server to")
	port := flag.String("port", "3000", "Port to bind the server to")
	flag.Parse()

	// Start server with HTTP/2 support
	s := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", *ip, *port),
		Handler: e,
	}
	e.Logger.Fatal(e.StartServer(s))
	log.Println("Server started on port", *port)
}
