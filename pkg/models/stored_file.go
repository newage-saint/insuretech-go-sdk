package models

import (
	"time"
)

// StoredFile represents a stored_file
type StoredFile struct {
	SizeBytes     string    `json:"size_bytes,omitempty"`
	StorageKey    string    `json:"storage_key,omitempty"`
	Bucket        string    `json:"bucket,omitempty"`
	CdnUrl        string    `json:"cdn_url,omitempty"`
	ReferenceId   string    `json:"reference_id,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
	FileId        string    `json:"file_id,omitempty"`
	Filename      string    `json:"filename,omitempty"`
	Url           string    `json:"url,omitempty"`
	UploadedBy    string    `json:"uploaded_by,omitempty"`
	ContentType   string    `json:"content_type,omitempty"`
	FileType      *FileType `json:"file_type,omitempty"`
	ReferenceType string    `json:"reference_type,omitempty"`
	IsPublic      bool      `json:"is_public,omitempty"`
	ExpiresAt     time.Time `json:"expires_at,omitempty"`
	TenantId      string    `json:"tenant_id,omitempty"`
}
