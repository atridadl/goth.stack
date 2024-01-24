package pages

import (
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/uptrace/bunrouter"
	"goth.stack/lib"
)

type BlogProps struct {
	Posts []lib.CardLink
}

func Blog(w http.ResponseWriter, req bunrouter.Request) error {
	var posts []lib.CardLink

	files, err := os.ReadDir("./content/")
	if err != nil {
		http.Error(w, "There was an issue finding posts!", http.StatusInternalServerError)
		return nil
	}

	for _, file := range files {
		frontMatter, err := lib.ExtractFrontMatter(file, "./content/")
		if err != nil {
			http.Error(w, "There was an issue rendering the posts!", http.StatusInternalServerError)
			return nil
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

	// Specify the partials used by this page
	partials := []string{"header", "navitems", "cardlinks"}

	// Render the template
	return lib.RenderTemplate(w, "base", partials, props)
}
