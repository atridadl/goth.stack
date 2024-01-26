package pages

import (
	"github.com/labstack/echo/v4"
	"goth.stack/lib"
)

func SSEDemo(c echo.Context) error {
	// Specify the partials used by this page
	partials := []string{"header", "navitems"}

	// Render the template
	return lib.RenderTemplate(c.Response().Writer, "base", partials, nil)
}
