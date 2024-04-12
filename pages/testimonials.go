package pages

import (
	"atri.dad/lib"
	"github.com/labstack/echo/v4"
)

type TestimonialsProps struct {
	Images []string
}

func Testimonials(c echo.Context) error {
	images := []string{
		"/public/img/testimonials/1.png",
		"/public/img/testimonials/2.png",
		"/public/img/testimonials/3.png",
	}

	props := TestimonialsProps{
		Images: images,
	}

	// Specify the partials used by this page
	partials := []string{"header", "navitems"}

	// Render the template
	return lib.RenderTemplate(c.Response().Writer, "base", partials, props)
}
