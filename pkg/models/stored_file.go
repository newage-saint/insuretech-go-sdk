package models

import (
	"time"
)

// StoredFile represents a stored_file
type StoredFile struct {
	StorageKey    string    `json:"storage_key,omitempty"`
	CdnUrl        string    `json:"cdn_url,omitempty"`
	FileType      *FileType `json:"file_type,omitempty"`
	IsPublic      bool      `json:"is_public,omitempty"`
	ExpiresAt     time.Time `json:"expires_at,omitempty"`
	Filename      string    `json:"filename,omitempty"`
	Url           string    `json:"url,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
	ContentType   string    `json:"content_type,omitempty"`
	SizeBytes     string    `json:"size_bytes,omitempty"`
	ReferenceId   string    `json:"reference_id,omitempty"`
	FileId        string    `json:"file_id,omitempty"`
	TenantId      string    `json:"tenant_id,omitempty"`
	Bucket        string    `json:"bucket,omitempty"`
	ReferenceType string    `json:"reference_type,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UploadedBy    string    `json:"uploaded_by,omitempty"`
}
