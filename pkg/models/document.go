package models

import (
	"time"
)

// Document represents a document
type Document struct {
	FileSize     string    `json:"file_size,omitempty"`
	Hash         string    `json:"hash,omitempty"`
	UploadedAt   time.Time `json:"uploaded_at,omitempty"`
	UploadedBy   string    `json:"uploaded_by,omitempty"`
	FileUrl      string    `json:"file_url,omitempty"`
	DocumentId   string    `json:"document_id,omitempty"`
	DocumentType string    `json:"document_type,omitempty"`
	FileName     string    `json:"file_name,omitempty"`
	MimeType     string    `json:"mime_type,omitempty"`
}
