package pages

import (
	"io/fs"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	contentfs "goth.stack/content"
	"goth.stack/lib"
)

type BlogProps struct {
	Posts []lib.CardLink
}

func Blog(c echo.Context) error {
	var posts []lib.CardLink

	files, err := fs.ReadDir(contentfs.FS, ".")
	if err != nil {
		log.Println(err)
		http.Error(c.Response().Writer, "There was an issue finding posts!", http.StatusInternalServerError)
		return nil
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {
			frontMatter, err := lib.ExtractFrontMatter(file, contentfs.FS)
			if err != nil {
				log.Println(err)
				http.Error(c.Response().Writer, "There was an issue rendering the posts!", http.StatusInternalServerError)
				return nil
			}

			frontMatter.Href = "post/" + strings.TrimSuffix(file.Name(), ".md")
			frontMatter.Internal = true

			posts = append(posts, frontMatter)
		}
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

	// Specify the partials used by this page
	partials := []string{"header", "navitems", "cardlinks"}

	// Render the template
	return lib.RenderTemplate(c.Response().Writer, "base", partials, props)
}
