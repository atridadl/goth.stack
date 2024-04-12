package lib

import (
	"html/template"
)

type IconLink struct {
	Name string
	Href string
	Icon template.HTML
}

type CardLink struct {
	Name        string
	Href        string
	Description string
	Date        string
	Tags        []string
	Internal    bool
}

type ButtonLink struct {
	Name     string
	Href     string
	Internal bool
}
