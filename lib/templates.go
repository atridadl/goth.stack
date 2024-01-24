package lib

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
)

func RenderTemplate(w http.ResponseWriter, layout string, partials []string, props interface{}) error {
	// Get the name of the current file
	_, filename, _, _ := runtime.Caller(1)
	page := filepath.Base(filename)
	page = page[:len(page)-len(filepath.Ext(page))] // remove the file extension

	// Build the list of templates
	templates := []string{
		"./pages/templates/layouts/" + layout + ".html",
		"./pages/templates/" + page + ".html",
	}
	for _, partial := range partials {
		templates = append(templates, "./pages/templates/partials/"+partial+".html")
	}

	// Parse the templates
	ts, err := template.ParseFiles(templates...)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	// Execute the layout template
	err = ts.ExecuteTemplate(w, layout, props)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	return nil
}
