package lib

import (
	"html/template"
	"net/http"
	"path/filepath"
	"runtime"

	templatefs "goth.stack/pages/templates"
)

func RenderTemplate(w http.ResponseWriter, layout string, partials []string, props interface{}) error {
	// Get the name of the current file
	_, filename, _, _ := runtime.Caller(1)
	page := filepath.Base(filename)
	page = page[:len(page)-len(filepath.Ext(page))] // remove the file extension

	// Build the list of templates
	templates := []string{
		"layouts/" + layout + ".html",
		page + ".html",
	}
	for _, partial := range partials {
		templates = append(templates, "partials/"+partial+".html")
	}

	// Parse the templates
	ts, err := template.ParseFS(templatefs.FS, templates...)
	if err != nil {
		LogError.Print(err.Error())
		return err
	}

	// Execute the layout template
	err = ts.ExecuteTemplate(w, layout, props)
	if err != nil {
		LogError.Print(err.Error())
		return err
	}

	return nil
}
