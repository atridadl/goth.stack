package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"goth.stack/api"
	"goth.stack/lib"
	"goth.stack/pages"
)

func main() {
	// Load environment variables
	godotenv.Load(".env")

	// Initialize Redis client
	lib.RedisClient = lib.NewRedisClient()

	// Test Redis connection
	_, err := lib.RedisClient.Ping(context.Background()).Result()

	// Initialize pubsub
	var pubSub lib.PubSub
	if err != nil {
		log.Printf("Failed to connect to Redis: %v", err)
		log.Println("Falling back to LocalPubSub")
		pubSub = &lib.LocalPubSub{}
	} else {
		pubSub = &lib.RedisPubSub{
			Client: lib.RedisClient,
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
	e.Static("/public", "public")

	// Page routes
	e.GET("/", pages.Home)
	e.GET("/blog", pages.Blog)
	e.GET("/post/:post", pages.Post)
	e.GET("/sse", pages.SSEDemo)

	// API Routes:
	apiGroup := e.Group("/api")
	apiGroup.GET("/ping", api.Ping)

	apiGroup.GET("/sse", func(c echo.Context) error {
		return api.SSE(c, pubSub)
	})

	apiGroup.POST("/sendsse", func(c echo.Context) error {
		return api.SSEDemoSend(c, pubSub)
	})

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
