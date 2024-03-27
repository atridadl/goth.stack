package api

import (
	"fmt"
	"net/http"
	"strconv"

	"atri.dad/lib"
	"github.com/labstack/echo/v4"
)

func ResizeHandler(c echo.Context) error {

	// Extract file from request
	file, _, err := c.Request().FormFile("image")
	if err != nil {
		return c.String(http.StatusBadRequest, "Error getting image file")
	}
	defer file.Close()

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

	fileBlob, err := lib.ResizeImg(file, width, height)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "resized.png"))

	return c.Blob(http.StatusOK, "image/png", fileBlob)
}
