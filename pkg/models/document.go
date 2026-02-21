package models

import (
	"time"
)

// Document represents a document
type Document struct {
	FileName     string    `json:"file_name,omitempty"`
	FileSize     string    `json:"file_size,omitempty"`
	UploadedBy   string    `json:"uploaded_by,omitempty"`
	DocumentId   string    `json:"document_id,omitempty"`
	FileUrl      string    `json:"file_url,omitempty"`
	MimeType     string    `json:"mime_type,omitempty"`
	Hash         string    `json:"hash,omitempty"`
	UploadedAt   time.Time `json:"uploaded_at,omitempty"`
	DocumentType string    `json:"document_type,omitempty"`
}
