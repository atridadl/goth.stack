package pages

import (
	"bytes"
	"html/template"
	"io/fs"
	"net/http"

	contentfs "goth.stack/content"
	"goth.stack/lib"
	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/labstack/echo/v4"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"gopkg.in/yaml.v2"
)

type PostProps struct {
	Content template.HTML
	Name    string
	Date    string
	Tags    []string
}

func Post(c echo.Context) error {
	postName := c.Param("post")

	filePath := postName + ".md"

	md, err := fs.ReadFile(contentfs.FS, filePath)
	if err != nil {
		println(err.Error())
		http.Error(c.Response().Writer, "This post does not exist!", http.StatusNotFound)
		return nil
	}

	frontmatterBytes, content, err := lib.SplitFrontmatter(md)
	if err != nil {
		http.Error(c.Response().Writer, "There was an issue rendering this post!", http.StatusInternalServerError)
		return nil
	}

	var frontmatter lib.FrontMatter
	if err := yaml.Unmarshal(frontmatterBytes, &frontmatter); err != nil {
		http.Error(c.Response().Writer, "There was an issue rendering this post!", http.StatusInternalServerError)
		return nil
	}

	var buf bytes.Buffer
	markdown := goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithStyle("fruity"),
				highlighting.WithFormatOptions(
					chromahtml.WithLineNumbers(true),
				),
			),
		),
	)

	if err := markdown.Convert(content, &buf); err != nil {
		http.Error(c.Response().Writer, "There was an issue rendering this post!", http.StatusInternalServerError)
		return nil
	}

	props := PostProps{
		Content: template.HTML(buf.String()),
		Name:    frontmatter.Name,
		Date:    frontmatter.Date,
		Tags:    frontmatter.Tags,
	}

	// Specify the partials used by this page
	partials := []string{"header", "navitems"}

	// Render the template
	return lib.RenderTemplate(c.Response().Writer, "post", partials, props)
}
