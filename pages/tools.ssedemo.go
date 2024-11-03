package pages

import (
	"goth.stack/lib"
	"github.com/labstack/echo/v4"
)

func SSEDemo(c echo.Context) error {
	// Specify the partials used by this page
	partials := []string{"header", "navitems"}

	// Render the template
	return lib.RenderTemplate(c.Response().Writer, "base", partials, nil)
}
