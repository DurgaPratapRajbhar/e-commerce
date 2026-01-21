package processor

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
)

type ImageProcessor struct {
	ResizeWidth     uint
	ResizeHeight    uint
	ThumbnailWidth  uint
	ThumbnailHeight uint
}

func NewImageProcessor(resizeWidth, resizeHeight, thumbnailWidth, thumbnailHeight uint) *ImageProcessor {
	return &ImageProcessor{
		ResizeWidth:     resizeWidth,
		ResizeHeight:    resizeHeight,
		ThumbnailWidth:  thumbnailWidth,
		ThumbnailHeight: thumbnailHeight,
	}
}

func (p *ImageProcessor) ProcessImage(filePath string) error {
	// Open the image file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open image file: %w", err)
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	// Create resized version
	resizedImg := resize.Resize(p.ResizeWidth, p.ResizeHeight, img, resize.Lanczos3)

	// Create path for resized image
	dir := filepath.Dir(filePath)
	base := filepath.Base(filePath)
	ext := filepath.Ext(base)
	name := base[:len(base)-len(ext)]
	resizedPath := filepath.Join(dir, fmt.Sprintf("%s_medium%s", name, ext))

	// Save resized image
	out, err := os.Create(resizedPath)
	if err != nil {
		return fmt.Errorf("failed to create resized image file: %w", err)
	}
	defer out.Close()

	// Encode based on original format
	switch ext {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(out, resizedImg, nil)
	case ".png":
		err = png.Encode(out, resizedImg)
	default:
		// Default to JPEG
		err = jpeg.Encode(out, resizedImg, nil)
	}

	if err != nil {
		return fmt.Errorf("failed to encode resized image: %w", err)
	}

	// Create thumbnail
	thumbnailImg := resize.Resize(p.ThumbnailWidth, p.ThumbnailHeight, img, resize.Lanczos3)

	// Create path for thumbnail
	thumbnailPath := filepath.Join(dir, fmt.Sprintf("%s_thumbnail%s", name, ext))

	// Save thumbnail
	outThumb, err := os.Create(thumbnailPath)
	if err != nil {
		return fmt.Errorf("failed to create thumbnail file: %w", err)
	}
	defer outThumb.Close()

	// Encode thumbnail based on original format
	switch ext {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(outThumb, thumbnailImg, nil)
	case ".png":
		err = png.Encode(outThumb, thumbnailImg)
	default:
		// Default to JPEG
		err = jpeg.Encode(outThumb, thumbnailImg, nil)
	}

	if err != nil {
		return fmt.Errorf("failed to encode thumbnail: %w", err)
	}

	return nil
}