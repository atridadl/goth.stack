package pages

import (
	"html/template"
	"log"

	"github.com/labstack/echo/v4"
)

func SSEDemo(c echo.Context) error {
	templates := []string{
		"./pages/templates/layouts/base.html",
		"./pages/templates/partials/header.html",
		"./pages/templates/partials/navitems.html",
		"./pages/templates/ssedemo.html",
	}

	ts, err := template.ParseFiles(templates...)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	return ts.ExecuteTemplate(c.Response().Writer, "base", nil)
}
