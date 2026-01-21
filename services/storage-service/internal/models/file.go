package models

import "time"

// FileMetadata represents uploaded file information
type FileMetadata struct {
	OriginalName string    `json:"original_name"` // User's original filename
	FileName     string    `json:"file_name"`     // Unique generated filename
	FilePath     string    `json:"file_path"`     // Full path on disk
	FileType     string    `json:"file_type"`     // MIME type (image/jpeg, application/pdf)
	FileSize     int64     `json:"file_size"`     // Size in bytes
	Category     string    `json:"category"`      // products, users, documents, categories
	Subcategory  string    `json:"subcategory"`   // images, avatars, invoices, banners
	URL          string    `json:"url"`           // Public accessible URL
	ThumbnailURL string    `json:"thumbnail_url,omitempty"` // For images only
	MediumURL    string    `json:"medium_url,omitempty"`    // For images only
	UploadedAt   time.Time `json:"uploaded_at"`
}

// UploadResponse is returned after successful upload
type UploadResponse struct {
	Success      bool   `json:"success"`
	FileName     string `json:"file_name"`
	OriginalName string `json:"original_name"`
	FileType     string `json:"file_type"`
	FileSize     int64  `json:"file_size"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnail_url,omitempty"` // Only for images
	MediumURL    string `json:"medium_url,omitempty"`    // Only for images
	Message      string `json:"message"`
}

// ErrorResponse is returned on failure
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

// DeleteResponse is returned after file deletion
type DeleteResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// UploadRequest contains upload parameters
type UploadRequest struct {
	Category    string `json:"category"`    // Required: products, users, documents
	Subcategory string `json:"subcategory"` // Required: images, avatars, invoices
}