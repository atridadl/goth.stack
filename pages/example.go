package pages

import (
	"atri.dad/lib"
	"github.com/labstack/echo/v4"
)

type ExampleProps struct {
	ExamplePropText string
}

func Example(c echo.Context) error {
	props := ExampleProps{
		ExamplePropText: "EXAMPLE TEXT HERE",
	}

	// Specify the partials used by this page
	partials := []string{"header", "navitems"}

	// Render the template
	return lib.RenderTemplate(c.Response().Writer, "base", partials, props)
}
