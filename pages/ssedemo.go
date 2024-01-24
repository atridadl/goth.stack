package pages

import (
	"net/http"

	"github.com/uptrace/bunrouter"
	"goth.stack/lib"
)

func SSEDemo(w http.ResponseWriter, req bunrouter.Request) error {
	// Specify the partials used by this page
	partials := []string{"header", "navitems"}

	// Render the template
	return lib.RenderTemplate(w, "base", partials, nil)
}
