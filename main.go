package main

import (
	"io"
	"log"
	"net/http"
	"text/template"

	"goth.stack/api"
	"goth.stack/pages"

	bmw "github.com/atridadl/bmw"
	"github.com/joho/godotenv"
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

// Template Type
type Template struct {
	templates *template.Template
}

// Template Render function
func (t *Template) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

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
	pageGroup := router.NewGroup("", bunrouter.Use(bmw.NewRateLimiter(50).RateLimit))
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
