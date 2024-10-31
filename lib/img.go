package lib

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"mime/multipart"

	"golang.org/x/image/draw"
)

func ResizeImg(file multipart.File, width int, height int) ([]byte, error) {
    // Read and decode image
    img, _, err := image.Decode(file)
    if err != nil {
        return nil, fmt.Errorf("decode error: %w", err)
    }

    // Create new RGBA image
    dst := image.NewRGBA(image.Rect(0, 0, width, height))

    // Resize using high-quality interpolation
    draw.CatmullRom.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)

    // Encode to PNG
    buf := new(bytes.Buffer)
    if err := png.Encode(buf, dst); err != nil {
        return nil, fmt.Errorf("encode error: %w", err)
    }

    return buf.Bytes(), nil
}
