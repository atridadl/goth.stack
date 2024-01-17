package pages

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"goth.stack/lib"
)

type BlogProps struct {
	Posts []lib.CardLink
}

func Blog(c echo.Context) error {
	var posts []lib.CardLink

	files, err := os.ReadDir("./content/")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "There was an finding posts!")
	}

	for _, file := range files {
		frontMatter, err := lib.ExtractFrontMatter(file, "./content/")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "There was an issue rendering the posts!")
		}

		frontMatter.Href = "post/" + strings.TrimSuffix(file.Name(), ".md")
		frontMatter.Internal = true

		posts = append(posts, frontMatter)
	}

	const layout = "January 2 2006"

	sort.Slice(posts, func(i, j int) bool {
		iDate, err := time.Parse(layout, posts[i].Date)
		if err != nil {
			log.Fatal(err)
		}

		jDate, err := time.Parse(layout, posts[j].Date)
		if err != nil {
			log.Fatal(err)
		}

		return iDate.Before(jDate)
	})

	props := BlogProps{
		Posts: posts,
	}

	templates := []string{
		"./pages/templates/layouts/base.html",
		"./pages/templates/partials/header.html",
		"./pages/templates/partials/navitems.html",
		"./pages/templates/partials/cardlinks.html",
		"./pages/templates/blog.html",
	}

	ts, err := template.ParseFiles(templates...)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	return ts.ExecuteTemplate(c.Response().Writer, "base", props)
}
