package main

import (
	"log"
	"net/http"

	"goth.stack/api"
	"goth.stack/pages"

	"github.com/atridadl/bmw"
	"github.com/joho/godotenv"
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

func main() {
	godotenv.Load(".env")

	// Initialize router
	router := bunrouter.New(
		bunrouter.Use(reqlog.NewMiddleware(), bmw.RequestID, bmw.SecureHeaders),
	)

	// Static server
	fs := http.FileServer(http.Dir("public"))

	router.GET("/public/*filepath", func(w http.ResponseWriter, req bunrouter.Request) error {
		http.StripPrefix("/public", fs).ServeHTTP(w, req.Request)
		return nil
	})

	// Page routes
	rateLimiter := bmw.NewRateLimiter(50)
	pageGroup := router.NewGroup("", bunrouter.Use(rateLimiter.RateLimit))
	pageGroup.GET("/", pages.Home)
	pageGroup.GET("/blog", pages.Blog)
	pageGroup.GET("/post/:post", pages.Post)
	pageGroup.GET("/sse", pages.SSEDemo)

	// API Routes:
	apiGroup := router.NewGroup("/api")
	apiGroup.GET("/ping", api.Ping)

	apiGroup.GET("/sse", api.SSE)
	apiGroup.POST("/sendsse", api.SSEDemoSend)

	log.Fatal(http.ListenAndServe(":3000", router))
}
