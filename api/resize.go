package api

import (
	"atri.dad/lib"
	"github.com/labstack/echo/v4"
)

func ResizeHandler(c echo.Context) error {
	return lib.ResizeImg(c)
}
