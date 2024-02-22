package lib

import (
	"bytes"
	"errors"
	"image"
	"image/png"
	"io"
	"mime/multipart"

	"github.com/anthonynsimon/bild/transform"
)

func ResizeImg(file multipart.File, width int, height int) ([]byte, error) {
	// Read file content
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.New("Error reading image file")

	}

	// Decode image
	img, _, err := image.Decode(bytes.NewReader(fileContent))
	if err != nil {
		return nil, errors.New("Error decoding image")
	}

	// Resize the image
	resizedImg := transform.Resize(img, width, height, transform.Linear)

	// Encode the resized image as PNG
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, resizedImg); err != nil {
		return nil, errors.New("Error encoding image to PNG")
	}

	// Return the resized image as response
	return buf.Bytes(), nil
}
