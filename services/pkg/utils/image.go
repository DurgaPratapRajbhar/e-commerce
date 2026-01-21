package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

// ImageConfig holds image processing configuration
type ImageConfig struct {
	MaxFileSize    int64
	AllowedFormats []string
	MaxWidth       uint
	MaxHeight      uint
	Quality        uint
}

// DefaultImageConfig provides default image processing settings
var DefaultImageConfig = &ImageConfig{
	MaxFileSize:    5 * 1024 * 1024, // 5MB
	AllowedFormats: []string{"jpeg", "jpg", "png", "gif", "webp"},
	MaxWidth:       1920,
	MaxHeight:      1080,
	Quality:        80,
}

// ValidateImageFile validates an uploaded image file
func ValidateImageFile(file *multipart.FileHeader) error {
	// Check file size
	if file.Size > DefaultImageConfig.MaxFileSize {
		return fmt.Errorf("file size too large: %d bytes, max allowed: %d bytes", file.Size, DefaultImageConfig.MaxFileSize)
	}

	// Check file extension
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !isAllowedFormat(ext) {
		return fmt.Errorf("file format not allowed: %s, allowed formats: %v", ext, DefaultImageConfig.AllowedFormats)
	}

	return nil
}

// isAllowedFormat checks if the file format is allowed
func isAllowedFormat(ext string) bool {
	// Remove the dot from extension
	ext = strings.TrimPrefix(ext, ".")
	
	for _, format := range DefaultImageConfig.AllowedFormats {
		if strings.ToLower(ext) == format {
			return true
		}
	}
	return false
}

// ResizeImage resizes an image to the specified dimensions
func ResizeImage(src []byte, width, height uint) ([]byte, error) {
	// Decode the image
	img, format, err := image.Decode(bytes.NewReader(src))
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}

	// Resize the image
	resizedImg := resize.Resize(width, height, img, resize.Lanczos3)

	// Encode the resized image
	var buf bytes.Buffer
	switch format {
	case "jpeg", "jpg":
		err = jpeg.Encode(&buf, resizedImg, &jpeg.Options{Quality: int(DefaultImageConfig.Quality)})
	case "png":
		err = png.Encode(&buf, resizedImg)
	default:
		return nil, fmt.Errorf("unsupported image format: %s", format)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to encode resized image: %v", err)
	}

	return buf.Bytes(), nil
}

// ResizeImageByMaxDimension resizes an image while maintaining aspect ratio
func ResizeImageByMaxDimension(src []byte, maxDimension uint) ([]byte, error) {
	img, format, err := image.Decode(bytes.NewReader(src))
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}

	// Get original dimensions
	bounds := img.Bounds()
	originalWidth := uint(bounds.Dx())
	originalHeight := uint(bounds.Dy())

	// Calculate new dimensions while maintaining aspect ratio
	var newWidth, newHeight uint
	if originalWidth > originalHeight {
		if originalWidth > maxDimension {
			newWidth = maxDimension
			newHeight = (originalHeight * maxDimension) / originalWidth
		} else {
			newWidth = originalWidth
			newHeight = originalHeight
		}
	} else {
		if originalHeight > maxDimension {
			newHeight = maxDimension
			newWidth = (originalWidth * maxDimension) / originalHeight
		} else {
			newWidth = originalWidth
			newHeight = originalHeight
		}
	}

	// Resize the image
	resizedImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)

	// Encode the resized image
	var buf bytes.Buffer
	switch format {
	case "jpeg", "jpg":
		err = jpeg.Encode(&buf, resizedImg, &jpeg.Options{Quality: int(DefaultImageConfig.Quality)})
	case "png":
		err = png.Encode(&buf, resizedImg)
	default:
		return nil, fmt.Errorf("unsupported image format: %s", format)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to encode resized image: %v", err)
	}

	return buf.Bytes(), nil
}

// ProcessImage uploads and processes an image file
func ProcessImage(file multipart.File, filename string) ([]byte, error) {
	// Read the file contents
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	// Validate image format by trying to decode it
	_, _, err = image.DecodeConfig(bytes.NewReader(fileBytes))
	if err != nil {
		return nil, fmt.Errorf("invalid image file: %v", err)
	}

	// Resize image if it exceeds max dimensions
	img, _, err := image.Decode(bytes.NewReader(fileBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}

	bounds := img.Bounds()
	width := uint(bounds.Dx())
	height := uint(bounds.Dy())

	var processedImage []byte

	if width > DefaultImageConfig.MaxWidth || height > DefaultImageConfig.MaxHeight {
		processedImage, err = ResizeImageByMaxDimension(fileBytes, DefaultImageConfig.MaxWidth)
		if err != nil {
			return nil, fmt.Errorf("failed to resize image: %v", err)
		}
	} else {
		processedImage = fileBytes
	}

	return processedImage, nil
}

// SaveImage saves an image to the specified path
func SaveImage(imageData []byte, path string) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	// Create file
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// Write image data to file
	_, err = file.Write(imageData)
	if err != nil {
		return fmt.Errorf("failed to write image data: %v", err)
	}

	return nil
}