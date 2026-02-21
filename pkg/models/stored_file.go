package models

import (
	"time"
)

// StoredFile represents a stored_file
type StoredFile struct {
	FileId        string    `json:"file_id,omitempty"`
	Filename      string    `json:"filename,omitempty"`
	StorageKey    string    `json:"storage_key,omitempty"`
	Bucket        string    `json:"bucket,omitempty"`
	ContentType   string    `json:"content_type,omitempty"`
	SizeBytes     string    `json:"size_bytes,omitempty"`
	Url           string    `json:"url,omitempty"`
	IsPublic      bool      `json:"is_public,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	TenantId      string    `json:"tenant_id,omitempty"`
	CdnUrl        string    `json:"cdn_url,omitempty"`
	FileType      *FileType `json:"file_type,omitempty"`
	ExpiresAt     time.Time `json:"expires_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
	UploadedBy    string    `json:"uploaded_by,omitempty"`
	ReferenceId   string    `json:"reference_id,omitempty"`
	ReferenceType string    `json:"reference_type,omitempty"`
}
