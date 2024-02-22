package lib

import (
	"bytes"
	"image"
	"image/png"
	"io"
	"net/http"
	"strconv"

	"github.com/anthonynsimon/bild/transform"
	"github.com/labstack/echo/v4"
)

func ResizeImg(c echo.Context) error {
	// Extract file from request
	file, _, err := c.Request().FormFile("image")
	if err != nil {
		return c.String(http.StatusBadRequest, "Error getting image file")
	}
	defer file.Close()

	// Read file content
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error reading image file")
	}

	// Decode image
	img, _, err := image.Decode(bytes.NewReader(fileContent))
	if err != nil {
		return c.String(http.StatusBadRequest, "Error decoding image")
	}

	// Get dimensions from form data parameters
	widthStr := c.FormValue("width")
	heightStr := c.FormValue("height")

	// Validate and convert dimensions to integers
	width, err := strconv.Atoi(widthStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid width parameter")
	}

	height, err := strconv.Atoi(heightStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid height parameter")
	}

	// Resize the image
	resizedImg := transform.Resize(img, width, height, transform.Linear)

	// Encode the resized image as PNG
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, resizedImg); err != nil {
		return c.String(http.StatusInternalServerError, "Error encoding image to PNG")
	}

	// Return the resized image as response
	return c.Blob(http.StatusOK, "image/png", buf.Bytes())
}
