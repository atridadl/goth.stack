package pages

import (
	"atri.dad/lib"
	"github.com/labstack/echo/v4"
)

type PubProps struct {
	Pubs []lib.CardLink
}

func Pubs(c echo.Context) error {
	pubs := []lib.CardLink{}

	props := PubProps{
		Pubs: pubs,
	}

	// Specify the partials used by this page
	partials := []string{"header", "navitems", "cardlinks"}

	// Render the template
	return lib.RenderTemplate(c.Response().Writer, "base", partials, props)
}
