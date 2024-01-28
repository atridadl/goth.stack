package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"goth.stack/api"
	"goth.stack/pages"
)

func main() {
	// Load environment variables
	godotenv.Load(".env")

	// Initialize Echo router
	e := echo.New()

	// Generate a unique version identifier
	version := time.Now().Format(time.RFC3339)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(50)))
	// Use middleware to set ETag and Cache-Control headers
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get the path of the requested resource
			path := c.Request().URL.Path

			// If the requested resource is a CSS file
			if strings.HasSuffix(path, ".css") {
				log.Println(path)
				// Read the CSS file
				data, err := os.ReadFile(filepath.Join("public", strings.TrimPrefix(path, "/public")))
				if err != nil {
					// Log the error and return a 500 status
					log.Println(err)
					return c.NoContent(http.StatusInternalServerError)
				}

				// Compute the hash of the CSS file contents
				hash := sha256.Sum256(data)

				// Set the ETag to the hash
				c.Response().Header().Set("ETag", hex.EncodeToString(hash[:]))

				// Set the Content-Type to text/css
				c.Response().Header().Set("Content-Type", "text/css")
			} else {
				// For other resources, set the ETag to the server start time
				c.Response().Header().Set("ETag", version)
			}

			c.Response().Header().Set("Cache-Control", "public, no-cache")
			return next(c)
		}
	})

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
	apiGroup.GET("/sse", api.SSE)
	apiGroup.POST("/sendsse", api.SSEDemoSend)

	// Start server with HTTP/2 support
	s := &http.Server{
		Addr:    ":3000",
		Handler: e,
	}
	e.Logger.Fatal(e.StartServer(s))
	log.Println("Server started on port 3000")
}
