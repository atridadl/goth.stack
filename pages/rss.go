package pages

import (
	"io/fs"
	"net/http"
	"strings"
	"time"

	contentfs "atri.dad/content"
	"atri.dad/lib"
	"github.com/gorilla/feeds"
	"github.com/labstack/echo/v4"
)

func RSSFeedHandler(c echo.Context) error {
	files, err := fs.ReadDir(contentfs.FS, ".")

	protocol := "http"
	if c.Request().TLS != nil {
		protocol = "https"
	}

	feed := &feeds.Feed{
		Title: "Atridad Lahiji's Blog",
		Link:  &feeds.Link{Href: protocol + "://" + c.Request().Host + "/api/rss"},
	}

	if err != nil {
		http.Error(c.Response().Writer, "There was an issue finding posts!", http.StatusInternalServerError)
		return nil
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {

			frontMatter, err := lib.ExtractFrontMatter(file, contentfs.FS)
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
	}

	rss, _ := feed.ToRss()
	return c.Blob(http.StatusOK, "application/rss+xml", []byte(rss))
}
