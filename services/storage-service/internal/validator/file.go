package validator

import (
	"errors"
	"fmt"
	"mime/multipart"
)

type FileValidator struct {
	MaxFileSize  int64
	AllowedTypes []string
}

func NewFileValidator(maxFileSize int64, allowedTypes []string) *FileValidator {
	return &FileValidator{
		MaxFileSize:  maxFileSize,
		AllowedTypes: allowedTypes,
	}
}

func (v *FileValidator) ValidateFile(file *multipart.FileHeader) error {
	// Check file size
	if file.Size > v.MaxFileSize {
		return errors.New(fmt.Sprintf("file size exceeds maximum allowed size of %d bytes", v.MaxFileSize))
	}

	// Check file type is done separately in handlers since we need to check the content type
	return nil
}

func (v *FileValidator) GetAllowedTypes() []string {
	return v.AllowedTypes
}