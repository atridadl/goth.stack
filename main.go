package main

import (
	"embed"
	"flag"
	"fmt"
	"net/http"

	"goth.stack/api"
	"goth.stack/lib"
	"goth.stack/pages"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "goth.stack/docs"
)

//go:embed public/*
var PublicFS embed.FS

// @title GOTH Stack API
// @version 1.0
// @description This is the API for GOTH Stack - Go + HTMX + Tailwind
// @host localhost:3000
// @BasePath /api
func main() {
	// Load environment variables
	godotenv.Load(".env")

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
	e.GET("/posts", pages.Posts)
	e.GET("/posts/:post", pages.Post)
	e.GET("/tools", pages.Tools)
	e.GET("/tools/resize", pages.Resize)
	e.GET("/tools/ssedemo", pages.SSEDemo)

	// API Routes:
	apiGroup := e.Group("/api")
	// Swagger endpoint
	apiGroup.GET("/swagger/*", echoSwagger.WrapHandler)
	apiGroup.GET("/ping", api.Ping)
	apiGroup.GET("/rss", api.RSSFeedHandler)
	apiGroup.GET("/post/copy", api.PostCopy)

	apiGroup.GET("/sse", func(c echo.Context) error {
		return api.SSE(c)
	})

	apiGroup.POST("/tools/sendsse", func(c echo.Context) error {
		return api.SSEDemoSend(c)
	})

	apiGroup.POST("/tools/resize", api.ResizeHandler)

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
	lib.LogSuccess.Println("Server started on port", *port)
}
