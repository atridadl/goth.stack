package pages

import (
	"atri.dad/lib"
	"github.com/labstack/echo/v4"
)

type TestimonialsProps struct {
	Images []string
}

func Testimonials(c echo.Context) error {
	images := lib.GeneratePublicURLsFromDirectory("testimonials/")

	props := TestimonialsProps{
		Images: images,
	}

	// Specify the partials used by this page
	partials := []string{"header", "navitems"}

	// Render the template
	return lib.RenderTemplate(c.Response().Writer, "base", partials, props)
}
