package models

import (
	"time"
)

// FileUploadRequest represents a file_upload_request
type FileUploadRequest struct {
	ReferenceType string    `json:"reference_type,omitempty"`
	IsPublic      bool      `json:"is_public,omitempty"`
	ExpiresAt     time.Time `json:"expires_at,omitempty"`
	Content       string    `json:"content,omitempty"`
	TenantId      string    `json:"tenant_id"`
	Filename      string    `json:"filename,omitempty"`
	ContentType   string    `json:"content_type,omitempty"`
	FileType      *FileType `json:"file_type,omitempty"`
	ReferenceId   string    `json:"reference_id"`
}
