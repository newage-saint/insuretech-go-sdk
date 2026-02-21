package models

import (
	"time"
)

// Document represents a document
type Document struct {
	DocumentType string    `json:"document_type,omitempty"`
	FileUrl      string    `json:"file_url,omitempty"`
	MimeType     string    `json:"mime_type,omitempty"`
	FileSize     string    `json:"file_size,omitempty"`
	Hash         string    `json:"hash,omitempty"`
	UploadedBy   string    `json:"uploaded_by,omitempty"`
	DocumentId   string    `json:"document_id,omitempty"`
	FileName     string    `json:"file_name,omitempty"`
	UploadedAt   time.Time `json:"uploaded_at,omitempty"`
}
