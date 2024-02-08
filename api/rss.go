package api

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/feeds"
	"github.com/labstack/echo/v4"
	"goth.stack/lib"
)

func RSSFeedHandler(c echo.Context) error {
	files, err := os.ReadDir("./content/")

	protocol := "http"
	if c.Request().TLS != nil {
		protocol = "https"
	}

	feed := &feeds.Feed{
		Title: "GOTH Stack Demo Blog",
		Link:  &feeds.Link{Href: protocol + "://" + c.Request().Host + "/api/rss"},
	}

	if err != nil {
		http.Error(c.Response().Writer, "There was an issue finding posts!", http.StatusInternalServerError)
		return nil
	}

	for _, file := range files {
		frontMatter, err := lib.ExtractFrontMatter(file, "./content/")
		if err != nil {
			http.Error(c.Response().Writer, "There was an issue rendering the posts!", http.StatusInternalServerError)
			return nil
		}

		date, _ := time.Parse("January 2 2006", frontMatter.Date)

		feed.Add(&feeds.Item{
			Title:   frontMatter.Name,
			Link:    &feeds.Link{Href: protocol + "://" + c.Request().Host + "/post/" + strings.TrimSuffix(file.Name(), ".md")},
			Created: date,
		})
	}

	rss, _ := feed.ToRss()
	return c.Blob(http.StatusOK, "application/rss+xml", []byte(rss))
}
