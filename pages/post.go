package pages

import (
	"bytes"
	"html/template"
	"net/http"
	"os"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/labstack/echo/v4"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"gopkg.in/yaml.v2"
	"goth.stack/lib"
)

type PostProps struct {
	Content template.HTML
	Name    string
	Date    string
	Tags    []string
}

func Post(c echo.Context) error {
	postName := c.ParamValues()[0]

	filePath := "content/" + postName + ".md"

	md, err := os.ReadFile(filePath)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "This post does not exist!")
	}

	frontmatterBytes, content, err := lib.SplitFrontmatter(md)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "There was an issue rendering this post!")
	}

	var frontmatter lib.FrontMatter
	if err := yaml.Unmarshal(frontmatterBytes, &frontmatter); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "There was an issue rendering this post!")
	}

	var buf bytes.Buffer
	markdown := goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithStyle("dracula"),
				highlighting.WithFormatOptions(
					chromahtml.WithLineNumbers(true),
				),
			),
		),
	)

	if err := markdown.Convert(content, &buf); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "There was an issue rendering this post!")
	}

	props := PostProps{
		Content: template.HTML(buf.String()),
		Name:    frontmatter.Name,
		Date:    frontmatter.Date,
		Tags:    frontmatter.Tags,
	}

	templates := []string{
		"./pages/templates/layouts/post.html",
		"./pages/templates/partials/header.html",
		"./pages/templates/partials/navitems.html",
		"./pages/templates/post.html",
	}

	ts, err := template.ParseFiles(templates...)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "There was an issue rendering this post!")
	}

	return ts.ExecuteTemplate(c.Response().Writer, "post", props)
}
