package models

import (
	"time"
)

// FileUploadRequest represents a file_upload_request
type FileUploadRequest struct {
	ExpiresAt     time.Time `json:"expires_at,omitempty"`
	TenantId      string    `json:"tenant_id"`
	Filename      string    `json:"filename,omitempty"`
	FileType      *FileType `json:"file_type,omitempty"`
	ReferenceType string    `json:"reference_type,omitempty"`
	Content       string    `json:"content,omitempty"`
	ContentType   string    `json:"content_type,omitempty"`
	ReferenceId   string    `json:"reference_id"`
	IsPublic      bool      `json:"is_public,omitempty"`
}
