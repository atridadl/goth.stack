package pages

import (
	"atri.dad/lib"
	"github.com/labstack/echo/v4"
)

type ToolsProps struct {
	Tools []lib.CardLink
}

func Tools(c echo.Context) error {
	tools := []lib.CardLink{
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
