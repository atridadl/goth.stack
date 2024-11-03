package pages

import (
	"goth.stack/lib"
	"github.com/labstack/echo/v4"
)

type ToolsProps struct {
	Tools []lib.CardLink
}

func Tools(c echo.Context) error {
	tools := []lib.CardLink{
		{
			Name:        "Server Sent Events Demo",
			Description: "Server Sent Events Demo",
			Href:        "/tools/ssedemo",
			Internal:    true,
		},
		{
			Name:        "Image Resizer",
			Description: "Image Resizer Tool",
			Href:        "/tools/resize",
			Internal:    true,
		},
	}

	props := ToolsProps{
		Tools: tools,
	}

	// Specify the partials used by this page
	partials := []string{"header", "navitems", "cardlinks"}

	// Render the template
	return lib.RenderTemplate(c.Response().Writer, "base", partials, props)
}
