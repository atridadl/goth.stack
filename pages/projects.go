package pages

import (
	"atri.dad/lib"
	"github.com/labstack/echo/v4"
)

type ProjectProps struct {
	Projects []lib.CardLink
}

func Projects(c echo.Context) error {
	projects := []lib.CardLink{
		{
			Name:        "Sprint Padawan",
			Description: "A scrum poker tool that helps agile teams plan their sprints in real-time.",
			Tags:        []string{"react", "remix.js", "product"},
			Href:        "https://sprintpadawan.dev",
		},
		{
			Name:        "Atash",
			Description: "The 🔥hottest🔥 full-stack Remix template!",
			Tags:        []string{"react", "remix.js", "template"},
			Href:        "http://atash.atri.dad",
		},
		{
			Name:        "Darius TS",
			Description: "🚀 A Web Application Template Powered by HTMX + Elysia + Tailwind 🚀",
			Tags:        []string{"typescript", "htmx", "template"},
			Href:        "https://ts.darius.atri.dad",
		},
		{
			Name:        "Darius Go",
			Description: "🚀 A Web Application Template Powered by HTMX + Go Fiber + Tailwind 🚀",
			Tags:        []string{"golang", "htmx", "template"},
			Href:        "https://go.darius.atri.dad",
		},
		{
			Name:        "Himbot",
			Description: "A discord bot written in Go. Loosly named after my username online (HimbothySwaggins).",
			Tags:        []string{"golang", "bot"},
			Href:        "https://github.com/atridadl/HimBot",
		},
		{
			Name:        "Commodore",
			Description: "Helpful Nightbot Helpers for Twitch",
			Tags:        []string{"react", "remix.js", "template"},
			Href:        "https://commodore.atri.dad",
		},
		{
			Name:        "loadr",
			Description: "A lightweight REST load testing tool with robust support for different verbs, token auth, and performance reports.",
			Tags:        []string{"golang", "cli"},
			Href:        "https://github.com/atridadl/loadr",
		},
		{
			Name:        "GOTH Stack",
			Description: "A template using Go + Templates + HTMX",
			Tags:        []string{"golang", "htmx", "template"},
			Href:        "https://github.com/atridadl/goth.stack",
		},
	}

	props := ProjectProps{
		Projects: projects,
	}

	// Specify the partials used by this page
	partials := []string{"header", "navitems", "cardlinks"}

	// Render the template
	return lib.RenderTemplate(c.Response().Writer, "base", partials, props)
}
