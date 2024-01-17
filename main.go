package main

import (
	"io"
	"text/template"

	"goth.stack/api"
	"goth.stack/pages"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Template Type
type Template struct {
	templates *template.Template
}

// Template Render function
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	godotenv.Load(".env")

	// Initialize router
	e := echo.New()

	// Middlewares
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(69)))

	// Template Parsing
	t := &Template{
		templates: template.Must(template.ParseGlob("pages/**/*.html")),
	}

	e.Renderer = t

	// // Static server
	e.Static("/public", "public")

	// Page routes
	e.GET("/", pages.Home)
	e.GET("/blog", pages.Blog)
	e.GET("/post/:post", pages.Post)
	e.GET("/sse", pages.SSEDemo)

	// API Routes:
	e.GET("/api/ping", api.Ping)
	e.GET("/api/ssedemo", api.SSEDemo)
	e.POST("/api/sendsse", api.SSEDemoSend)

	e.Logger.Fatal(e.Start(":3000"))
}
