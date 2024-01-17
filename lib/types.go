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

type Post struct {
	Content template.HTML
	Name    string
	Date    string
	Tags    []string
}
type FrontMatter struct {
	Name string
	Date string
	Tags []string
}

type PubSubMessage struct {
	Channel string `json:"channel"`
	Data    string `json:"data"`
}
